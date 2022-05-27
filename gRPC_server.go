package main

import (
	"test-api/controller"
	"test-api/database"
	"test-api/handler/gRPC"
)

func main() {
	db := database.NewDB()
	pc := controller.NewProductController(db)
	test := gRPC.NewGrpcService(*pc)
	test.NewGrpc()
}
