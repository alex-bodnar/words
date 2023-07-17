package translations

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Translation struct represents the translation domain model
	Translation struct {
		ID            uuid.UUID
		Number        uint64
		Translation   string
		Transcription string
		LanguageID    uint64
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
)
