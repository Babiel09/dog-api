package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetCaesPerID(c *gin.Context) {
	var req models.Caes
	id := c.Params.ByName("id")
	//Procura dentro do banco de dados
	database.DB.First(&req, id)
	//Caso ocorra um erro
	if req.Id == 0 {
		c.JSON(400, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return
	}
	//Caso não encontre um erro
	c.JSON(200, req)
}
