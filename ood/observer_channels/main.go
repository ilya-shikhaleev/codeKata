package main

import (
	"github.com/ilya-shikhaleev/codeKata/ood/observer_channels/weather"
	"time"
)

func main() {
	inWd := weather.NewWeatherData("in")
	outWd := weather.NewWeatherData("out")

	d := weather.NewDisplay()
	defer d.Close()

	sd := weather.NewStatisticsDisplay()
	defer sd.Close()

	inWd.RegisterObserver([]string{weather.TEMPERATURE_CHANGED}, d.InChan)
	inWd.RegisterObserver([]string{weather.TEMPERATURE_CHANGED}, sd.InChan)

	outWd.RegisterObserver([]string{weather.HUMIDITY_CHANGED}, d.InChan)
	outWd.RegisterObserver([]string{weather.PRESSURE_CHANGED}, sd.InChan)

	inWd.SetMeasurements(3, 0.7, 760)
	inWd.SetMeasurements(4, 0.8, 761)

	inWd.RemoveObserver(d.InChan)

	inWd.SetMeasurements(10, 0.8, 761)
	inWd.SetMeasurements(-10, 0.8, 761)

	outWd.SetMeasurements(23, 10, 700)
	time.Sleep(1 * time.Second)
}
