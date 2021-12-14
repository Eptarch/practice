// Package weather provides tools to predict
// weather conditions.
package weather

// CurrentCondition stores a string and can be used in Forecast(). 
var CurrentCondition string
// CurrentLocation stores a string and can be used in Forecast(). 
var CurrentLocation string

// Forecast  returns a string value equal to concatenation of
// CurrentLocation, some text and then CurrentCondition
// Forecast accepts strings of a city and a condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

