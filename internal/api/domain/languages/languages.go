package languages

import (
	"time"
)

type (
	// Language struct represents the language domain model
	Language struct {
		ID           uint64
		LanguageName string
		Code         string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
