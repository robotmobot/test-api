package main

import (
	"test-api/database"
	"test-api/router"
)

func main() {
	db := database.NewDB()
	rc := database.ConnectRedis()
	r := router.NewEcho(*db, *rc)

	r.Logger.Fatal(r.Start("localhost:1324"))
}
