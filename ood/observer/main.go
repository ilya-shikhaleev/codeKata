package main

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer/weather"
)

func main() {
	//Create subject
	inWeatherData := weather.NewWeatherData("in")
	outWeatherData := weather.NewWeatherData("out")

	//Create and register first observer
	statisticsDisplay := weather.NewStatisticsDisplay()
	inWeatherData.RegisterObserver(statisticsDisplay, 1.0)
	outWeatherData.RegisterObserver(statisticsDisplay, 2.0)

	//Create and register second observer
	display := &weather.Display{}
	inWeatherData.RegisterObserver(display, 2.0)
	outWeatherData.RegisterObserver(display, 1.0)

	inWeatherData.SetMeasurements(3, 0.7, 760)
	inWeatherData.SetMeasurements(4, 0.8, 761)

	inWeatherData.RemoveObserver(display)

	inWeatherData.SetMeasurements(10, 0.8, 761)
	inWeatherData.SetMeasurements(-10, 0.8, 761)

	outWeatherData.SetMeasurements(23, 10, 700)
}
