package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

// Testcase represents a generic test case with dynamic parameters
type Testcase struct {
	Params []interface{}
	Answer interface{}
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "leetcode"
	password = "root"
)

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s "+
		"sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	testcases := []Testcase{}
	rows, err := db.Query("SELECT param1, param2, param3, param4, param5, " + "param6, answer FROM testcases_non_btree")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var param1, param2, param3, param4, param5, param6, answer interface{}
		err = rows.Scan(&param1, &param2, &param3, &param4, &param5, &param6, &answer)
		if err != nil {
			log.Fatal(err)
		}
		params := []interface{}{param1, param2, param3, param4, param5, param6}
		// Remove nil parameters from the slice
		filteredParams := filterNilParams(params)
		testcases = append(testcases, Testcase{Params: filteredParams, Answer: answer})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	funcMap := map[string]interface{}{
		"twoSum": twoSum,
	}

	for _, testcase := range testcases {
		funcName := "twoSum" // This should be dynamically determined based on your use case
		result := callFunction(funcMap[funcName], testcase.Params...)

		expectedAns, ok := testcase.Answer.([]int)
		if !ok || !reflect.DeepEqual(result, expectedAns) {
			log.Fatalf("Test failed: got %v, want %v", result, expectedAns)
		} else {
			fmt.Println("Test passed")
		}
	}
}

// filterNilParams removes nil values from the params slice
func filterNilParams(params []interface{}) []interface{} {
	filtered := []interface{}{}
	for _, param := range params {
		if param != nil {
			filtered = append(filtered, param)
		}
	}
	return filtered
}

// callFunction uses reflection to call a function with a dynamic number of arguments
func callFunction(fn interface{}, params ...interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)
	if len(params) != fnValue.Type().NumIn() {
		log.Fatalf("Number of parameters mismatch: expected %d, got %d", fnValue.Type().NumIn(), len(params))
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	result := fnValue.Call(in)
	if len(result) == 0 {
		return nil
	}

	return result[0].Interface()
}

// twoSum is an example function that can be dynamically called
func twoSum(nums []int, target int) []int {
	count := map[int][]int{}
	for i, num := range nums {
		count[num] = append(count[num], i)
		if len(count[target-num]) > 0 {
			if count[target-num][0] != i {
				return []int{count[target-num][0], i}
			}
		}
	}
	return []int{}
}
