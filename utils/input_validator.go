package utils

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ValidateAndParseFloat checks if the input is a valid float and returns it, or an error if invalid.
func ValidateAndParseFloat(rawInput string) (float64, error) {
	trimmedInput := strings.TrimSpace(rawInput)
	parsedFloat, parseErr := strconv.ParseFloat(trimmedInput, 64)

	validationErrors := []error{ //list of all error occurring
		validateNumericCharacters(trimmedInput),
		validateParseSuccess(parseErr),
		validateNonNegativeValue(parsedFloat),
	}

	//Return first error found
	for _, validationErr := range validationErrors { //range makes for loop act like for (Error validationErr : validationErrors) in java
		if validationErr != nil {
			return 0.0, validationErr
		}
	}

	return parsedFloat, nil
}

// validateNumericCharacters ensures the input contains only digits, decimal points, or minus signs.
func validateNumericCharacters(input string) error {
	for _, char := range input {
		if unicode.IsLetter(char) || (!unicode.IsDigit(char) && char != '.' && char != '-') {
			return errors.New("invalid input: only numeric values are allowed")
		}
	}
	return nil
}

// validateParseSuccess returns an error if parsing failed.
func validateParseSuccess(err error) error {
	if err != nil {
		return errors.New("failed to parse number: please enter a valid numeric value")
	}
	return nil
}

// validateNonNegativeValue ensures the parsed float is not negative.
func validateNonNegativeValue(value float64) error {
	if value < 0.0 {
		return errors.New("negative values are not allowed")
	}
	return nil
}
