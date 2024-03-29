package app

import (
	"context"
	"embed"

	"github.com/alex-bodnar/lib/http/responder"
	"github.com/alex-bodnar/lib/log"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/alex-bodnar/words/internal/api/delivery"
	"github.com/alex-bodnar/words/internal/api/repository"
	"github.com/alex-bodnar/words/internal/api/services"
	"github.com/alex-bodnar/words/internal/config"
)

type (
	// Meta defines meta for application.
	Meta struct {
		Info       Info
		ConfigPath string
	}

	// Info defines metadata of application.
	Info struct {
		AppName       string
		Tag           string
		Version       string
		Commit        string
		Date          string
		FortuneCookie string
	}

	// App defines main application struct.
	App struct {
		// meta information about application.
		meta Meta

		// tech dependencies.
		config *config.Config
		logger log.Logger
		valid  *validator.Validate

		dbMigrationsFS embed.FS
		db             *pgxpool.Pool

		responder responder.Responder

		// Repository dependencies.
		txRepo           repository.TX
		descriptionRepo  repository.Description
		groupsRepo       repository.Groups
		languagesRepo    repository.Languages
		translationsRepo repository.Translations
		wordsRepo        repository.Words

		// Service dependencies.
		languagesService services.LanguagesService

		// Delivery dependencies.
		statusHTTPHandler    delivery.StatusHTTP
		languagesHTTPHandler delivery.LanguagesHTTP
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta) *App {
	return &App{
		meta: meta,
	}
}

// WithMigrationFS is a setup for database migration filesystem
func (a *App) WithMigrationFS(f embed.FS) *App {
	a.dbMigrationsFS = f
	return a
}

// Run – registers graceful shutdown.
// populate configuration and application dependencies,
// run workers.
func (a *App) Run(ctx context.Context) {
	// Initialize configuration
	a.populateConfiguration()

	// Register Dependencies
	a.initLogger()
	a.initValidator()
	a.initDatabase(ctx)

	// Domain registration.
	a.registerRepositories()
	a.registerServices(ctx)

	// Register Handlers
	a.registerHTTPHandlers()

	// Run Workers
	a.runWorkers(ctx)
}
