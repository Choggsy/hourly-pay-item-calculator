package utils

// CalculateWorkHours function requires 2 double parameters and returns a double.
// In Go, functions are only accessible outside their package if they start with a capital letter!!! weird
func CalculateWorkHours(hourlyRate, itemPrice float64) float64 {
	if hourlyRate == 0 {
		return 0 // can't divide by 0 :)
	}
	return itemPrice / hourlyRate
}
