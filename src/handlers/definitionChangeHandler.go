package handlers

import (
	basemodal "artanis/src/models/base"
	"artanis/src/models/requests"
	"artanis/src/repositories/definitionChangeRepository"
	"github.com/gofiber/fiber/v2"
)

type DefinitionChangeHandler struct {
	db *definitionChangeRepository.DefinitionChangeRepository
}

func NewDefinitionChangeHandler(db *definitionChangeRepository.DefinitionChangeRepository) *DefinitionChangeHandler {
	return &DefinitionChangeHandler{db: db}
}

func (h *DefinitionChangeHandler) Paginate(c *fiber.Ctx) error {
	var definitionRequest requests.UpdateDefinitionChange
	if err := c.BodyParser(&definitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	if err := h.db.UpdateDefinitionChangeState(definitionRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to update the definition change"})
	}

	return c.Status(fiber.StatusCreated).JSON(basemodal.Response{
		Success: true,
		Message: "definition successfully updated",
	})
}
