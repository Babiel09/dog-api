package routes

import (
	"github.com/Babiel09/Gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	request := gin.Default()
	request.GET("/cao", controllers.GetCaes)
	request.POST("/cao", controllers.PostCao)
	request.PUT("/cao/:id", controllers.PutCao)
	request.PATCH("/cao/:id", controllers.PatchCaoNome)
	request.PATCH("/cao/:id", controllers.PatchCaoIdade)
	request.PATCH("/cao/:id", controllers.PatchCaoPeso)
	request.PATCH("/cao/:id", controllers.PatchCaoTemp)
	request.PATCH("/cao/:id", controllers.DeleteCao)
	request.GET("/cao/:id", controllers.GetCaesPerID)
	request.Run()
}
