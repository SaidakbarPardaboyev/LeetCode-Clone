package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Answer struct {
	AnswerId            string           `json:"answer_id"`
	ProblemTitle        string           `json:"problem_title"`
	Int2dArrayAnswer    *pq.Int64Array   `json:"int_2d_array_answer"`
	Float2dArrayAnswer  *pq.Float64Array `json:"float_2d_array_answer"`
	String2dArrayAnswer *pq.StringArray  `json:"string_2d_array_answer"`
	Byte2dArrayAnswer   *[]byte          `json:"byte_2d_array_answer"`
	Bool2dArrayAnswer   *pq.BoolArray    `json:"bool_2d_array_answer"`
	IntArrayAnswer      pq.Int64Array    `json:"int_array_answer"`
	FloatArrayAnswer    pq.Float64Array  `json:"float_array_answer"`
	StringArrayAnswer   pq.StringArray   `json:"string_array_answer"`
	ByteArrayAnswer     []byte           `json:"byte_array_answer"`
	BoolArrayAnswer     pq.BoolArray     `json:"bool_array_answer"`
	IntValAnswer        sql.NullInt64    `json:"int_val_answer"`
	FloatValAnswer      sql.NullFloat64  `json:"float_val_answer"`
	StringValAnswer     sql.NullString   `json:"string_val_answer"`
	ByteValAnswer       sql.NullByte     `json:"byte_val_answer"`
	BoolValAnswer       sql.NullBool     `json:"bool_val_answer"`
}

type TestCasee struct {
	Id             string       `json:"id"`
	ProblemTitle   string       `json:"problem_title"`
	AnswerId       string       `json:"answer_id"`
	Int2dArray     [][]int      `json:"int_2d_array"`
	Int2dArray2    [][]int      `json:"int_2d_array2"`
	Float2dArray   [][]float64  `json:"float_2d_array"`
	Float2dArray2  [][]float64  `json:"float_2d_array2"`
	String2dArray  [][]string   `json:"string_2d_array"`
	String2dArray2 [][]string   `json:"string_2d_array2"`
	Byte2dArray    [][]byte     `json:"byte_2d_array"`
	Byte2dArray2   [][]byte     `json:"byte_2d_array2"`
	Bool2dArray    [][]bool     `json:"bool_2d_array"`
	Bool2dArray2   [][]bool     `json:"bool_2d_array2"`
	IntArray       []int        `json:"int_array"`
	IntArray2      []int        `json:"int_array2"`
	FloatArray     []float64    `json:"float_array"`
	FloatArray2    []float64    `json:"float_array2"`
	StringArray    []string     `json:"string_array"`
	StringArray2   []string     `json:"string_array2"`
	ByteArray      []byte       `json:"byte_array"`
	ByteArray2     []byte       `json:"byte_array2"`
	BoolArray      []bool       `json:"bool_array"`
	BoolArray2     []bool       `json:"bool_array2"`
	IntVal         int          `json:"int_val"`
	IntVal2        int          `json:"int_val2"`
	IntVal3        int          `json:"int_val3"`
	FloatVal       float64      `json:"float_val"`
	FloatVal2      float64      `json:"float_val2"`
	FloatVal3      float64      `json:"float_val3"`
	StringVal      string       `json:"string_val"`
	StringVal2     string       `json:"string_val2"`
	StringVal3     string       `json:"string_val3"`
	ByteVal        byte         `json:"byte_val"`
	ByteVal2       byte         `json:"byte_val2"`
	ByteVal3       byte         `json:"byte_val3"`
	BoolVal        bool         `json:"bool_val"`
	BoolVal2       bool         `json:"bool_val2"`
	BoolVal3       bool         `json:"bool_val3"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
	DeletedAt      sql.NullTime `json:"deleted_at"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "leetcode"
	password = "root"
)

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	title := "Two Sum"

	answers := map[string]Answer{}

	queryAnswers := "with problem_title_of_question as ( select problem_title from testcases_non_btree WHERE problem_title = $1 limit 1 ) select answer_id, a.problem_title, int_2d_array_answer, float_2d_array_answer, string_2d_array_answer, byte_2d_array_answer, bool_2d_array_answer, int_array_answer, float_array_answer, string_array_answer, byte_array_answer, bool_array_answer, int_val_answer, float_val_answer, string_val_answer, byte_val_answer, bool_val_answer from answers as a inner join problem_title_of_question as p on p.problem_title = a.problem_title;"
	rows, err := db.Query(queryAnswers, title)
	for rows.Next() {
		ans := Answer{}
		err := rows.Scan(&ans.AnswerId, &ans.ProblemTitle, &ans.Int2dArrayAnswer,
			&ans.Float2dArrayAnswer, &ans.String2dArrayAnswer, &ans.Byte2dArrayAnswer,
			&ans.Bool2dArrayAnswer, &ans.IntArrayAnswer, &ans.FloatArrayAnswer,
			&ans.StringArrayAnswer, &ans.ByteArrayAnswer, &ans.BoolArrayAnswer,
			&ans.IntValAnswer, &ans.FloatValAnswer, &ans.StringValAnswer,
			&ans.ByteValAnswer, &ans.BoolValAnswer)
		if err != nil {
			panic(err)
		}
		answers[ans.AnswerId] = ans
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	// for _, val := range answers {
	// 	fmt.Printf("%+v\n", val)
	// }

	queryTestcases := "select * from testcases_non_btree WHERE problem_title = $1;"
	rows, err = db.Query(queryTestcases, "Two Sum")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cases := []TestCasee{}
	for rows.Next() {
		tc := TestCasee{}
		err = rows.Scan(&tc.Id, &tc.ProblemTitle, &tc.AnswerId, &tc.Int2dArray,
			&tc.Int2dArray2, &tc.Float2dArray, &tc.Float2dArray2, &tc.String2dArray,
			&tc.String2dArray2, &tc.Byte2dArray, &tc.Byte2dArray2, &tc.Bool2dArray,
			&tc.Bool2dArray2, &tc.IntArray, &tc.IntArray2, &tc.FloatArray,
			&tc.FloatArray2, &tc.StringArray, &tc.StringArray2, &tc.ByteArray,
			&tc.ByteArray2, &tc.BoolArray, &tc.BoolArray2, &tc.IntVal, &tc.IntVal2,
			&tc.IntVal3, &tc.FloatVal, &tc.FloatVal2, &tc.FloatVal3, &tc.StringVal,
			&tc.StringVal2, &tc.StringVal3, &tc.ByteVal, &tc.ByteVal2,
			&tc.ByteVal3, &tc.BoolVal, &tc.BoolVal2, &tc.BoolVal3, &tc.CreatedAt,
			&tc.UpdatedAt, &tc.DeletedAt)

		if err != nil {
			panic(err)
		}
		cases = append(cases, tc)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	for _, val := range cases {
		fmt.Printf("%+v\n", val)
	}

	// for _, c := range cases {
	// 	ans := twoSum(c.IntArray, c.IntVal)
	// 	if fmt.Sprintf("%v", ans) != c.Answer {
	// 		panic(fmt.Errorf("wrong answer: got %v, expected %v", ans, c.Answer))
	// 	}
	// }
	fmt.Println("Accepted")
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
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

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
	rows, err := db.Query("SELECT param1, param2, param3, param4, param5, param6, answer FROM testcases_non_btree")
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
