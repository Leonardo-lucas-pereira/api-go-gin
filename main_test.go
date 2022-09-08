package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Leonardo-lucas-pereira/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacoes)
	req, _ := http.NewRequest("GET", "/leo", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d eo esperado era %d",
			resposta.Code, http.StatusOK)
	}
}
