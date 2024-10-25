package routes

import (
	"github.com/Babiel09/Gin-rest-api/controllers"
	"github.com/Babiel09/Gin-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	request := gin.Default()
	request.Use(middlewares.CorsMiddleware())
	request.GET("/cao", controllers.GetCaes)
	request.POST("/cao", controllers.PostCao)
	request.PUT("/cao/:id", controllers.PutCao)
	request.PATCH("/cao/:id/nome", controllers.PatchCaoTitle)
	request.PATCH("/cao/:id/legenda", controllers.PatchCaoLegenda)
	request.DELETE("/cao/:id", controllers.DeleteCao)
	request.GET("/cao/:id", controllers.GetCaesPerID)
	request.Run()
}
