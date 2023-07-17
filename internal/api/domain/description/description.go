package description

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Description struct represents the description domain model
	Description struct {
		ID          uuid.UUID
		Number      uint64
		WordID      uuid.UUID
		Description string
		Level       LanguageLevel
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
