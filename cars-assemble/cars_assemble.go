package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed >= 9 {
        return 0.77
    } else if speed >= 5 {
        return 0.9
    } else if speed >= 1 {
        return 1.0
    } else {
        return 0.0
    }
    
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
    if speed == 0 {return 0.0}
    rate := SuccessRate(speed)
    return float64(221*speed)*rate
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
    produced := CalculateProductionRatePerHour(speed)
    return int(produced / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
    produced := CalculateProductionRatePerHour(speed)
    if limit < produced {
        return limit
    } else {
        return produced
    }
}
