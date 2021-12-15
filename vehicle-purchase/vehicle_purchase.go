package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var result string
	if option1 <= option2 {
		result = option1
	} else {
		result = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", result)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var ageFactor float64
	if age < 3 {
		ageFactor = 0.8
	} else if age < 10 {
		ageFactor = 0.7
	} else {
		ageFactor = 0.5
	}
	return originalPrice * ageFactor
}
