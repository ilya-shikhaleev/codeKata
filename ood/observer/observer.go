package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

//Observers
type Observer interface {
	Update(data interface{})
}

type Display struct{}

func (this Display) Update(data interface{}) {
	weatherData, ok := data.(WeatherData)
	if ok {
		fmt.Printf("Current Temperature: %v\nCurrent Humidity: %v\nCurrent Pressure: %v\n", weatherData.temperature, weatherData.humidity, weatherData.pressure)
		fmt.Println(strings.Repeat("-", 21))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (Display::Update)\n", data)
	}
}

type StatisticsDisplay struct {
	minTemperature          float64
	maxTemperature          float64
	accumulativeTemperature float64
	countMeasurements       float64
}

func (this *StatisticsDisplay) Update(data interface{}) {
	weatherData, ok := data.(WeatherData)
	if ok {
		if this.minTemperature > weatherData.temperature {
			this.minTemperature = weatherData.temperature
		}
		if this.maxTemperature < weatherData.temperature {
			this.maxTemperature = weatherData.temperature
		}
		this.countMeasurements++
		this.accumulativeTemperature += weatherData.temperature

		fmt.Printf("Min Temperature: %v\nMax Temperature: %v\nMean Temperature: %v\n", this.minTemperature, this.maxTemperature, this.accumulativeTemperature/this.countMeasurements)
		fmt.Println(strings.Repeat("-", 21))
	} else {
		fmt.Fprintf(os.Stderr, "%T is not WeatherData (StatisticsDisplay::Update)\n", data)
	}

}

func NewStatisticsDisplay() *StatisticsDisplay {
	statisticsDisplay := &StatisticsDisplay{math.Inf(1), math.Inf(-1), 0, 0}
	return statisticsDisplay
}

//Observable subject
type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(data interface{})
}

type Subject struct {
	observers []Observer
}

func (this *Subject) RegisterObserver(observer Observer) {
	this.observers = append(this.observers, observer)
}

func (this *Subject) RemoveObserver(observer Observer) {
	for index, currentObserver := range this.observers {
		if currentObserver == observer {
			this.observers = append(this.observers[:index], this.observers[index+1:]...)
		}
	}
}

func (this *Subject) NotifyObservers(data interface{}) {
	for _, observer := range this.observers {
		observer.Update(data)
	}
}

type WeatherData struct {
	temperature float64
	humidity    float64
	pressure    float64
	Subject
}

func (this *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	this.temperature = temperature
	this.humidity = humidity
	this.pressure = pressure

	this.OnMeasurementChange()
}

func (this *WeatherData) OnMeasurementChange() {
	this.NotifyObservers(*this)
}

func NewWeatherData() *WeatherData {
	weatherData := &WeatherData{}
	weatherData.temperature = 0
	weatherData.humidity = 0
	weatherData.pressure = 760
	return weatherData
}

func main() {
	//Create subject
	weatherData := NewWeatherData()

	//Create and register first observer
	statisticsDisplay := NewStatisticsDisplay()
	weatherData.RegisterObserver(statisticsDisplay)

	//Create and register second observer
	display := &Display{}
	weatherData.RegisterObserver(display)

	temperature, humidity, pressure := 0.0, 0.0, 760.0
	for i := 0; i < 3; i++ {
		weatherData.SetMeasurements(temperature, humidity, pressure)
		temperature += 2
		humidity += 2
		pressure += 10
	}

	weatherData.RemoveObserver(display)
	for i := 0; i < 3; i++ {
		weatherData.SetMeasurements(temperature, humidity, pressure)
		temperature -= 1.5
		humidity -= 2
		pressure -= 10
	}
}
