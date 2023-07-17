package app

import (
	"context"

	"github.com/alex-bodnar/lib/database"
)

// initDatabase init database in app struct.
func (a *App) initDatabase(ctx context.Context) {
	a.db = database.InitDatabase(ctx, a.config.Storage.Postgres, a.logger, a.dbMigrationsFS)
}
