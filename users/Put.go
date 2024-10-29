package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PutUser(c *gin.Context) {
	var putUser models.User
	id := c.Params.ByName("id")
	database.DB.First(&putUser, id)
	if err := models.UserValidation(&putUser); err != nil {
		c.JSON(400, gin.H{"server": "Ocorreu um erro para a modificação do usuário, por favor verifique o valor das credênciais."})
		return
	}
	database.DB.Save(&putUser)
	c.JSON(202, putUser)
}
