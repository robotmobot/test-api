package main

import (
	"test-api/controller"
	"test-api/database"
	"test-api/gRPC/service"
	"test-api/router"
)

func main() {
	db := database.NewDB()
	pc := controller.NewProductController(db)
	test := service.NewGrpcService(*pc)
	test.GRPCSERVER()

	r := router.NewEcho(*db)
	r.Logger.Fatal(r.Start("localhost:1324"))

}
