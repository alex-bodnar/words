package languages

import (
	"time"

	"github.com/alex-bodnar/words/internal/api/domain/languages"
)

type (
	createLanguageRequest struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	createLanguageResponse struct {
		ID           uint64    `json:"id"`
		LanguageName string    `json:"name"`
		Code         string    `json:"code"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

// toDomain converts createLanguageRequest to language domain model.
func (req createLanguageRequest) toDomain() languages.Language {
	return languages.Language{
		LanguageName: req.Name,
		Code:         req.Code,
	}
}

// toCreateLanguageResponse converts language domain model to createLanguageResponse.
func toCreateLanguageResponse(lang languages.Language) createLanguageResponse {
	return createLanguageResponse{
		ID:           lang.ID,
		LanguageName: lang.LanguageName,
		Code:         lang.Code,
		CreatedAt:    lang.CreatedAt,
		UpdatedAt:    lang.UpdatedAt,
	}
}
