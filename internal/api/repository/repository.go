package repository

import (
	"context"

	"github.com/alex-bodnar/words/internal/api/domain/description"
	"github.com/alex-bodnar/words/internal/api/domain/groups"
	"github.com/alex-bodnar/words/internal/api/domain/languages"
	"github.com/alex-bodnar/words/internal/api/domain/translations"
	"github.com/alex-bodnar/words/internal/api/domain/words"
)

//go:generate mockgen -source repository.go -destination ./repository_mock.go -package repository

type (
	// Description - describe an interface for working with description database models.
	Description interface {
		// Create new description in database.
		Create(ctx context.Context, val description.Description) (description.Description, error)
	}

	// Groups - describe an interface for working with groups database models.
	Groups interface {
		// Create new group in database.
		Create(ctx context.Context, name string) (groups.Group, error)
	}

	// Languages - describe an interface for working with languages database models.
	Languages interface {
		// Create new language in database.
		Create(ctx context.Context, val languages.Language) (languages.Language, error)
	}

	// Translations - describe an interface for working with translations database models.
	Translations interface {
		// Create new translation in database.
		Create(ctx context.Context, val translations.Translation) (translations.Translation, error)
	}

	// Words - describe an interface for working with words database models.
	Words interface {
		// Create new word in database.
		Create(ctx context.Context, val words.Word) (words.Word, error)
	}
)
