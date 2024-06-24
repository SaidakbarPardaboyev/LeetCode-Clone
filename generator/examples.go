package generator

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type example struct {
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation"`
}

type exampleSet struct {
	Title    string    `json:"title"`
	Examples []example `json:"examples"`
}

func InsertExamples(db *sql.DB) {
	f, err := os.Open("json/examples.json")
	if err != nil {
		log.Fatalf("Failed to open json file: %s\n", err)
	}
	defer f.Close()

	var examples []exampleSet
	err = json.NewDecoder(f).Decode(&examples)
	if err != nil {
		log.Fatalf("Failed to decode json file: %s\n", err)

	}

	for _, exampleSet := range examples {
		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Error: Opening tranaction for insert examples: %s", err)
		}
		defer tx.Commit()
		
		// find problem id
		var problemId string
		query := "select id from problems where title = $1"
		err = tx.QueryRow(query, exampleSet.Title).Scan(&problemId)
		if err != nil {
			log.Fatalf("Error taking problem_id from problems by problem title: %s", err)
		}
		query = `insert into examples(problem_id, input, output, explanation) values($1,$2,$3,$4)`
		for _, example := range exampleSet.Examples {
			// insert into examples
			result, err := tx.Exec(query, problemId, example.Input, example.Output, example.Explanation)
			if err != nil {
				log.Printf("Failed to insert example into database: %s", err)
			}
			if check, _ := result.RowsAffected(); check <= 0 {
				log.Printf("Error: there is no problem for problem_title(%s): %s", exampleSet.Title, err)
			}
		}
	}
}
