package router

import (
	"Canteen/controller/login"
	"Canteen/middleware"
	"github.com/gin-gonic/gin"
)

func LoginRegister(c *gin.Engine) {
	api := c.Group("/canteen")
	{
		api.POST("/user/login", login.Auth)
		api.POST("/user/password", middleware.JwtAuth(), login.Password)
	}
}
