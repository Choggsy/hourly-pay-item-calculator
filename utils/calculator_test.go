package utils

import "testing"

func TestCalculateWorkHours(t *testing.T) {
	t.Run("StandardCase", func(t *testing.T) {
		expected := 10.0
		result := CalculateWorkHours(10.0, 100.0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("HalfHourRate", func(t *testing.T) {
		expected := 2.0
		result := CalculateWorkHours(25.0, 50.0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("ZeroHourlyRate", func(t *testing.T) {
		expected := 0.0
		result := CalculateWorkHours(0.0, 100.0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("PenceValue", func(t *testing.T) {
		expected := 5.0
		result := CalculateWorkHours(15.5, 77.5)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
