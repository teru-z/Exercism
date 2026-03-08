// Package weather provides helpers for tracking and reporting weather forecasts.
package weather

var (
	// CurrentCondition represents the current weather condition of a certain city.
	CurrentCondition string
	// CurrentLocation represents the city associated with CurrentCondition.
	CurrentLocation string
)

// Forecast updates the current weather state and returns a formatted forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
