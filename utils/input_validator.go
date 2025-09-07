package utils

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ValidateAndParseFloat checks if the input is a valid float and returns it, or an error if invalid.
func ValidateAndParseFloat(rawInput string) (float64, error) {
	const zeroValue = 0.0
	trimmedInput := strings.TrimSpace(rawInput)
	characterValidationErr := validateCharacters(trimmedInput)
	parsedFloat, parseErr := strconv.ParseFloat(trimmedInput, 64)
	parseValidationErr := validateParseError(parseErr)
	negativeValueErr := validateNonNegative(parsedFloat)

	if characterValidationErr != nil {
		return zeroValue, characterValidationErr
	}
	if parseValidationErr != nil {
		return zeroValue, parseValidationErr
	}
	if negativeValueErr != nil {
		return zeroValue, negativeValueErr
	}

	return parsedFloat, nil
}

// validateCharacters ensures the input contains only digits, decimal points, or minus signs.
func validateCharacters(input string) error {
	for _, char := range input {
		if unicode.IsLetter(char) || (!unicode.IsDigit(char) && char != '.' && char != '-') {
			return errors.New("invalid input: only numeric values are allowed")
		}
	}
	return nil
}

// validateParseError returns an error if parsing failed.
func validateParseError(err error) error {
	if err != nil {
		return errors.New("failed to parse number: please enter a valid numeric value")
	}
	return nil
}

// validateNonNegative ensures the parsed float is not negative.
func validateNonNegative(value float64) error {
	if value < 0 {
		return errors.New("negative values are not allowed")
	}
	return nil
}
