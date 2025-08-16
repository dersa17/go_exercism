// Package weather provides functionality to forecast the weather for a given city.
// It allows setting the current location and retrieving a forecast message that
// describes the weather condition in that location.
package weather


// CurrentCondition stores the current weather condition for the given location.
// This value can be updated to reflect changes in the weather forecast.
var CurrentCondition string

// CurrentLocation stores the name of the location for which the weather
// forecast applies. This value can be set to any city or region name.
var CurrentLocation string


// Forecast returns a string with the weather forecast for the current location.
// The forecast includes both the location name and the current weather condition. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
