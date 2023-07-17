package description

import "github.com/alex-bodnar/words/internal/api/domain/description"

// toCreateParams converts a domain Description model to a createParams.
func toCreateParams(desc description.Description) createParams {
	return createParams{
		WordID:      desc.WordID,
		Description: desc.Description,
		Level:       LanguageLevel(desc.Level),
	}
}

// toDomain converts a Description model to a domain Description.
func (d Description) toDomain() description.Description {
	return description.Description{
		ID:          d.ID,
		Number:      d.Number,
		WordID:      d.WordID,
		Description: d.Description,
		Level:       description.LanguageLevel(d.Level),
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
