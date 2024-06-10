package main

import (
	// "leetcode/generator"

	"leetcode/handler"
	"leetcode/router"
	"leetcode/storage/postgres"
	"net/http"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// generator.GenerateAllMockData(db)

	h := handler.NewHandler(db)
	r := router.CreateServer(h)
	http.ListenAndServe(":8080", r)

}
