package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetCaesPerID(c *gin.Context) {
	var req models.Caes
	id := c.Param("id")

	if id == "" {
		c.JSON(400, req)
		return
	}

	result := database.DB.First(&req, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return
	}
	c.JSON(200, req)
}
