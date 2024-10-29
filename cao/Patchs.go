package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PatchCaoTitle(c *gin.Context) {
	var patchCao models.Caes
	id := c.Params.ByName("id")
	//Procura as informações do database
	database.DB.First(&patchCao, id)
	//Caso ocorra um erro
	if err := models.ValidarInformations(&patchCao); err != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro, verifique se todos os campos foram preenchidos."})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&patchCao).UpdateColumn("title", patchCao.Title)
	c.JSON(202, patchCao)

}

func PatchCaoLegenda(c *gin.Context) {
	var patchCao models.Caes
	id := c.Params.ByName("id")
	//Procura as informações do database
	database.DB.First(&patchCao, id)
	//Caso ocorra um erro
	if err := models.ValidarInformations(&patchCao); err != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro, verifique se todos os campos foram preenchidos."})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&patchCao).UpdateColumn("legenda", patchCao.Legenda)
	c.JSON(202, patchCao)

}
