package database

import (
	"log"
	"time"

	config "test-api/config"
	"test-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)
	<-time.After(time.Second * 4)
	db, err := gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&model.Product{})

	return db
}

func GetDBInstance() *gorm.DB {
	return DB
}
