package app

import (
	"context"

	"github.com/alex-bodnar/words/internal/api/services/languages"
)

// registerServices register services in app struct.
func (a *App) registerServices(ctx context.Context) {
	a.languagesService = languages.NewService(a.languagesRepo, a.valid, a.logger)
}
