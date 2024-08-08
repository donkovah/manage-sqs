package routes

import (
	"be/src/app/controllers"
	"be/src/domain/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, projectService *service.ProjectService, taskService *service.TaskService) {
	projectController := controllers.NewProjectController(projectService)
	taskController := controllers.NewTaskController(taskService)

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
