package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func DelteUser(c *gin.Context) {
	var req models.User
	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(400, req)
		return
	}
	delete := database.DB.Delete(&req, id)
	if delete.Error != nil {
		c.JSON(400, delete)
		return
	}

	c.JSON(202, gin.H{"server": "Dog is sucefully deleted"})
}
