package controllers

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

// Get
func GetCaes(c *gin.Context) {
	var getCao []models.Caes
	//PRocura no banco de dados
	database.DB.Find(&getCao)
	//Caso ocorra um erro
	if len(getCao) == 0 {
		c.JSON(400, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return
	}
	//Caso não ocorra um erro
	c.JSON(200, getCao)
}

// Get : id

func GetCaesPerID(c *gin.Context) {
	var getIdCao models.Caes
	id := c.Params.ByName("id")
	//Procura dentro do banco de dados
	database.DB.First(&getIdCao, id)
	//Caso ocorra um erro
	if getIdCao.Id == 0 {
		c.JSON(400, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return
	}
	//Caso não encontre um erro
	c.JSON(200, getIdCao)
}

//Post

func PostCao(c *gin.Context) {
	var postCao models.Caes
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&postCao); err != nil {
		c.JSON(400, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return

	}
	//Caso não ocorra um erro
	database.DB.Create(&postCao)
	c.JSON(201, postCao)
}

//Put

func PutCao(c *gin.Context) {
	var putCao models.Caes
	id := c.Params.ByName("id")
	//Procura no banco de dados o devido id na devida table
	database.DB.First(&putCao, id)
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&putCao); err != nil {
		c.JSON(304, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&putCao).Updates(putCao)
	c.JSON(202, putCao)
}

//Patchs

func PatchCaoNome(c *gin.Context) {
	var patchCao models.Caes
	id := c.Params.ByName("id")
	//Procura as informações do database
	database.DB.First(&patchCao, id)
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&patchCao); err != nil {
		c.JSON(304, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&patchCao).UpdateColumn("nome", patchCao.Nome)
	c.JSON(202, patchCao)

}
func PatchCaoIdade(c *gin.Context) {
	var patchCao models.Caes
	id := c.Params.ByName("id")
	//Procura as informações do database
	database.DB.First(&patchCao, id)
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&patchCao); err != nil {
		c.JSON(304, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&patchCao).UpdateColumn("idade", patchCao.Idade)
	c.JSON(202, patchCao)

}
func PatchCaoPeso(c *gin.Context) {
	var patchCao models.Caes
	id := c.Params.ByName("id")
	//Procura as informações do database
	database.DB.First(&patchCao, id)
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&patchCao); err != nil {
		c.JSON(304, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&patchCao).UpdateColumn("peso", patchCao.Peso)
	c.JSON(202, patchCao)

}
func PatchCaoTemp(c *gin.Context) {
	var patchCao models.Caes
	id := c.Params.ByName("id")
	//Procura as informações do database
	database.DB.First(&patchCao, id)
	//Caso ocorra um erro
	if err := c.ShouldBindJSON(&patchCao); err != nil {
		c.JSON(304, gin.H{"error": "ocorreu um erro, verifique os parâmetros da URL"})
		return

	}
	//Caso não ocorra um erro
	database.DB.Model(&patchCao).UpdateColumn("temperamento", patchCao.Temperamento)
	c.JSON(202, patchCao)
}

//Delete

func DeleteCao(c *gin.Context) {
	var DeleteCao models.Caes
	id := c.Params.ByName("id")
	//Tenta deletar no banco de dados
	database.DB.Delete(&DeleteCao, id)
	//Caso o cão seja deletado
	if DeleteCao.Id == 0 {
		c.JSON(202, gin.H{"server": "cão deletado com sucesso"})
		return
	}

}
