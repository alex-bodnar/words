package translations

import (
	"github.com/alex-bodnar/words/internal/api/domain/translations"
	"github.com/alex-bodnar/words/internal/api/repository"
)

// toCreateParams converts a domain Translation model to a createParams.
func toCreateParams(val translations.Translation) createParams {
	return createParams{
		Translation:   val.Translation,
		Transcription: repository.ToPgText(val.Transcription),
		LanguageID:    val.LanguageID,
	}
}

// toDomain converts a Translation model to a domain Translation.
func (t Translation) toDomain() translations.Translation {
	return translations.Translation{
		ID:            t.ID,
		Translation:   t.Translation,
		Transcription: repository.FromPgText(t.Transcription),
		LanguageID:    t.LanguageID,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}
