package caseconv

import (
	"bytes"
	"strings"
	"unicode"
)

// A Converter allows initialising with a string and provides
// various methods to convert to specific cases.
type Converter struct {
	str string
}

// Convert initialised string to Camel Case.
// Example: 'foo bar' -> 'fooBar'
func (c *Converter) ToCamel() string {
	return ToCamel(c.str)
}

// Convert initialised string to Snake Case.
// Example: 'foo bar' -> 'foo_bar'
func (c *Converter) ToSnake() string {
	return ToSnake(c.str)
}

// Convert initialised string to Pascal Case.
// Example: 'foo bar' -> 'FooBar'
func (c *Converter) ToPascal() string {
	return ToPascal(c.str)
}

// Convert initialised string to Kebab Case.
// Example: 'foo bar' -> 'foo-bar'
func (c *Converter) ToKebab() string {
	return ToKebab(c.str)
}

// Convert s to Camel Case.
// Example: 'foo bar' -> 'fooBar'
func ToCamel(s string) string {
	var result bytes.Buffer

	passedDelim := false

	for _, r := range s {
		if isDelim(r) {
			passedDelim = true
		} else {
			if passedDelim {
				r = unicode.ToUpper(r)
				passedDelim = false
			}
			result.WriteRune(r)
		}
	}

	return result.String()
}

// Convert s to Snake Case.
// Example: 'foo bar' -> 'foo_bar'
func ToSnake(s string) string {
	return lowerWithSeparator(s, '_')
}

// Convert s to Pascal Case.
// Example: 'foo bar' -> 'FooBar'
func ToPascal(s string) string {
	return strings.Title(ToCamel(s))
}

// Convert s to Kebab Case.
// Example: 'foo bar' -> 'foo-bar'
func ToKebab(s string) string {
	return lowerWithSeparator(s, '-')
}

func NewConverter(s string) Converter {
	return Converter{
		str: s,
	}
}

func isDelim(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsSymbol(r) || r == '_' || r == '-'
}

func lowerWithSeparator(s string, sep rune) string {
	var result bytes.Buffer

	for i, r := range s {
		if isDelim(r) {
			result.WriteRune(sep)
		} else if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune(sep)
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
