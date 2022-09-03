package main

import (
	"github.com/Leonardo-lucas-pereira/api-go-gin/database"
	"github.com/Leonardo-lucas-pereira/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
