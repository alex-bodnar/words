package app

import (
	stdlog "log"

	"github.com/alex-bodnar/words/internal/config"
)

// PopulateConfiguration load configuration from file.
func (a *App) populateConfiguration() {
	var err error

	if a.config, err = config.New(a.meta.Info.AppName, a.meta.ConfigPath); err != nil {
		stdlog.Fatal(err)
	}
}
