package generator

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

type topicsProblems struct {
	ProblemTitle string   `json:"problem_title"`
	Topics       []string `json:"topics"`
}

func InsertTopicProblems(db *sql.DB) {
	f, err := os.Open("json/topicsProblems.json")
	if err != nil {
		log.Fatalf("Failed to open json file: %s\n", err)
	}
	defer f.Close()

	var topicsProblems []topicsProblems
	err = json.NewDecoder(f).Decode(&topicsProblems)
	if err != nil {
		log.Fatalf("Failed to decode json file: %s\n", err)
	}

	// fmt.Println(testcases)
	for _, topicsProblemsSet := range topicsProblems {
		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Error: Opening tranaction for insert testcases_non_btree: %s", err)
		}
		defer tx.Commit()

		// find problem id
		var problemId string
		query := "select id from problems where title = $1"
		err = tx.QueryRow(query, topicsProblemsSet.ProblemTitle).Scan(&problemId)
		if err != nil {
			log.Fatalf("Error taking problem_id from problems by problem title: %s", err)
		}

		query = `insert into topics_problems(problem_id, topic_id) values($1,$2)`
		for _, topicName := range topicsProblemsSet.Topics {

			// find topic id
			var topicId string
			query1 := "select id from topics where name = $1"
			err = tx.QueryRow(query1, topicName).Scan(&topicId)
			if err != nil {
				log.Fatalf("Error taking topic_id from topics by topic name: %s", err)
			}

			// insert into topics_problems
			result, err := tx.Exec(query, problemId, topicId)
			if err != nil {
				log.Printf("Failed to insert topic_problem into database: %s", err)
			}
			if check, _ := result.RowsAffected(); check <= 0 {
				log.Printf("Error: there is no topics_problems for topic_title(%s): %s", topicName, err)
			}
		}
	}
}
