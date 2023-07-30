package languages

import (
	"context"

	"github.com/alex-bodnar/lib/log"
	"github.com/go-playground/validator/v10"

	"github.com/alex-bodnar/words/internal/api/domain/languages"
	"github.com/alex-bodnar/words/internal/api/repository"
	"github.com/alex-bodnar/words/internal/api/services"
)

var _ services.LanguagesService = &Service{}

// Service -  defines service for working with languages.
type Service struct {
	langRepo repository.Languages

	valid  *validator.Validate
	logger log.Logger
}

// NewService - creates new service for working with languages.
func NewService(langRepo repository.Languages, valid *validator.Validate, logger log.Logger) *Service {
	return &Service{
		langRepo: langRepo,
		valid:    valid,
		logger:   logger,
	}
}

// Create - creates new language.
func (s *Service) Create(ctx context.Context, lang languages.Language) (languages.Language, error) {
	lang.Sanitize()

	if err := lang.Validate(s.valid); err != nil {
		return languages.Language{}, err
	}

	return s.langRepo.Create(ctx, lang)
}
