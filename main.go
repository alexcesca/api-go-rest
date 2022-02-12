package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandleResquest()
}
