package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(200, user)
}
