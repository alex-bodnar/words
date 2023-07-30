package app

import "github.com/gofiber/fiber/v2"

func (a *App) registerHTTPRoutes(app *fiber.App) {
	router := app.Group("/v1/words-api")

	// Status
	router.Get("/status", a.statusHTTPHandler.CheckStatus)

	// Languages
	router.Post("/languages", a.languagesHTTPHandler.Create)
}
