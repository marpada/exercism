// Package weather : forecast current weather condition.
package weather

// CurrentCondition : foo.
var CurrentCondition string

// CurrentLocation : bar.
var CurrentLocation string

// Forecast: Run forecast condition for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
