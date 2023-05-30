package app

import (
	"github.com/alex-bodnar/lib/log"
)

// initLogger initializes logger.
func (a *App) initLogger() {
	a.logger = log.InitLogger(a.config.Logger, map[string]string{
		"service": a.meta.Info.AppName,
	})
}
