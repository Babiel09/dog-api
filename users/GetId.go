package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetUserPerId(c *gin.Context) {
	var userId models.User
	id := c.Params.ByName("id")
	database.DB.First(&userId, id)
	if userId.Id == 0 {
		c.JSON(404, gin.H{"server": "O usuário que procura não existe, verifique se as informações estão corretas por favor."})
		return
	}
	c.JSON(200, userId)
}
