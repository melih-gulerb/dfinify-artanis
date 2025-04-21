package handlers

import (
	"artanis/src/configs"
	basemodal "artanis/src/models/base"
	clientmodal "artanis/src/models/clients"
	"artanis/src/models/entities"
	"artanis/src/models/enums"
	"artanis/src/models/requests"
	"artanis/src/models/responses"
	"artanis/src/repositories/collectionRepository"
	"artanis/src/repositories/projectUserRepository"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CollectionHandler struct {
	db  *collectionRepository.CollectionRepository
	pDb projectUserRepository.ProjectUserRepository
	cfg *configs.Config
}

func NewCollectionHandler(db *collectionRepository.CollectionRepository, pDb projectUserRepository.ProjectUserRepository,
	cfg *configs.Config) *CollectionHandler {
	return &CollectionHandler{db: db, pDb: pDb, cfg: cfg}
}

func (h *CollectionHandler) Register(c *fiber.Ctx) error {
	var collectionRequest requests.RegisterCollection
	if err := c.BodyParser(&collectionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	user := c.Context().UserValue("user").(*clientmodal.User)
	if validateAuth := h.validateAuth(user.Id, collectionRequest.ProjectId); validateAuth != nil {
		return c.Status(fiber.StatusForbidden).JSON(basemodal.Error{Message: validateAuth.Error()})
	}

	collection := entities.Collection{
		Id:          uuid.New().String(),
		Name:        collectionRequest.Name,
		Description: collectionRequest.Description,
		ProjectId:   collectionRequest.ProjectId,
	}

	if err := h.db.RegisterCollection(collection); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to register the collection"})
	}

	return c.Status(fiber.StatusCreated).JSON(basemodal.Response{
		Success: true,
		Message: "collection successfully created",
	})
}

func (h *CollectionHandler) Paginate(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	offset := c.QueryInt("offset")
	projectId := c.Params("id")

	user := c.Context().UserValue("user").(*clientmodal.User)
	if validateAuth := h.pDb.GetProjectUser(user.Id, projectId); validateAuth == nil {
		return c.Status(fiber.StatusForbidden).JSON(basemodal.Error{Message: "not enough credentials to create a collection"})
	}

	collections, err := h.db.PaginateCollections(projectId, limit, offset)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to register the collection"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "success",
		Data:    responses.PaginateCollectionResponse{CollectionResponse: mapPaginateCollectionResponse(collections), TotalCount: len(collections)},
	})
}

func (h *CollectionHandler) Update(c *fiber.Ctx) error {
	var collectionRequest requests.UpdateCollection
	if err := c.BodyParser(&collectionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	user := c.Context().UserValue("user").(*clientmodal.User)
	if validateAuth := h.validateAuth(user.Id, collectionRequest.ProjectId); validateAuth != nil {
		return c.Status(fiber.StatusForbidden).JSON(basemodal.Error{Message: validateAuth.Error()})
	}

	if err := h.db.UpdateCollection(collectionRequest.Id, collectionRequest.Name, collectionRequest.Description); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to update the collection"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "collection successfully updated",
	})
}

func (h *CollectionHandler) Delete(c *fiber.Ctx) error {
	collectionId := c.Params("id")
	projectId := c.Query("projectId")

	user := c.Context().UserValue("user").(*clientmodal.User)
	if validateAuth := h.validateAuth(user.Id, projectId); validateAuth != nil {
		return c.Status(fiber.StatusForbidden).JSON(basemodal.Error{Message: validateAuth.Error()})
	}

	if err := h.db.DeleteCollection(collectionId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(basemodal.Error{Message: "Failed to delete the collection"})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "collection successfully deleted",
	})
}

func mapPaginateCollectionResponse(collections []entities.Collection) []responses.CollectionResponse {
	var collectionsResponse []responses.CollectionResponse
	for _, collection := range collections {
		collectionsResponse = append(collectionsResponse, responses.CollectionResponse{
			Id:          collection.Id,
			Name:        collection.Name,
			Description: collection.Description,
			CreatedAt:   collection.CreatedAt,
		})
	}
	return collectionsResponse
}

func (h *CollectionHandler) validateAuth(userId, projectId string) error {
	projectUser := h.pDb.GetProjectUser(userId, projectId)
	if projectUser == nil || *projectUser == enums.ProjectUser {
		return errors.New("not enough credentials to create a collection")
	}

	return nil
}
