package controllers

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

//Get

func GetUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(200, user)
}

//Ger p/r id

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

//Post

func PostUser(c *gin.Context) {
	var userPost models.User
	if err := models.UserValidation(&userPost); err != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro para a criação do usuário, por favor verifique o valor das credênciais."})
		return
	}
	database.DB.Create(&userPost)
	c.JSON(201, userPost)
}

//Put

func PutUser(c *gin.Context) {
	var putUser models.User
	id := c.Params.ByName("id")
	database.DB.First(&putUser, id)
	if err := models.UserValidation(&putUser); err != nil {
		c.JSON(400, gin.H{"server": "Ocorreu um erro para a modificação do usuário, por favor verifique o valor das credênciais."})
		return
	}
	database.DB.Model(&putUser).Updates(putUser)
	c.JSON(202, putUser)
}

// Delete
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
