package weather

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer_channels/observer"
)

type WeatherData struct {
	location    string
	temperature float64
	humidity    float64
	pressure    float64
	observer.Subject
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

func NewWeatherData(location string) *WeatherData {
	wd := &WeatherData{}
	wd.location = location
	wd.temperature = 0
	wd.humidity = 0
	wd.pressure = 760
	return wd
}
