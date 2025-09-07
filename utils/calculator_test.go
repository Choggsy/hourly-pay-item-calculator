package utils

import "testing"

// Test case names must begin with Test!
// it uses snakecase after that I think based on the auto generated cases
func Test_Calculates(t *testing.T) {
	expected := 10.0
	result := CalculateWorkHours(10.0, 100.0)
	if result != expected {
		t.Errorf("Test failed, expected: %v, got: %v", expected, result)
		// like javas assertThat(expected).isEqualTo(result); but with built in error message
	}
}

// to run an individual test its - go test -run TestName
func Test_HalfHourRate(t *testing.T) {
	expected := 2.0
	result := CalculateWorkHours(25.0, 50.0)
	if result != expected {
		t.Errorf("Test failed, expected: %v, got: %v", expected, result)
	}
}

func Test_ZeroHourlyRate(t *testing.T) {
	expected := 0.0
	result := CalculateWorkHours(25.0, 0.0)
	if result != expected {
		t.Errorf("Test failed, expected: %v, got: %v", expected, result)
	}
}

func Test_PenceValue(t *testing.T) {
	expected := 5.0
	result := CalculateWorkHours(15.5, 77.5)
	if result != expected {
		t.Errorf("Test failed, expected: %v, got: %v", expected, result)
	}
}
