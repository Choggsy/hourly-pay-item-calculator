package utils

import "testing"

func TestBasicCalculator(t *testing.T) {
	calculator := BasicCalculator{} // implements WorkCalculator because in Go any type that has the required methods automatically satisfies an interface! that's crazy
	//could alternatively be var calculator WorkCalculator = BasicCalculator{}

	t.Run("StandardCase", func(t *testing.T) {
		expected := 10.0
		result := calculator.Calculate(10.0, 100.0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("HalfHourRate", func(t *testing.T) {
		expected := 2.0
		result := calculator.Calculate(25.0, 50.0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("ZeroHourlyRate", func(t *testing.T) {
		expected := 0.0
		result := calculator.Calculate(0.0, 100.0)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("PenceValue", func(t *testing.T) {
		expected := 5.0
		result := calculator.Calculate(15.5, 77.5)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
