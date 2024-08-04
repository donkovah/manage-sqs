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
	// Create controller instance with injected repository
	projectController := controllers.NewProjectController(projectRepo)

	v1 := r.Group("/v1")

	projectRoute := v1.Group("/projects")
	{
		// Define routes and attach controller methods
		projectRoute.GET("", projectController.GetProducts)
		projectRoute.GET("/:id", projectController.GetProduct)
		projectRoute.POST("", projectController.CreateProject)
		projectRoute.PUT("/:id", projectController.UpdateProduct)
		projectRoute.DELETE("/:id", projectController.DeleteProduct)
	}

	taskRoute := v1.Group("/task")
	{
		taskRoute.GET("/:id", controllers.GetTask)

		taskRoute.GET("/", controllers.GetTasks)

		taskRoute.POST("/", controllers.CreateTask)

		taskRoute.PATCH("/:id", controllers.UpdateTask)

		taskRoute.DELETE("/:id", controllers.DeleteTask)
	}

}
