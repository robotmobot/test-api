package main

import (
	"test-api/controller"
	"test-api/database"
	"test-api/handler/gRPC"
)

func main() {
	db := database.NewDB()
	rc := database.ConnectRedis()
	pc := controller.NewProductController(db, *rc)
	test := gRPC.NewGrpcService(*pc)
	test.NewGrpc()
}
