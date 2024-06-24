package generator

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type testcase struct {
	RunOrSubmit string      `json:"run_or_submit"`
	Arg1        interface{} `json:"arg1"`
	Arg2        interface{} `json:"arg2"`
	Arg3        interface{} `json:"arg3"`
	Arg4        interface{} `json:"arg4"`
	Arg5        interface{} `json:"arg5"`
	Arg6        interface{} `json:"arg6"`
	Answer      interface{} `json:"answer"`
	Arg1Type    string      `json:"arg1_type"`
	Arg2Type    string      `json:"arg2_type"`
	Arg3Type    string      `json:"arg3_type"`
	Arg4Type    string      `json:"arg4_type"`
	Arg5Type    string      `json:"arg5_type"`
	Arg6Type    string      `json:"arg6_type"`
	AnswerType  string      `json:"answer_type"`
}

type testcaseSet struct {
	ProblemTile  string `json:"problem_title"`
	FunctionName string `json:"function_name"`
	Testcases    []testcase
}

func InsertTestcasesNonBtree(db *sql.DB) {
	f, err := os.Open("json/testcasesNonBtree.json")
	if err != nil {
		log.Fatalf("Failed to open json file: %s\n", err)
	}
	defer f.Close()

	var testcases []testcaseSet
	err = json.NewDecoder(f).Decode(&testcases)
	if err != nil {
		log.Fatalf("Failed to decode json file: %s\n", err)

	}
	// fmt.Println(testcases)
	for _, testcaseSet := range testcases {
		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Error: Opening tranaction for insert testcases_non_btree: %s", err)
		}
		defer tx.Commit()

		// find problem id
		var problemId string
		query := "select id from problems where title = $1"
		err = tx.QueryRow(query, testcaseSet.ProblemTile).Scan(&problemId)
		if err != nil {
			log.Fatalf("Error taking problem_id from problems by problem title: %s", err)
		}

		functionName := testcaseSet.FunctionName

		query = `insert into testcases_non_btree(problem_id, function_name, run_or_submit, 
		arg1, arg2, arg3, arg4, arg5, arg6, answer, arg1_type, arg2_type, arg3_type, arg4_type,
		arg5_type, arg6_type, answer_type) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,
		$14,$15,$16,$17)`
		for _, t := range testcaseSet.Testcases {
			// insert into examples
			result, err := tx.Exec(query, problemId, functionName, t.RunOrSubmit, t.Arg1, t.Arg2,
				t.Arg3, t.Arg4, t.Arg5, t.Arg6, t.Answer, t.Arg1Type, t.Arg2Type, t.Arg3Type,
				t.Arg4Type, t.Arg5Type, t.Arg6Type, t.AnswerType)
			if err != nil {
				log.Printf("Failed to insert example into database: %s", err)
			}
			if check, _ := result.RowsAffected(); check <= 0 {
				log.Printf("Error: there is no problem for problem_title(%s): %s", testcaseSet.ProblemTile, err)
			}
		}
	}
}
