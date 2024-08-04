package main

import (
	"be/src/app/routes"
	"be/src/infrastructure/config"
	"fmt"

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
	config.LoadConfig()

	r := gin.Default()
	routes.InitRoutes(r)

	// Listen and serve on the specified port
	if err := r.Run(fmt.Sprintf(":%d", config.App.Port)); err != nil {
		panic(err)
	}
}
