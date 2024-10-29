package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteCao(c *gin.Context) {
	var req models.Caes
	id := c.Params.ByName("id")
	//Tenta deletar no banco de dados
	database.DB.Delete(&req, id)
	//Caso o cão seja deletado
	if req.Id == 0 {
		c.JSON(202, gin.H{"server": "cão deletado com sucesso"})
		return
	}
	c.JSON(202, gin.H{"server": "Dog is sucefully deleted"})

}
