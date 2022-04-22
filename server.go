package main

import (
	"net/http"
	"test-api/database"
	"test-api/router"
)

func main() {
	db := database.NewDB()

	r := router.New(db)

	r.Logger.Fatal(r.Start(":1324"))
	http.ListenAndServe(":1324", r)
}
