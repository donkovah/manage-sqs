package controllers

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	repo repository.ProjectRepository
}

func NewProjectController(repo repository.ProjectRepository) *ProjectController {
	return &ProjectController{repo: repo}
}

func (pc *ProjectController) GetProject(c *gin.Context) {
	id := c.Param("id")

	project, err := pc.repo.GetProject(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) GetProjects(c *gin.Context) {
	projects, err := pc.repo.GetProjects(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)

}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProject, err := pc.repo.CreateProject(context.Background(), &project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdProject)
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	idStr := c.Param("id") // ID from URL parameters

	// Convert string ID to uint64
	productID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Convert uint64 to uint
	id := uint(productID)

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project.ID = id // Set the ID for the update

	updatedProject, err := pc.repo.UpdateProject(context.Background(), &project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedProject)
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	if err := pc.repo.DeleteProject(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
