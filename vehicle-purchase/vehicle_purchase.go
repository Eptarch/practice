package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    text := " is clearly the better choice."
	if option1 < option2 {
		return option1 + text
    }
	return option2 + text
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    ageFactor := 0.5
	if age < 3 {
		ageFactor = 0.8
	} else if age < 10 {
		ageFactor = 0.7
    }
	return originalPrice * ageFactor
}
