package translations

import (
	"time"

	"github.com/gofrs/uuid"
)

type (
	// Translation struct represents the translation domain model
	Translation struct {
		ID          uuid.UUID
		Number      uint64
		Translation string
		LanguageID  uint64
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
