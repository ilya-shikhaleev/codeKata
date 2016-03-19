package weather

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer_channels/observer"
)

const (
	TEMPERATURE_CHANGED = "temperature_changed"
	HUMIDITY_CHANGED    = "humidity_changed"
	PRESSURE_CHANGED    = "pressure_changed"
)

type WeatherData struct {
	location    string
	temperature float64
	humidity    float64
	pressure    float64
	observer.Subject
}

func (this *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	var events []string
	if this.temperature != temperature {
		this.temperature = temperature
		events = append(events, TEMPERATURE_CHANGED)
	}
	if this.humidity != humidity {
		this.humidity = humidity
		events = append(events, HUMIDITY_CHANGED)
	}
	if this.pressure != pressure {
		this.pressure = pressure
		events = append(events, PRESSURE_CHANGED)
	}
	this.OnMeasurementChange(events)
}

func (this *WeatherData) OnMeasurementChange(events []string) {
	this.NotifyObservers(events, *this)
}

func NewWeatherData(location string) *WeatherData {
	wd := &WeatherData{}
	wd.location = location
	wd.temperature = 0
	wd.humidity = 0
	wd.pressure = 760
	wd.Observers = make(map[string][]chan *observer.Event)
	return wd
}
