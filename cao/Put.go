package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PutCao(c *gin.Context) {
	var req models.Caes
	id := c.Params.ByName("id")
	//Procura no banco de dados o devido id na devida table
	database.DB.First(&req, id)
	//Caso ocorra um erro
	if err := models.ValidarInformations(&req); err != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro, verifique se todos os campos foram preenchidos."})
		return

	}
	//Caso não ocorra um erro
	database.DB.Save(&req)
	c.JSON(202, req)
}
