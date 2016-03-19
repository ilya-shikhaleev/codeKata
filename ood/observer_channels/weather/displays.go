package weather

import (
	"fmt"
	"github.com/ilya-shikhaleev/codeKata/ood/observer_channels/observer"
	"math"
	"os"
	"strings"
)

const DELIMITER_LINE_LENGTH = 21

type Display struct {
	InChan chan interface{}
}

func (this Display) Handle(data interface{}) {
	weatherData, ok := data.(WeatherData)
	if ok {
		fmt.Println("Display ", weatherData.location)
		//fmt.Printf("Location: %s measurements:\nCurrent Temperature: %v\nCurrent Humidity: %v\nCurrent Pressure: %v\n",
		//	weatherData.location, weatherData.temperature, weatherData.humidity, weatherData.pressure)
		//fmt.Println(strings.Repeat("-", DELIMITER_LINE_LENGTH))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (Display::Update)\n", data)
	}
}

func (this *Display) Close() {
	close(this.InChan)
}

func NewDisplay() *Display {
	d := &Display{}
	inChan := make(chan interface{})
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

func (this *StatisticsObject) updateStatistics(weatherData WeatherData) {
	if this.minValue > weatherData.temperature {
		this.minValue = weatherData.temperature
	}
	if this.maxValue < weatherData.temperature {
		this.maxValue = weatherData.temperature
	}
	this.countMeasurements++
	this.accumulativeValue += weatherData.temperature

	meanValue := this.accumulativeValue / this.countMeasurements
	fmt.Printf("Min %s: %v\nMax %s: %v\nMean %s: %v\n", this.name, this.minValue, this.name, this.maxValue, this.name, meanValue)
	fmt.Println(strings.Repeat("*", DELIMITER_LINE_LENGTH))
}

type StatisticsDisplay struct {
	temperatureStatisticsObject *StatisticsObject
	humidityStatisticsObject    *StatisticsObject
	pressureStatisticsObject    *StatisticsObject
	InChan                      chan interface{}
}

func (this *StatisticsDisplay) Handle(data interface{}) {
	weatherData, ok := data.(WeatherData)
	if ok {
		fmt.Printf("Location %v statistics:\n", weatherData.location)
		//this.temperatureStatisticsObject.updateStatistics(weatherData)
		//this.humidityStatisticsObject.updateStatistics(weatherData)
		//this.pressureStatisticsObject.updateStatistics(weatherData)
		//fmt.Println(strings.Repeat("-", DELIMITER_LINE_LENGTH))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (StatisticsDisplay::Update)\n", data)
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
	inChan := make(chan interface{})

	d := &StatisticsDisplay{temperatureStatisticsObject, humidityStatisticsObject, pressureStatisticsObject, inChan}
	observer.StartSubscribing(inChan, d)

	return d
}

func (this *StatisticsDisplay) Close() {
	close(this.InChan)
}
