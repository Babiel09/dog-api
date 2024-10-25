package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Babiel09/Gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func RequestPadrao() *gin.Engine {
	routes := gin.Default()
	return routes
}

func VeirificaOStatusCode(t *testing.T) {
	//Defino que o request dos teste via ser igual a função do request padrão
	requestTest := RequestPadrao()
	//Defino a rota que eu quero testar
	requestTest.GET("/cao", controllers.GetCaes)
	//Aqui eu faço a requisição, por enquanto sem verificar o erro
	req, _ := http.NewRequest("GET", "/cao", nil)
	resposta := httptest.NewRecorder()
	requestTest.ServeHTTP(resposta, req)
	if resposta.Code != http.StatusOK {
		t.Fatalf("O teste resultou em %d, mas deveria resultar em %d", resposta.Code, http.StatusOK)
		return
	}
}
