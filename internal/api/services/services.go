package services

import (
	"context"

	"github.com/alex-bodnar/words/internal/api/domain/languages"
)

//go:generate mockgen -source services.go -destination ./services_mock.go -package services

type (
	// LanguagesService -  describe an interface for working with languages.
	LanguagesService interface {
		Create(ctx context.Context, lang languages.Language) (languages.Language, error)
	}
)
