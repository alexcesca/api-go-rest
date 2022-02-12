package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	connStr := "user=postgres dbname=personalidades password=.A9l13ed0m7., host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Panic("Erro ao conectar o banco de dados")
	} else {
		log.Println("Banco conectado")
	}
}
