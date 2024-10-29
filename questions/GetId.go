package questions

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetQuestionsPId(c *gin.Context) {
	var req models.Perguntas
	id := c.Params.ByName("id")
	//Procura no datagbase
	database.DB.First(&req, id)
	//Verifica se não há erros
	if req.Id == 0 {
		c.JSON(404, gin.H{"server": "Unxpected error"})
	}
	//Caso não tenha erros
	c.JSON(200, req)
}
