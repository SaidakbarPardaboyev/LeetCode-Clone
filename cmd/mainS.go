package main

import (
	// "leetcode/generator"

	"fmt"
	"leetcode/pkg"
	"leetcode/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// generator.GenerateAllMockData(db)

	// h := handler.NewHandler(db)
	// server := router.CreateServer(h)
	// server.ListenAndServe()

	res, err := pkg.Run("twoSum", `func twoSum(nums []int, target int) []int {
				count := map[int][]int{}
				for i, num := range nums {
					count[num] = append(count[num], i)
					if len(count[target-num]) > 0 {
					if count[target-num][0] != i {
						return []int{count[target-num][0], i}
					}
					}
				}
					for i := 0; i < 11000000000000000; i++ {
						fmt.Println("Hello")
					}
				return []int{}
				}
				`)

	fmt.Println(*res, err)
}
