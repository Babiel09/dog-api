package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro para a criação do usuário, por favor verifique o valor das credênciais."})
		return
	}
	database.DB.Create(&req)
	c.JSON(201, req)
}
