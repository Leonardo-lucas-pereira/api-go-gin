package main

import (
	"github.com/Leonardo-lucas-pereira/api-go-gin/database"
	"github.com/Leonardo-lucas-pereira/api-go-gin/models"
	"github.com/Leonardo-lucas-pereira/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Leonardo", CPF: "14258714587", RG: "42005479"},
		{Nome: "Ana", CPF: "14248514587", RG: "92005439"},
	}

	routes.HandleRequests()
}
