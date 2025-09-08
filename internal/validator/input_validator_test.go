package validator

import "testing"

func Test_ValidInputReturnsParsedValue(t *testing.T) {
	t.Run("successful numeric parse", func(t *testing.T) {
		result, _ := ValidateAndParseFloat("123.45")
		expected := 123.45
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})
	t.Run("valid input is trimmed and parsed", func(t *testing.T) {
		result, _ := ValidateAndParseFloat("   99.9  ")
		expected := 99.9
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})
	t.Run("integer value returns float", func(t *testing.T) {
		result, _ := ValidateAndParseFloat("100")
		expected := 100.00
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})
	t.Run("zero return zero", func(t *testing.T) {
		result, _ := ValidateAndParseFloat("0")
		expected := 0.0
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
	})
}

func Test_ReturnsCharacterValidationError(t *testing.T) {
	t.Run("letter input returns character validation error", func(t *testing.T) {
		_, err := ValidateAndParseFloat("12a.3")
		if err == nil || err.Error() != "invalid input: only numeric values are allowed" {
			t.Errorf("Expected character validation error: %v", err)
		}
	})
	t.Run("special character input returns CharacterValidationError", func(t *testing.T) {
		_, err := ValidateAndParseFloat("12.3$")
		if err == nil || err.Error() != "invalid input: only numeric values are allowed" {
			t.Errorf("Expected character validation error: %v", err)
		}
	})
}

func Test_ReturnsNegativeValueError(t *testing.T) {
	t.Run("negative input returns NegativeValueError", func(t *testing.T) {
		_, err := ValidateAndParseFloat("-10.0")
		if err == nil || err.Error() != "negative values are not allowed: please enter a valid numeric value" {
			t.Errorf("Expected error: %v", err)
		}
	})
}

func TestReturnsParseError(t *testing.T) {
	t.Run("multiple decimal points returns parse error", func(t *testing.T) {
		_, err := ValidateAndParseFloat("12.3.4")
		if err == nil || err.Error() != "failed to parse number: please enter a valid numeric value" {
			t.Errorf("Expected error: %v", err)
		}
	})
	t.Run("all non numeric input returns parse validation error", func(t *testing.T) {
		result, err := ValidateAndParseFloat("...")
		expected := 0.0
		if result != expected {
			t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		}
		if err == nil || err.Error() != "failed to parse number: please enter a valid numeric value" {
			t.Errorf("Expected error: %v", err)
		}
	})
	t.Run("blank value returns parse error", func(t *testing.T) {
		_, err := ValidateAndParseFloat("")
		if err == nil || err.Error() != "failed to parse number: please enter a valid numeric value" {
			t.Errorf("Expected error: %v", err)
		}
	})
}
