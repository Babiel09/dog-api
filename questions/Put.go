package questions

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func PutQuestion(c *gin.Context) {
	var req models.Perguntas
	id := c.Params.ByName("id")
	database.DB.First(&req, id)
	//Veirica se não há nenhum erro
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"server": "Unxpected error"})
	}
	//Caso não ocorra nenhum erro
	database.DB.Save(&req)

}
