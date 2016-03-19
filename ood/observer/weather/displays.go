package weather

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const DELIMITER_LINE_LENGTH = 21

type Display struct{}

func (this Display) Update(data interface{}) {
	weatherData, ok := data.(WeatherData)
	if ok {
		fmt.Printf("Location: %s measurements:\nCurrent Temperature: %v\nCurrent Humidity: %v\nCurrent Pressure: %v\n",
			weatherData.location, weatherData.temperature, weatherData.humidity, weatherData.pressure)
		fmt.Println(strings.Repeat("-", DELIMITER_LINE_LENGTH))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (Display::Update)\n", data)
	}
}

type StatisticsObject struct {
	statisticsName    string
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
	fmt.Printf("Min %s: %v\nMax %s: %v\nMean %s: %v\n", this.statisticsName, this.minValue, this.statisticsName, this.maxValue, this.statisticsName, meanValue)
	fmt.Println(strings.Repeat("*", DELIMITER_LINE_LENGTH))
}

type StatisticsDisplay struct {
	temperatureStatisticsObject *StatisticsObject
	humidityStatisticsObject    *StatisticsObject
	pressureStatisticsObject    *StatisticsObject
}

func (this *StatisticsDisplay) Update(data interface{}) {
	weatherData, ok := data.(WeatherData)
	if ok {
		fmt.Printf("Location %v statistics:\n", weatherData.location)
		this.temperatureStatisticsObject.updateStatistics(weatherData)
		this.humidityStatisticsObject.updateStatistics(weatherData)
		this.pressureStatisticsObject.updateStatistics(weatherData)
		fmt.Println(strings.Repeat("-", DELIMITER_LINE_LENGTH))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (StatisticsDisplay::Update)\n", data)
	}
}

func newStatisticsObject(statisticsName string) *StatisticsObject {
	statisticsObject := &StatisticsObject{statisticsName, math.Inf(1), math.Inf(-1), 0, 0}
	return statisticsObject
}

func NewStatisticsDisplay() *StatisticsDisplay {
	temperatureStatisticsObject := newStatisticsObject("Temperature")
	humidityStatisticsObject := newStatisticsObject("Humidity")
	pressureStatisticsObject := newStatisticsObject("Pressure")
	statisticsDisplay := &StatisticsDisplay{temperatureStatisticsObject, humidityStatisticsObject, pressureStatisticsObject}
	return statisticsDisplay
}
