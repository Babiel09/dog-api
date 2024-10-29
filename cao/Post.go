package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PostCao(c *gin.Context) {
	var req models.Caes
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro, verifique se todos os campos foram preenchidos."})
		return

	}
	//Caso não ocorra um erro
	database.DB.Create(&req)
	c.JSON(201, req)
}
