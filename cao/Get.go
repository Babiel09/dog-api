package cao

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetCaes(c *gin.Context) {
	var req []models.Caes
	//PRocura no banco de dados
	procurar := database.DB.Find(&req)
	if procurar.Error != nil {
		c.JSON(400, gin.H{"error": "Ocorreu um erro, verifique os parâmetros do corpo do resposta"})
		return
	}
	//Caso ocorra um erro
	if len(req) == 0 {
		c.JSON(400, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return
	}
	//Caso não ocorra um erro
	c.JSON(200, req)
}
