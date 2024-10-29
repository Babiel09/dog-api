package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteCao(c *gin.Context) {
	var req models.Caes
	id := c.Param("id")

	if id == "" {
		c.JSON(400, req)
		return
	}
	//Tenta deletar no banco de dados
	delete := database.DB.Delete(&req, id)
	if delete.Error != nil {
		c.JSON(400, delete)
		return
	}
	c.JSON(202, gin.H{"server": "Dog is sucefully deleted"})

}
