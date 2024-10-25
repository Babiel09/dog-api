package database

import (
	"log"
	"time"

	"github.com/Babiel09/Gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectarComBase() {
	log.Println("Conectando com o banco de dados, isso pode levar alguns segundos")
	time.Sleep(20 * time.Second)
	stringDeConexao := "host=go_db2 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Falha ao conectar com o banco de dados")
	}
	//Criando a tablea com o gorm:
	DB.AutoMigrate(models.Caes{})
	DB.AutoMigrate(models.Perguntas{})
	log.Println("Conexão com o banco de dados estabelecia")
}
