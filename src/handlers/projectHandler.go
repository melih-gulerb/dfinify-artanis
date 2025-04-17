package handlers

import (
	"artanis/src/configs"
	"artanis/src/models"
	"artanis/src/models/base"
	"artanis/src/models/clients"
	"artanis/src/models/enums"
	"artanis/src/models/requests"
	"artanis/src/models/responses"
	"artanis/src/repositories/projectRepository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProjectHandler struct {
	db  *projectRepository.ProjectRepository
	cfg *configs.Config
}

func NewProjectHandler(db *projectRepository.ProjectRepository, cfg *configs.Config) *ProjectHandler {
	return &ProjectHandler{db: db, cfg: cfg}
}

func (h *ProjectHandler) Register(c *fiber.Ctx) error {
	var projectRequest requests.RegisterProject
	tokenUser := c.Context().UserValue("user").(*clients.User)
	if err := c.BodyParser(&projectRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(base.Error{Message: err.Error()})
	}

	if tokenUser.OrganizationRole == enums.OrganizationUser {
		return c.Status(fiber.StatusForbidden).JSON(base.Error{Message: "Not enough credentials to create a project"})
	}

	project := models.Project{
		Id:             uuid.New().String(),
		Name:           projectRequest.Name,
		Description:    projectRequest.Description,
		OrganizationId: tokenUser.OrganizationId,
	}

	if err := h.db.RegisterProject(project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to register the project"})
	}

	return c.Status(fiber.StatusCreated).JSON(base.Response{
		Success: true,
		Message: "project successfully created",
	})
}

func (h *ProjectHandler) Paginate(c *fiber.Ctx) error {
	user := c.Context().UserValue("user").(*clients.User)
	limit := c.QueryInt("limit")
	offset := c.QueryInt("offset")

	projects, err := h.db.PaginateProjects(user.OrganizationId, limit, offset)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to register the project"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "success",
		Data:    responses.PaginateProjectResponse{ProjectResponse: mapPaginateResponse(projects), TotalCount: len(projects)},
	})
}

func (h *ProjectHandler) Update(c *fiber.Ctx) error {
	tokenUser := c.Context().UserValue("user").(*clients.User)
	if tokenUser.OrganizationRole == enums.OrganizationUser {
		return c.Status(fiber.StatusForbidden).JSON(base.Error{Message: "Not enough credentials to create a project"})
	}

	var projectRequest requests.UpdateProject
	if err := c.BodyParser(&projectRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(base.Error{Message: err.Error()})
	}
	if err := h.db.UpdateProject(projectRequest.Id, projectRequest.Name, projectRequest.Description); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to update the project"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "project successfully updated",
	})
}

func (h *ProjectHandler) Delete(c *fiber.Ctx) error {
	tokenUser := c.Context().UserValue("user").(*clients.User)
	if tokenUser.OrganizationRole == enums.OrganizationUser {
		return c.Status(fiber.StatusForbidden).JSON(base.Error{Message: "Not enough credentials to create a project"})
	}

	projectId := c.Params("id")

	if err := h.db.DeleteProject(projectId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(base.Error{Message: "Failed to delete the project"})
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "project successfully deleted",
	})
}

func (h *ProjectHandler) GetDashboardCounts(c *fiber.Ctx) error {
	user := c.Context().UserValue("user").(*clients.User)
	counts, err := h.db.GetDashboardCounts(user.OrganizationId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(base.Response{
		Success: true,
		Message: "counts successfully fetched",
		Data:    counts,
	})
}

func mapPaginateResponse(projects []models.Project) []responses.ProjectResponse {
	var projectsResponse []responses.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, responses.ProjectResponse{
			Id:          project.Id,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
		})
	}
	return projectsResponse
}
