package weather

import (
	"bytes"
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/observer_channels/observer"
	"math"
	"os"
	"strings"
)

const DELIMITER_LINE_LENGTH = 21

type Display struct {
	InChan chan *observer.Event
}

func (this Display) Handle(event *observer.Event) {
	wd, ok := event.Data.(WeatherData)
	if ok {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("Location: %s, measurements:\n", wd.location))
		switch event.Type {
		case TEMPERATURE_CHANGED:
			buffer.WriteString(fmt.Sprintf("Current Temperature: %v\n", wd.temperature))
		case HUMIDITY_CHANGED:
			buffer.WriteString(fmt.Sprintf("Current Humidity: %v\n", wd.humidity))
		case PRESSURE_CHANGED:
			buffer.WriteString(fmt.Sprintf("Current Pressure: %v\n", wd.pressure))
		}
		fmt.Print(buffer.String())
		fmt.Println(strings.Repeat("-", DELIMITER_LINE_LENGTH))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (Display::Update)\n", wd)
	}
}

func (this *Display) Close() {
	close(this.InChan)
}

func NewDisplay() *Display {
	d := &Display{}
	inChan := make(chan *observer.Event)
	d.InChan = inChan
	observer.StartSubscribing(inChan, d)

	return d
}

type StatisticsObject struct {
	name              string
	minValue          float64
	maxValue          float64
	accumulativeValue float64
	countMeasurements float64
}

func (this *StatisticsObject) updateStatistics(wd WeatherData) (result string) {
	if this.minValue > wd.temperature {
		this.minValue = wd.temperature
	}
	if this.maxValue < wd.temperature {
		this.maxValue = wd.temperature
	}
	this.countMeasurements++
	this.accumulativeValue += wd.temperature
	meanValue := this.accumulativeValue / this.countMeasurements

	result = fmt.Sprintf("Min %s: %v\nMax %s: %v\nMean %s: %v\n", this.name, this.minValue, this.name, this.maxValue, this.name, meanValue)
	result += fmt.Sprintln(strings.Repeat("*", DELIMITER_LINE_LENGTH))
	return
}

type StatisticsDisplay struct {
	temperatureStatisticsObject *StatisticsObject
	humidityStatisticsObject    *StatisticsObject
	pressureStatisticsObject    *StatisticsObject
	InChan                      chan *observer.Event
}

func (this *StatisticsDisplay) Handle(event *observer.Event) {
	wd, ok := event.Data.(WeatherData)
	if ok {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("Location: %s, statistics:\n", wd.location))
		switch event.Type {
		case TEMPERATURE_CHANGED:
			buffer.WriteString(this.temperatureStatisticsObject.updateStatistics(wd))
		case HUMIDITY_CHANGED:
			buffer.WriteString(this.humidityStatisticsObject.updateStatistics(wd))
		case PRESSURE_CHANGED:
			buffer.WriteString(this.pressureStatisticsObject.updateStatistics(wd))
		}
		fmt.Print(buffer.String())
		fmt.Println(strings.Repeat("-", DELIMITER_LINE_LENGTH))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (StatisticsDisplay::Update)\n", wd)
	}
}

func newStatisticsObject(name string) *StatisticsObject {
	statisticsObject := &StatisticsObject{name, math.Inf(1), math.Inf(-1), 0, 0}
	return statisticsObject
}

func NewStatisticsDisplay() *StatisticsDisplay {
	temperatureStatisticsObject := newStatisticsObject("Temperature")
	humidityStatisticsObject := newStatisticsObject("Humidity")
	pressureStatisticsObject := newStatisticsObject("Pressure")
	inChan := make(chan *observer.Event)

	d := &StatisticsDisplay{temperatureStatisticsObject, humidityStatisticsObject, pressureStatisticsObject, inChan}
	observer.StartSubscribing(inChan, d)

	return d
}

//Remove from subject first!
func (this *StatisticsDisplay) Close() {
	close(this.InChan)
}
