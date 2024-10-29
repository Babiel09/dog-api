package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func DelteUser(c *gin.Context) {
	var deleteUser models.User
	id := c.Params.ByName("id")
	database.DB.Delete(&deleteUser, id)
	if deleteUser.Id == 0 {
		c.JSON(202, gin.H{"server": "Usuário deletado com sucesso"})
		return
	}
	c.JSON(500, gin.H{"server": "Ocorreu um erro, o usuário não foi deletado, verifique os parâmetros da URL"})
}
