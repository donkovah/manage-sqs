package main

import (
	"be/src/app/routes"
	"be/src/domain/service"
	"be/src/infrastructure/config"
	"be/src/infrastructure/persistence"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Load application configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set Gin mode based on config
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	projectService := service.NewProjectService(persistence.NewProjectRepository(config.DB))
	taskService := service.NewTaskService(persistence.NewTaskRepository(config.DB))

	routes.InitRoutes(r, projectService, taskService)

	// Listen and serve on the specified port
	if err := r.Run(fmt.Sprintf(":%d", config.App.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
