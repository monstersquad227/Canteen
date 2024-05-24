package login

import (
	"Canteen/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	数据库: 登录授权
*/

func Auth(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.ShouldBindJSON(&requestBody)
	response := service.LoginAuth(requestBody.Username, requestBody.Password)
	c.JSON(http.StatusOK, response)
}

/*
	数据库: 修改密码
*/

func Password(c *gin.Context) {
	var requestBody struct {
		Username           string `json:"username"`
		CurrentPassword    string `json:"currentPassword"`
		NewPassword        string `json:"newPassword"`
		ConfirmNewPassword string `json:"confirmNewPassword"`
	}
	c.ShouldBindJSON(&requestBody)
	response := service.LoginPassword(requestBody.Username, requestBody.CurrentPassword, requestBody.NewPassword)
	c.JSON(http.StatusOK, response)
}
