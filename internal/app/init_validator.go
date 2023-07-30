package app

import "github.com/go-playground/validator/v10"

// initValidator initializes validator.
func (a *App) initValidator() {
	a.valid = validator.New()
}
