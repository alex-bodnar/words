package words

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Word struct represents the word domain model
	Word struct {
		ID            uuid.UUID
		Number        uint64
		Word          string
		Transcription string
		LanguageID    uint64
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}
)
