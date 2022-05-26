package main

import (
	"test-api/controller"
	"test-api/database"
	"test-api/gRPC/service"
)

func main() {
	db := database.NewDB()
	pc := controller.NewProductController(db)
	test := service.NewGrpcService(*pc)
	test.GRPCSERVER()
}
