package router

import (
	"Canteen/controller/suggest"
	"Canteen/middleware"
	"github.com/gin-gonic/gin"
)

func SuggestRegister(c *gin.Engine) {
	api := c.Group("/canteen/suggest")
	{
		api.GET("/list", middleware.JwtAuth(), suggest.List)
		api.POST("/save", suggest.Save)
	}
}
