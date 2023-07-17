package languages

import "github.com/alex-bodnar/words/internal/api/domain/languages"

// toCreateParams converts a domain Language model to a createParams.
func toCreateParams(lang languages.Language) createParams {
	return createParams{
		LanguageName: lang.LanguageName,
		Code:         lang.Code,
	}
}

// toDomain converts a Language model to a domain Language.
func (l Language) toDomain() languages.Language {
	return languages.Language{
		ID:           l.ID,
		LanguageName: l.LanguageName,
		Code:         l.Code,
		CreatedAt:    l.CreatedAt,
		UpdatedAt:    l.UpdatedAt,
	}
}
