package database

import (
	"log"
	"sync"
	"time"

	config "test-api/config"
	"test-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

var once *sync.Once

func NewDB() *gorm.DB {
	once.Do(func() {
		var err error
		DB, err = connectDB()
		if err != nil {
			log.Panic(err)
		}

		DB.AutoMigrate(&model.Product{})

	})

	return DB
}

func connectDB() (*gorm.DB, error) {
	conString := config.GetPostgresConnectionString()

	<-time.After(time.Second * 4)

	return gorm.Open(config.GetDBType(), conString)
}

func GetDBInstance() *gorm.DB {
	return DB
}
