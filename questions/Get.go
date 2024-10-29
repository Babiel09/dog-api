package questions

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	var req []models.Perguntas
	database.DB.Find(&req)
	c.JSON(200, req)
}
