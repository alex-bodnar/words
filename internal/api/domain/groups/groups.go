package groups

import "time"

type (
	// Group struct represents the group domain model
	Group struct {
		ID        uint64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
