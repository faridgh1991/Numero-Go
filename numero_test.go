package numero

import (
	"testing"
)

func TestDigitOnly(t *testing.T) {
	tests := []struct {
		input_text      string
		expected_result bool
	}{
		{"abc", false},
		{"1234", true},
		{"۱۲۳۴", true},
		{"۱۲بA", false},
	}

	for _, test := range tests {
		result := DigitOnly(test.input_text)
		if result != test.expected_result {
			t.Errorf("Expected %v got %v.", test.expected_result, result)
		}
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		input_text, expected_text string
	}{
		{"abc", "abc"},
		{"1234", "1234"},
		{"۱۲۳۴", "1234"},
		{"۱۲بA", "12بA"},
	}

	for _, test := range tests {
		result := Normalize(test.input_text)
		if result != test.expected_text {
			t.Errorf("Expected %v got %v.", test.expected_text, result)
		}
	}
}

func TestNormalizeAsNumber(t *testing.T) {
	tests := []struct {
		input_text      string
		expected_number interface{}
	}{
		{"1234", 1234},
		{"12.34", 12.34},
		{"۱۲۳۴", 1234},
		{"۱۲.۳۴", 12.34},
	}

	for _, test := range tests {
		result, _ := NormalizeAsNumber(test.input_text)
		if result != test.expected_number {
			t.Errorf("Expected %v got %v.", test.expected_number, result)
		}
	}
}

func TestRemoveNonDigits(t *testing.T) {
	tests := []struct {
		input_text       string
		input_exceptions string
		expected_result  string
	}{
		{"1234abcd", "", "1234"},
		{"12.34abcd", "", "1234"},
		{"۱۲۳۴abcd", "", "1234"},
		{"۱۲.۳۴abcd", "", "1234"},

		{"1234abcd", "b", "1234b"},
		{"12.34abcd", "b", "1234b"},
	}

	for _, test := range tests {
		var result string
		if test.input_exceptions != "" {
			result = RemoveNonDigits(test.input_text, test.input_exceptions)
		} else {
			result = RemoveNonDigits(test.input_text)
		}
		if result != test.expected_result {
			t.Errorf("Expected %v got %v.", test.expected_result, result)
		}
	}
}
