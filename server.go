package main

import (
	"test-api/database"
	"test-api/router"
)

func main() {
	db := database.NewDB()

	r := router.New(db)

	r.Logger.Fatal(r.Start("localhost:1324"))
	//http.ListenAndServe("localhost:1324", r)

}
