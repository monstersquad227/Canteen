package suggest

import (
	"Canteen/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	获取建议列表
*/

func List(c *gin.Context) {
	response := service.SuggestList()
	c.JSON(http.StatusOK, response)
}

/*
	保存建议
*/

func Save(c *gin.Context) {
	type requestBody struct {
		Category string `json:"category"`
		Location string `json:"location"`
		Details  string `json:"details"`
		User     string `json:"user"`
		Contract string `json:"contract"`
		Date     string `json:"date"`
	}
	body := requestBody{}
	c.ShouldBindJSON(&body)
	response := service.SuggestSave(body.Category, body.Location, body.Details, body.User, body.Contract, body.Date)
	c.JSON(http.StatusOK, response)
}
