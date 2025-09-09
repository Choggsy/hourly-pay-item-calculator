package calculator

type BasicCalculator struct{}

// Calculate implements WorkCalculator for BasicCalculator.
func (b BasicCalculator) Calculate(hourlyRate, itemPrice float64) float64 {
	if hourlyRate == 0 {
		return 0
	}
	return itemPrice / hourlyRate
}
