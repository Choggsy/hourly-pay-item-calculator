package model

type WorkCalculator interface {
	Calculate(hourlyRate, itemPrice float64) float64
}
