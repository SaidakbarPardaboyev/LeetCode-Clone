package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost user=postgres dbname=just password=root sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	funcName := "Two Sum"

	rows, err := db.Query("SELECT function_name, arg1, arg2, arg3, arg4, "+
		"arg5, arg6, answer, arg1_type, arg2_type, arg3_type, arg4_type, arg5_type, "+
		"arg6_type, answer_type FROM function_calls WHERE function_name=$1", funcName)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var functionName string
		var arg1, arg2, arg3, arg4, arg5, arg6 sql.NullString
		var arg1Type, arg2Type, arg3Type, arg4Type, arg5Type, arg6Type sql.NullString
		var answer, answerType sql.NullString
		if err := rows.Scan(&functionName, &arg1, &arg2, &arg3, &arg4, &arg5, &arg6,
			&answer, &arg1Type, &arg2Type, &arg3Type, &arg4Type, &arg5Type, &arg6Type,
			&answerType); err != nil {
			log.Fatal(err)
		}

		args := []sql.NullString{arg1, arg2, arg3, arg4, arg5, arg6}
		argsTypes := []sql.NullString{arg1Type, arg2Type, arg3Type, arg4Type,
			arg5Type, arg6Type}

		results, err := callFunction(functionName, args, argsTypes, answerType)
		if err != nil {
			panic(err)
		}
		if len(results) > 0 {
			res, err := checkResult(results, answer)
			if err != nil {
				panic(err)
			}
			if !res {
				fmt.Printf("Wrong answer\nOutput: %v\nExpected: %v\n", results[0], answer.String)
				return
			}
		}
	}
	fmt.Println("Accepted")
}

func callFunction(name string, args []sql.NullString, argsTypes []sql.NullString, answertype sql.NullString) ([]reflect.Value, error) {
	fn, ok := funcMap[name]
	if !ok {
		log.Fatalf("Function %s not found", name)
	}

	fnType := reflect.TypeOf(fn)
	if fnType.NumIn() > 6 {
		log.Fatalf("Function %s has more than 6 arguments", name)
	}

	var callArgs []reflect.Value
	for i := 0; i < fnType.NumIn(); i++ {
		argType := fnType.In(i)
		argValue := reflect.New(argType).Elem()

		if err := checkType(argType, argsTypes[i], name, fmt.Sprintf("param_%d", i+1)); err != nil {
			return nil, err
		}

		if args[i].Valid {
			if err := json.Unmarshal([]byte(args[i].String), argValue.Addr().Interface()); err != nil {
				log.Fatalf("Error unmarshaling argument %d for function %s: %v", i+1, name, err)
			}
		} else {
			argValue = reflect.Zero(argType)
		}

		callArgs = append(callArgs, argValue)
	}

	result := reflect.ValueOf(fn).Call(callArgs)
	err := checkType(result[0].Type(), answertype, name, "ret")
	return result, err
}

func checkType(argType reflect.Type, correctType sql.NullString, functionName string, ParamOrReturnVal string) error {
	// check primitive types
	if argType.Kind() != reflect.Slice {
		if argType.Kind().String() != correctType.String {
			return fmt.Errorf("cannot use %s (variable of type %s) as %s value in argument to %s", ParamOrReturnVal, correctType.String, argType.Kind().String(), functionName)
		}
		return nil
	}

	// check 1D array types
	if argType.Elem().Kind() != reflect.Slice {
		if "[]"+argType.Elem().Kind().String() != correctType.String {
			return fmt.Errorf("cannot use %s (variable of type %s) as []%s value in argument to %s", ParamOrReturnVal, correctType.String, argType.Elem().Kind().String(), functionName)
		}
		return nil
	}

	// check 2D array types
	if "[][]"+argType.Elem().Elem().Kind().String() != correctType.String {
		return fmt.Errorf("cannot use %s (variable of type %s) as [][]%s value in argument to %s", ParamOrReturnVal, correctType.String, argType.Elem().Elem().Kind().String(), functionName)
	}
	return nil
}

func checkResult(result []reflect.Value, answerJson sql.NullString) (bool, error) {
	var answerInterface interface{}

	err := json.Unmarshal([]byte(answerJson.String), &answerInterface)
	if err != nil {
		return false, err
	}

	answer := reflect.ValueOf(answerInterface)

	switch answer.Kind() {
	case reflect.Int:
		if len(result) == 1 && result[0].Kind() == reflect.Int {
			return result[0].Int() == answer.Int(), nil
		}
	case reflect.String:
		if len(result) == 1 && result[0].Kind() == reflect.String {
			return result[0].String() == answer.String(), nil
		}
	case reflect.Float64:
		if len(result) == 1 && result[0].Kind() == reflect.Float64 {
			return result[0].Float() == answer.Float(), nil
		}
	case reflect.Bool:
		if len(result) == 1 && result[0].Kind() == reflect.Bool {
			return result[0].Bool() == answer.Bool(), nil
		}
	case reflect.Slice:
		if len(result) == 1 && result[0].Kind() == reflect.Slice {
			resultSlice := result[0]
			if resultSlice.Len() != answer.Len() {
				return false, nil
			}
			if !reflect.DeepEqual(answer.Interface().([]interface{}), answerInterface.([]interface{})) {
				return false, nil
			}
			return true, nil
		}
	}

	fmt.Println("result does not match answer")
	return false, nil
}

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

var funcMap = map[string]interface{}{
	"Two Sum": twoSum,
}
