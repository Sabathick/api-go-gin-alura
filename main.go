package main

import (
	"github.com/Sabathick/api-go-gin-alura/database"
	"github.com/Sabathick/api-go-gin-alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
