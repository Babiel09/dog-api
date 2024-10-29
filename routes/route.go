package routes

import (
	"github.com/Babiel09/Gin-rest-api/cao"
	"github.com/Babiel09/Gin-rest-api/middlewares"
	"github.com/Babiel09/Gin-rest-api/questions"
	"github.com/Babiel09/Gin-rest-api/users"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	request := gin.Default()

	//Parte de postagem na rede social da API
	request.Use(middlewares.CorsMiddleware())
	request.GET("/dogs", cao.GetCaes)
	request.POST("/dogs", cao.PostCao)
	request.PUT("/dogs/:id", cao.PutCao)
	request.PATCH("/dogs/:id/name", cao.PatchCaoTitle)
	request.PATCH("/dogs/:id/legenda", cao.PatchCaoLegenda)
	request.DELETE("/dogs/:id", cao.DeleteCao)
	request.GET("/dogs/:id", cao.GetCaesPerID)

	//Parte de perguntas da API
	request.GET("/questions", questions.GetQuestions)
	request.GET("/questions/:id", questions.GetQuestionsPId)
	request.POST("questions", questions.PostQuestion)
	request.PUT("questions/:id", questions.PutQuestion)
	request.DELETE("questions/:id", questions.DeleteQuestion)

	//Parte do User da API
	request.GET("/users", users.GetUser)
	request.GET("/users/:id", users.GetUserPerId)
	request.POST("/users", users.PostUser)
	request.PUT("/users/:id", users.PutUser)
	request.DELETE("/users/:id", users.DelteUser)

	request.Run()
}
