package description

import "strings"

type (
	// LanguageLevel type for language level
	LanguageLevel string
)

// Block which defines all possible LanguageLevel types for Checkout.
const (
	LanguageLevelA1 LanguageLevel = "A1"
	LanguageLevelA2 LanguageLevel = "A2"
	LanguageLevelB1 LanguageLevel = "B1"
	LanguageLevelB2 LanguageLevel = "B2"
	LanguageLevelC1 LanguageLevel = "C1"
	LanguageLevelC2 LanguageLevel = "C2"
)

// languageLevelMap map for validation.
var languageLevelMap = map[LanguageLevel]struct{}{
	LanguageLevelA1: {},
	LanguageLevelA2: {},
	LanguageLevelB1: {},
	LanguageLevelB2: {},
	LanguageLevelC1: {},
	LanguageLevelC2: {},
}

// String - method for casting LanguageLevel to string.
func (l LanguageLevel) String() string {
	return string(l)
}

// Validate - returns true if LanguageLevel is valid.
func (l LanguageLevel) Validate() bool {
	_, ok := languageLevelMap[l]
	return ok
}

// Sanitize - returns sanitized LanguageLevel.
func (l *LanguageLevel) Sanitize() {
	*l = LanguageLevel(strings.ToUpper(l.String()))
}
