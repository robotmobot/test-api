package database

import (
	"flag"
	"log"
	"math/rand"
	"sync"
	config "test-api/config"
	"test-api/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var once sync.Once

var Migrate = flag.Bool("m", false, "migrates model to db")

func init() {
	flag.Parse()
}
func NewDB() *gorm.DB {
	var err error
	once.Do(func() {
		DB, err = connectDB()
		if err != nil {
			log.Panic(err)
		}
		if *Migrate {
			DB.AutoMigrate(&model.Product{})
			DB.Create(seedDb())

		}
	})

	return DB
}

func connectDB() (*gorm.DB, error) {
	conString := config.GetPostgresConnectionString()

	<-time.After(time.Second * 4)

	return gorm.Open(postgres.Open(conString))
}

func GetDBInstance() *gorm.DB {
	return DB
}

//Trying to randomize seed data
func seedDb() []model.Product {
	randomName := []string{"Car", "Phone", "Pc", "Laptop", "Dress", "Shirt",
		"Perfume", "TV", "Refrigerator", "Oven", "Washing Machine",
		"Dishwashing Machine"}
	products := make([]model.Product, 50)

	for i := range products {
		products[i] = model.Product{Name: randomName[rand.Intn(len(randomName))], Detail: "Detail for " + randomName[rand.Intn(len(randomName))], Price: rand.Float64() * (100 - 2), IsCampaign: rand.Uint64()&(1<<63) == 0}
	}

	return products
}
