package users

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PutUser(c *gin.Context) {
	var req models.User
	id := c.Param("id")

	if id == "" {
		c.JSON(400, req)
		return
	}

	result := database.DB.First(&req, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"server": result})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	update := database.DB.Save(&req)
	if update.Error != nil {
		c.JSON(400, gin.H{"server": update})
		return
	}

	c.JSON(202, req)

}
