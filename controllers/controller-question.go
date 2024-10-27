package controllers

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

// Get Question
func GetQuestions(c *gin.Context) {
	var questionGet []models.Perguntas
	database.DB.Find(&questionGet)
	c.JSON(200, questionGet)
}

// Get id
func GetQuestionsPId(c *gin.Context) {
	var questionGetId models.Perguntas
	id := c.Params.ByName("id")
	//Procura no datagbase
	database.DB.First(&questionGetId, id)
	//Verifica se não há erros
	if questionGetId.Id == 0 {
		c.JSON(404, gin.H{"server": "Unxpected error"})
	}
	//Caso não tenha erros
	c.JSON(200, questionGetId)
}

// Post
func PostQuestion(c *gin.Context) {
	var postQuestion models.Perguntas
	//Verifica se não existem erros
	if err := c.ShouldBindJSON(&postQuestion); err != nil {
		c.JSON(400, gin.H{"server": "Unxpected error"})
	}
	//Caso ele não encontre nenhum erro
	database.DB.Create(&postQuestion)
	c.JSON(201, postQuestion)
}

// Put

func PutQuestion(c *gin.Context) {
	var putQuestionvar models.Perguntas
	id := c.Params.ByName("id")
	database.DB.First(&putQuestionvar, id)
	//Veirica se não há nenhum erro
	if err := c.ShouldBindJSON(&putQuestionvar); err != nil {
		c.JSON(400, gin.H{"server": "Unxpected error"})
	}
	//Caso não ocorra nenhum erro
	database.DB.Model(&putQuestionvar).Updates(putQuestionvar)

}

// Delete
func DeleteQuestion(c *gin.Context) {

	var deleteQuestionvar models.Perguntas
	id := c.Params.ByName("id")
	database.DB.First(&deleteQuestionvar, id)
	//Verifica os erros
	if deleteQuestionvar.Id == 0 {
		c.JSON(404, gin.H{"server": "Unxpected error"})
	}
	//Caso não ocorra nenhum erro
	database.DB.Delete(&deleteQuestionvar, id)
	c.JSON(202, gin.H{"server": "The question didin't exist more"})
}
