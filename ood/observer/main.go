package main

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer/weather"
)

func main() {
	//Create subject
	weatherData := weather.NewWeatherData()

	//Create and register first observer
	statisticsDisplay := weather.NewStatisticsDisplay()
	weatherData.RegisterObserver(statisticsDisplay)

	//Create and register second observer
	display := &weather.Display{}
	weatherData.RegisterObserver(display)

	weatherData.SetMeasurements(3, 0.7, 760)
	weatherData.SetMeasurements(4, 0.8, 761)

	weatherData.RemoveObserver(display)

	weatherData.SetMeasurements(10, 0.8, 761)
	weatherData.SetMeasurements(-10, 0.8, 761)
}
