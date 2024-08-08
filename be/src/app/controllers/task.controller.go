package controllers

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	repo repository.TaskRepository
}

func NewTaskController(repo repository.TaskRepository) *TaskController {
	return &TaskController{repo: repo}
}
func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.repo.GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
	}
	c.JSON(http.StatusOK, task)
}

func (tc TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.repo.GetTasks(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch task"})
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := tc.repo.CreateTask(context.Background(), &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
	}

	c.JSON(http.StatusOK, createdTask)
}

func (tc TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var taskBody *models.Task

	if err := c.ShouldBindJSON(&taskBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.repo.GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
	}
	task.Title = taskBody.Title
	task.Status = taskBody.Status
	task.Description = taskBody.Description
	task.Deadline = taskBody.Deadline

	updatedTask, err := tc.repo.UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (ts TaskController) DeleteTask(c *gin.Context) {
	id := c.Param(("id"))
	err := ts.repo.DeleteTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
	}
	c.JSON(http.StatusNoContent, nil)
}
