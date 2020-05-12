package db

import (
	"github.com/brunotitata/go-api/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConexaoDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error ao carregar variaveis de ambiente: %v", err)
	}

	dsn := os.Getenv("dsn")
	db, err := gorm.Open("postgres", dsn)

	//defer db.Close()

	if err != nil {
		log.Fatalf("Error ao conectar no postgresql: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.Usuario{})
	return db
}
