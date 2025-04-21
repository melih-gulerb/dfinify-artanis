package handlers

import (
	basemodal "artanis/src/models/base"
	"artanis/src/models/entities"
	"artanis/src/models/requests"
	"artanis/src/repositories/projectUserRepository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProjectUserHandler struct {
	repository projectUserRepository.ProjectUserRepository
}

func NewProjectUserHandler(repository projectUserRepository.ProjectUserRepository) *ProjectUserHandler {
	return &ProjectUserHandler{repository: repository}
}

func (h *ProjectUserHandler) AssignUser(c *fiber.Ctx) error {
	var userRequest requests.AssignUser
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	projectUser := entities.ProjectUser{
		Id:        uuid.New().String(),
		ProjectId: userRequest.ProjectId,
		UserId:    userRequest.UserId,
		RoleId:    userRequest.RoleId,
	}

	if err := h.repository.RegisterProjectUser(projectUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "User successfully assigned",
	})
}
