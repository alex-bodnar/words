package groups

import "github.com/alex-bodnar/words/internal/api/domain/groups"

// toDomain converts a Group model to a domain Group.
func (g Group) toDomain() groups.Group {
	return groups.Group{
		ID:        g.ID,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}
