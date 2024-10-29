package questions

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PostQuestion(c *gin.Context) {
	var req models.Perguntas
	//Verifica se não existem erros
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"server": "Unxpected error"})
	}
	//Caso ele não encontre nenhum erro
	database.DB.Create(&req)
	c.JSON(201, req)
}
