// Package weather tracking.
package weather

// CurrentCondition - current weather condition.
var CurrentCondition string

// CurrentLocation - current location.
var CurrentLocation string

// Forecast weather based on city and condition provided.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
