package questions

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)


func DeleteQuestion(c *gin.Context) {

	var req models.Perguntas
	id := c.Params.ByName("id")
	database.DB.First(&req, id)
	//Verifica os erros
	if req.Id == 0 {
		c.JSON(404, gin.H{"server": "Unxpected error"})
	}
	//Caso não ocorra nenhum erro
	database.DB.Delete(&req, id)
	c.JSON(202, gin.H{"server": "The question didin't exist more"})
}
