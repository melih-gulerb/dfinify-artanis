package handlers

import (
	"artanis/src/configs"
	"artanis/src/models"
	"artanis/src/models/base"
	"artanis/src/models/requests"
	"artanis/src/models/responses"
	"artanis/src/repositories/definitionRepository"
	"github.com/gofiber/fiber/v2"
)

type DefinitionHandler struct {
	db  *definitionRepository.DefinitionRepository
	cfg *configs.Config
}

func NewDefinitionHandler(db *definitionRepository.DefinitionRepository, cfg *configs.Config) *DefinitionHandler {
	return &DefinitionHandler{db: db, cfg: cfg}
}

func (h *DefinitionHandler) Register(c *fiber.Ctx) error {
	var DefinitionRequest requests.RegisterDefinition
	if err := c.BodyParser(&DefinitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(base.Error{Message: err.Error()})
	}

	definition := models.Definition{
		Name:  DefinitionRequest.Name,
		Value: DefinitionRequest.Value,
	}

	if err := h.db.RegisterDefinition(definition); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to register the definition"})
	}

	return c.Status(fiber.StatusCreated).JSON(base.Response{
		Success: true,
		Message: "definition successfully created",
	})
}

func (h *DefinitionHandler) Paginate(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	offset := c.QueryInt("offset")
	collectionId := c.Params("id")

	Definitions, err := h.db.PaginateDefinitions(collectionId, limit, offset)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to paginate the definitions"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "success",
		Data:    responses.PaginateDefinitionResponse{DefinitionResponse: mapPaginateDefinitionResponse(Definitions), TotalCount: len(Definitions)},
	})
}

func (h *DefinitionHandler) Update(c *fiber.Ctx) error {
	var definitionRequest requests.UpdateDefinition
	if err := c.BodyParser(&definitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(base.Error{Message: err.Error()})
	}

	if err := h.db.UpdateDefinition(definitionRequest.Id, definitionRequest.Name, definitionRequest.Value); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to update the definition"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "definition successfully updated",
	})
}

func (h *DefinitionHandler) Delete(c *fiber.Ctx) error {
	definitionId := c.Params("id")

	if err := h.db.DeleteDefinition(definitionId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to delete the definition"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "definition successfully deleted",
	})
}

func mapPaginateDefinitionResponse(definitions []models.Definition) []responses.DefinitionResponse {
	var DefinitionsResponse []responses.DefinitionResponse
	for _, definition := range definitions {
		DefinitionsResponse = append(DefinitionsResponse, responses.DefinitionResponse{
			Id:    definition.Id,
			Name:  definition.Name,
			Value: definition.Value,
		})
	}
	return DefinitionsResponse
}
