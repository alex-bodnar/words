package languages

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	// Language struct represents the language domain model
	Language struct {
		ID           uint64
		LanguageName string `validate:"required"`
		Code         string `validate:"bcp47_language_tag"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)

// Sanitize sanitizes the language.
func (l *Language) Sanitize() {
	l.LanguageName = strings.ToLower(strings.TrimSpace(l.LanguageName))
	l.Code = strings.ToLower(strings.TrimSpace(l.Code))
}

// Validate validates the language.
func (l Language) Validate(valid *validator.Validate) error {
	return valid.Struct(l)
}
