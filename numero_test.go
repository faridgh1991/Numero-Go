package numero

import (
	"testing"
)

func TestDigitOnly(t *testing.T) {
	tests := []struct {
		inputText      string
		expectedResult bool
	}{
		{"abc", false},
		{"1234", true},
		{"۱۲۳۴", true},
		{"۱۲بA", false},
	}

	for _, test := range tests {
		result := DigitOnly(test.inputText)
		if result != test.expectedResult {
			t.Errorf("Expected %v got %v.", test.expectedResult, result)
		}
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		inputText, expectedText string
	}{
		{"abc", "abc"},
		{"1234", "1234"},
		{"۱۲۳۴", "1234"},
		{"۱۲بA", "12بA"},
	}

	for _, test := range tests {
		result := Normalize(test.inputText)
		if result != test.expectedText {
			t.Errorf("Expected %v got %v.", test.expectedText, result)
		}
	}
}

func TestNormalizeAsNumber(t *testing.T) {
	tests := []struct {
		inputText      string
		expectedNumber interface{}
	}{
		{"1234", 1234},
		{"12.34", 12.34},
		{"۱۲۳۴", 1234},
		{"۱۲.۳۴", 12.34},
	}

	for _, test := range tests {
		result, _ := NormalizeAsNumber(test.inputText)
		if result != test.expectedNumber {
			t.Errorf("Expected %v got %v.", test.expectedNumber, result)
		}
	}
}

func TestRemoveNonDigits(t *testing.T) {
	tests := []struct {
		inputText       string
		inputExceptions []rune
		expectedResult  string
	}{
		{"1234abcd", []rune{}, "1234"},
		{"12.34abcd", []rune{}, "1234"},
		{"۱۲🙃۳۴abcd", []rune{}, "1234"},
		{"۱۲.۳۴abcd", []rune{}, "1234"},

		{"1234abcd", []rune{'b'}, "1234b"},
		{"12.34abcd", []rune{'b'}, "1234b"},
	}

	for _, test := range tests {
		var result string
		if len(test.inputExceptions) > 0 {
			result = RemoveNonDigits(test.inputText, test.inputExceptions...)
		} else {
			result = RemoveNonDigits(test.inputText)
		}
		if result != test.expectedResult {
			t.Errorf("Expected %v got %v.", test.expectedResult, result)
		}
	}
}
