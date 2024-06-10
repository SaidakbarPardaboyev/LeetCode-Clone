package main

import (
	// "leetcode/generator"

	"leetcode/handler"
	"leetcode/router"
	"leetcode/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// generator.GenerateAllMockData(db)

	h := handler.NewHandler(db)
	server := router.CreateServer(h)
	server.ListenAndServe()
}
