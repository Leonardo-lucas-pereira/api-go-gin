package database

import (
	"log"

	"github.com/Leonardo-lucas-pereira/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBancoDeDados() {
	//"host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	stringDeConexao := "host=172.20.0.2 user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err := gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o DB")
	}
	DB.AutoMigrate(&models.Aluno{})
}
