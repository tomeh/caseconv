package caseconv

import (
	"testing"
)

func TestToCamel(t *testing.T) {
	tests := map[string]string{
		"":         "",
		"1":        "1",
		"12":       "12",
		"one two":  "oneTwo",
		"one  two": "oneTwo",
		"one_-two": "oneTwo",
		"abc":      "abc",
		"abc-xyz":  "abcXyz",
		"abc_xyz":  "abcXyz",
		"abcXyz":   "abcXyz",
		"abc\nxyz": "abcXyz",
	}

	for test, expected := range tests {
		result := ToCamel(test)
		if result != expected {
			t.Errorf("with '%s', expected '%s', got '%s'", test, expected, result)
		}
	}
}

func TestToSnake(t *testing.T) {
	tests := map[string]string{
		"":           "",
		"1":          "1",
		"12":         "12",
		"one two":    "one_two",
		"one  two":   "one_two",
		"one_-two":   "one_two",
		"abc":        "abc",
		"abc-xyz":    "abc_xyz",
		"abc_xyz":    "abc_xyz",
		"abcXyz":     "abc_xyz",
		"abc\nxyz":   "abc_xyz",
		"abc\n\nxyz": "abc_xyz",
	}

	for test, expected := range tests {
		result := ToSnake(test)
		if result != expected {
			t.Errorf("with '%s', expected '%s', got '%s'", test, expected, result)
		}
	}
}

func TestToPascal(t *testing.T) {
	tests := map[string]string{
		"":           "",
		"1":          "1",
		"12":         "12",
		"one two":    "OneTwo",
		"one  two":   "OneTwo",
		"one_-two":   "OneTwo",
		"abc":        "Abc",
		"abc-xyz":    "AbcXyz",
		"abc_xyz":    "AbcXyz",
		"abcXyz":     "AbcXyz",
		"abc\nxyz":   "AbcXyz",
		"abc\n\nxyz": "AbcXyz",
	}

	for test, expected := range tests {
		result := ToPascal(test)
		if result != expected {
			t.Errorf("with '%s', expected '%s', got '%s'", test, expected, result)
		}
	}
}

func TestToKebab(t *testing.T) {
	tests := map[string]string{
		"":           "",
		"1":          "1",
		"12":         "12",
		"one two":    "one-two",
		"one  two":   "one-two",
		"one_-two":   "one-two",
		"abc":        "abc",
		"abc-xyz":    "abc-xyz",
		"abc_xyz":    "abc-xyz",
		"abcXyz":     "abc-xyz",
		"abc\nxyz":   "abc-xyz",
		"abc\n\nxyz": "abc-xyz",
	}

	for test, expected := range tests {
		result := ToKebab(test)
		if result != expected {
			t.Errorf("with '%s', expected '%s', got '%s'", test, expected, result)
		}
	}
}

func TestNewConverter(t *testing.T) {
	str := "Foo Bar"
	c := NewConverter(str)

	if c.ToCamel() != ToCamel(str) {
		t.Error("Converter returned incorrect value.")
	}
}
