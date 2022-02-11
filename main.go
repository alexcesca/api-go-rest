package main

import (
	"fmt"
	"go-rest-api/controllers/models"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConectaBanco()
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "hostiria 1"},
		{Id: 2, Nome: "nome 2", Historia: "hostiria 2"},
	}

	fmt.Println("iniciando servidor com Golang rest")
	routes.HandleRequest()
}
