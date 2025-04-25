package handlers

import (
	"artanis/src/configs"
	basemodal "artanis/src/models/base"
	clientmodal "artanis/src/models/clients"
	"artanis/src/models/entities"
	"artanis/src/models/enums"
	"artanis/src/models/requests"
	"artanis/src/models/responses"
	models "artanis/src/models/services"
	"artanis/src/repositories/definitionRepository"
	"artanis/src/repositories/projectUserRepository"
	"artanis/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	var definitionRequest requests.RegisterDefinition
	if err := c.BodyParser(&definitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	definition := entities.Definition{
		Id:           uuid.New().String(),
		CollectionId: definitionRequest.CollectionId,
		Name:         definitionRequest.Name,
		Value:        definitionRequest.Value,
	}

	if err := h.db.RegisterDefinition(definition); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to register the definition"})
	}

	return c.Status(fiber.StatusCreated).JSON(basemodal.Response{
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
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to paginate the definitions"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "success",
		Data:    responses.PaginateDefinitionResponse{DefinitionResponse: mapPaginateDefinitionResponse(definitions), TotalCount: len(definitions)},
	})
}

func (h *DefinitionHandler) UpdateName(c *fiber.Ctx) error {
	var definitionRequest requests.UpdateDefinitionName
	if err := c.BodyParser(&definitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	user := c.Context().UserValue("user").(*clientmodal.User)

	projectUser := h.pdb.GetProjectUser(user.Id, definitionRequest.ProjectId)
	if projectUser == nil || *projectUser == enums.ProjectUser {
		return c.Status(fiber.StatusForbidden).JSON(basemodal.Error{Message: "not enough credentials to create a definition"})
	}

	definition := h.db.GetDefinition(definitionRequest.Id)
	if definition == nil {
		return c.Status(fiber.StatusNotFound).JSON(basemodal.Error{Message: "Definition not found"})
	}

	if err := h.db.UpdateDefinitionName(definitionRequest.Id, definitionRequest.Name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to update the definition"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "definition name successfully updated",
	})
}

func (h *DefinitionHandler) UpdateValue(c *fiber.Ctx) error {
	var definitionRequest requests.UpdateDefinitionValue
	if err := c.BodyParser(&definitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	definition := h.db.GetDefinitionDetail(definitionRequest.Id)
	if definition == nil {
		return c.Status(fiber.StatusNotFound).JSON(basemodal.Error{Message: "Definition not found"})
	}

	user := c.Context().UserValue("user").(*clientmodal.User)

	projectUser := h.pdb.GetProjectUser(user.Id, definition.ProjectId)
	if projectUser == nil {
		return c.Status(fiber.StatusForbidden).JSON(basemodal.Error{Message: "not enough credentials to create a definition"})
	}

	if *projectUser == enums.ProjectUser {
		err := h.registerDefinitionChange(definition, definitionRequest.Value, user.Email, user.Id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to register the definition change"})
		}
		return c.Status(fiber.StatusOK).JSON(basemodal.Response{
			Success: true,
			Message: "definition value successfully submitted for approval",
		})
	}

	if err := h.db.UpdateDefinitionValue(definitionRequest.Id, definitionRequest.Value); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to update the definition"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "definition name successfully updated",
	})
}

func (h *DefinitionHandler) Delete(c *fiber.Ctx) error {
	definitionId := c.Params("id")

	if err := h.db.DeleteDefinition(definitionId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to delete the definition"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "definition successfully deleted",
	})
}

func mapPaginateDefinitionResponse(definitions []entities.Definition) []responses.DefinitionResponse {
	var definitionsResponse []responses.DefinitionResponse
	for _, definition := range definitions {
		definitionsResponse = append(definitionsResponse, responses.DefinitionResponse{
			Id:        definition.Id,
			Name:      definition.Name,
			Value:     definition.Value,
			CreatedAt: definition.CreatedAt,
		})
	}
	return definitionsResponse
}

func (h *DefinitionHandler) registerDefinitionChange(definitionDetail *models.DefinitionDetail, newValue, userMail, userId string) error {
	slackChannelIds := h.pdb.GetProjectAdminsForSlackUser(definitionDetail.ProjectId)

	request := models.RegisterDefinitionChange{
		DefinitionId:    definitionDetail.Id,
		OldValue:        definitionDetail.OldValue,
		NewValue:        newValue,
		ProjectName:     definitionDetail.ProjectName,
		CollectionName:  definitionDetail.CollectionName,
		DefinitionName:  definitionDetail.Name,
		UserMail:        userMail,
		UserId:          userId,
		SlackChannelIds: slackChannelIds,
	}

	err := h.ds.Register(request)
	if err != nil {
		return err
	}

	return nil
}
