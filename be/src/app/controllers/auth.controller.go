package controllers

import (
	"be/src/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (tc AuthController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (tc AuthController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (tc AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
