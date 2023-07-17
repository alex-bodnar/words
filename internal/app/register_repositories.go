package app

import (
	"github.com/alex-bodnar/words/internal/api/repository"
	"github.com/alex-bodnar/words/internal/api/repository/description"
	"github.com/alex-bodnar/words/internal/api/repository/groups"
	"github.com/alex-bodnar/words/internal/api/repository/languages"
	"github.com/alex-bodnar/words/internal/api/repository/translations"
	"github.com/alex-bodnar/words/internal/api/repository/words"
)

func (a *App) registerRepositories() {
	a.txRepo = repository.NewTxRepo(a.db)
	a.descriptionRepo = description.NewRepository(a.db)
	a.groupsRepo = groups.NewRepository(a.db)
	a.languagesRepo = languages.NewRepository(a.db)
	a.translationsRepo = translations.NewRepository(a.db)
	a.wordsRepo = words.NewRepository(a.db)
}
