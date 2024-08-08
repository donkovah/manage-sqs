package routes

import (
	"be/src/app/controllers"
	"be/src/domain/repository"
	"be/src/infrastructure/config"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Create repository instance
	projectRepo := repository.NewProjectRepository(config.DB)
	taskRepo := repository.NewTaskRepository(config.DB)

	// Create controller instance with injected repository
	projectController := controllers.NewProjectController(projectRepo)
	taskController := controllers.NewTaskController(taskRepo)

	v1 := r.Group("/v1")

	projectRoute := v1.Group("/projects")
	{
		projectRoute.GET("", projectController.GetProjects)
		projectRoute.GET("/:id", projectController.GetProject)
		projectRoute.POST("", projectController.CreateProject)
		projectRoute.PUT("/:id", projectController.UpdateProject)
		projectRoute.DELETE("/:id", projectController.DeleteProject)
	}

	taskRoute := v1.Group("/tasks")
	{
		taskRoute.GET("/:id", taskController.GetTask)
		taskRoute.GET("/", taskController.GetTasks)
		taskRoute.POST("/", taskController.CreateTask)
		taskRoute.PUT("/:id", taskController.UpdateTask)
		taskRoute.PATCH("/:id/status", taskController.GetTask)
		taskRoute.DELETE("/:id", taskController.DeleteTask)
	}

}
