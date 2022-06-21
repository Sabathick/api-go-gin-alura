package main

import (
	"github.com/Sabathick/api-go-gin-alura/models"
	"github.com/Sabathick/api-go-gin-alura/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Erick Martins", CPF: "00000000000", RG: "4444444444"},
		{Nome: "Erick Pereira", CPF: "00000000001", RG: "4444444443"},
	}
	routes.HandleRequests()
}
