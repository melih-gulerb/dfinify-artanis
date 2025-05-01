package handlers

import (
	"artanis/src/clients"
	basemodal "artanis/src/models/base"
	"artanis/src/models/entities"
	"artanis/src/models/enums"
	"artanis/src/models/requests"
	"artanis/src/repositories/projectUserRepository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProjectUserHandler struct {
	repository   projectUserRepository.ProjectUserRepository
	divineShield clients.DivineShieldClient
}

func NewProjectUserHandler(repository projectUserRepository.ProjectUserRepository) *ProjectUserHandler {
	return &ProjectUserHandler{repository: repository}
}

func (h *ProjectUserHandler) Register(c *fiber.Ctx) error {
	var userRequest requests.Register
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	projectUser := entities.ProjectUser{
		Id:        uuid.New().String(),
		ProjectId: userRequest.ProjectId,
		UserId:    userRequest.UserId,
		Role:      userRequest.Role,
	}

	if err := h.repository.RegisterProjectUser(projectUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "User successfully assigned",
	})
}

func (h *ProjectUserHandler) UpdateProjectUserRole(c *fiber.Ctx) error {
	projectUserId := c.Params("id")
	role := c.QueryInt("role")

	if err := h.repository.UpdateProjectUserRole(projectUserId, enums.ProjectRole(role)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "User successfully assigned",
	})
}

func (h *ProjectUserHandler) Delete(c *fiber.Ctx) error {
	projectUserId := c.Params("id")

	if err := h.repository.DeleteProjectUser(projectUserId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(basemodal.Error{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "User successfully assigned",
	})
}

func (h *ProjectUserHandler) Paginate(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	offset := c.QueryInt("offset")
	projectId := c.Params("id")

	projectUsers := h.repository.Paginate(projectId, limit, offset)

	return c.Status(fiber.StatusOK).JSON(basemodal.Response{
		Success: true,
		Message: "User successfully assigned",
		Data:    projectUsers,
	})
}
