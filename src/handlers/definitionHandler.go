package handlers

import (
	"artanis/src/configs"
	"artanis/src/models/base"
	"artanis/src/models/clients"
	"artanis/src/models/entities"
	"artanis/src/models/enums"
	"artanis/src/models/requests"
	"artanis/src/models/responses"
	models "artanis/src/models/services"
	"artanis/src/repositories/definitionRepository"
	"artanis/src/repositories/projectUserRepository"
	"artanis/src/services"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type DefinitionHandler struct {
	db  *definitionRepository.DefinitionRepository
	pdb *projectUserRepository.ProjectUserRepository
	ds  *services.DefinitionChangeService
	cfg *configs.Config
}

func NewDefinitionHandler(db *definitionRepository.DefinitionRepository, pdb *projectUserRepository.ProjectUserRepository,
	ds *services.DefinitionChangeService, cfg *configs.Config) *DefinitionHandler {
	return &DefinitionHandler{db: db, pdb: pdb, ds: ds, cfg: cfg}
}

func (h *DefinitionHandler) Register(c *fiber.Ctx) error {
	var DefinitionRequest requests.RegisterDefinition
	if err := c.BodyParser(&DefinitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(base.Error{Message: err.Error()})
	}

	definition := entities.Definition{
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

	definitions, err := h.db.PaginateDefinitions(collectionId, limit, offset)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to paginate the definitions"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "success",
		Data:    responses.PaginateDefinitionResponse{DefinitionResponse: mapPaginateDefinitionResponse(definitions), TotalCount: len(definitions)},
	})
}

func (h *DefinitionHandler) Update(c *fiber.Ctx) error {
	var definitionRequest requests.UpdateDefinition
	if err := c.BodyParser(&definitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(base.Error{Message: err.Error()})
	}

	user := c.Context().UserValue("user").(*clients.User)
	role, err := h.validateAuth(user.Id, definitionRequest.ProjectId)
	if err != nil {
		return err
	}

	definition := h.db.GetDefinition(definitionRequest.Id)
	if definition == nil {
		return c.Status(fiber.StatusNotFound).JSON(base.Error{Message: "Definition not found"})
	}

	slackChannelIds := h.pdb.GetProjectAdminsForSlackUser(definitionRequest.ProjectId)

	if role == enums.ProjectUser {

		definitionChangeRequest := models.RegisterDefinitionChange{
			DefinitionId:    definition.Id,
			ProjectId:       definitionRequest.ProjectId,
			OldValue:        definition.Value,
			NewValue:        definitionRequest.Value,
			ProjectName:     definitionRequest.ProjectName,
			CollectionName:  definitionRequest.CollectionName,
			DefinitionName:  definitionRequest.Name,
			UserName:        user.Name,
			UserMail:        user.Email,
			UserId:          user.Id,
			SlackChannelIds: slackChannelIds,
		}
		if err = h.ds.Register(definitionChangeRequest); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to register the definition change"})
		}

		return c.Status(fiber.StatusOK).JSON(base.Response{
			Success: true,
			Message: "Definition update change submitted for approval. You will receive an email when it is approved.",
			Data: responses.DefinitionResponse{
				Id:    definitionRequest.Id,
				Name:  definitionRequest.Name,
				Value: definitionRequest.Value,
			},
		})
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

func mapPaginateDefinitionResponse(definitions []entities.Definition) []responses.DefinitionResponse {
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

func (h *DefinitionHandler) validateAuth(userId, projectId string) (enums.ProjectRole, error) {
	projectUser := h.pdb.GetProjectUser(userId, projectId)
	if projectUser == nil {
		return 0, errors.New("not enough credentials to create a collection")
	}

	return *projectUser, nil
}
