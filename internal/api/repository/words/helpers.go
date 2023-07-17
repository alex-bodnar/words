package words

import (
	"github.com/alex-bodnar/words/internal/api/domain/words"
	"github.com/alex-bodnar/words/internal/api/repository"
)

// toCreateParams converts a domain Word model to a createParams.
func toCreateParams(word words.Word) createParams {
	return createParams{
		Word:          word.Word,
		Transcription: repository.ToPgText(word.Transcription),
		LanguageID:    word.LanguageID,
	}
}

// toDomain converts a Word model to a domain Word.
func (w Word) toDomain() words.Word {
	return words.Word{
		ID:            w.ID,
		Word:          w.Word,
		Transcription: repository.FromPgText(w.Transcription),
		LanguageID:    w.LanguageID,
		CreatedAt:     w.CreatedAt,
		UpdatedAt:     w.UpdatedAt,
	}
}
