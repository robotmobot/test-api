package main

import (
	"test-api/database"
	"test-api/router"
)

func main() {
	db := database.NewDB()
	r := router.NewEcho(*db)
	r.Logger.Fatal(r.Start("localhost:1324"))
}
