package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PatchCaoTitle(c *gin.Context) {
	var req models.Caes
	id := c.Param("id")

	if id == "" {
		c.JSON(500, req)
	}
	//Procura as informações do database
	database.DB.First(&req, id)
	//Caso não ocorra um erro
	database.DB.Model(&req).UpdateColumn("title", req.Title)
	c.JSON(202, req)

}

func PatchCaoLegenda(c *gin.Context) {
	var req2 models.Caes
	id := c.Param("id")

	if id == "" {
		c.JSON(500, req2)
	}
	//Procura as informações do database
	database.DB.First(&req2, id)
	//Caso ocorra um erro
	//Caso não ocorra um erro
	database.DB.Model(&req2).UpdateColumn("legenda", req2.Legenda)
	c.JSON(202, req2)

}
