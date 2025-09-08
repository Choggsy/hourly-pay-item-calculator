package main

import (
	"hourly-pay-item-calculator/utils"
	"testing"
)

func Test_CalculateWorkHours(t *testing.T) {
	calculator := utils.BasicCalculator{}

	t.Run("standard case", func(t *testing.T) {
		result := calculator.Calculate(10.0, 100.0)
		expected := 10.0
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})

	t.Run("zero hourly rate returns zero", func(t *testing.T) {
		result := calculator.Calculate(0.0, 100.0)
		expected := 0.0
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})

	t.Run("pence value returns correct float", func(t *testing.T) {
		result := calculator.Calculate(15.5, 77.5)
		expected := 5.0
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})
}
