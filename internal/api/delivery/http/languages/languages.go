package languages

import (
	"github.com/alex-bodnar/lib/errs"
	"github.com/alex-bodnar/lib/http/responder"
	"github.com/gofiber/fiber/v2"

	"github.com/alex-bodnar/words/internal/api/services"
)

// Handler - defines handler for working with languages.
type Handler struct {
	responder.Responder
	service services.LanguagesService
}

// NewHandler - creates new handler for working with languages.
func NewHandler(service services.LanguagesService) *Handler {
	return &Handler{
		service: service,
	}
}

// Create - creates new language.
func (h *Handler) Create(ctx *fiber.Ctx) error {
	var req createLanguageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errs.BadRequest{Cause: "invalid body"}
	}

	lang, err := h.service.Create(ctx.Context(), req.toDomain())
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusCreated, toCreateLanguageResponse(lang))
}
