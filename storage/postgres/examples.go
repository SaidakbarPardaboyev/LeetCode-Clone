package postgres

import (
	"database/sql"
	"leetcode/models"
)

type ExampleRepo struct {
	Db *sql.DB
}

func NewExampleRepo(db *sql.DB) *ExampleRepo {
	return &ExampleRepo{db}
}

func (e *ExampleRepo) GetExamplesByProblemId(problemId string) (*[]models.Example, error) {
	query := `select input, output, explanation from examples where problem_id = $1`

	examples := []models.Example{}

	rows, err := e.Db.Query(query, problemId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		example := models.Example{}
		err := rows.Scan(&example.Input, &example.Output, &example.Explanation)
		if err != nil {
			return nil, err
		}
		examples = append(examples, example)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &examples, nil
}
