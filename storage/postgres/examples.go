package postgres

import (
	"database/sql"
	"leetcode/models"
	"time"

	"github.com/google/uuid"
)

type ExampleRepo struct {
	Db *sql.DB
}

func NewExampleRepo(db *sql.DB) *ExampleRepo {
	return &ExampleRepo{db}
}

func (e *ExampleRepo) CreateExample(example *models.ExampleCreate) (*string, error) {
	tx, err := e.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := `insert into examples(
					id, problem_id, input, output, explanation
				) values (
					$1, $2, $3, $4, $5 
				)`
	newId := uuid.NewString()
	_, err = tx.Exec(query, newId, example.ProblemId, example.Input,
		example.Output, example.Explanation)
	if err != nil {
		return nil, err
	}
	return &newId, nil
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

// Update
func (e *ExampleRepo) UpdateExample(tp *models.ExampleUpdate) error {
	tx, err := e.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update examples
	set
		problem_id = $1,
		input = $2,
		output = $3,
		explanation = $4,
		updated_at = $5
	where
		deleted_at is null and id = $4`
	_, err = tx.Exec(query, tp.ProblemId, tp.Input, tp.Output,
		tp.Explanation, time.Now(), tp.Id)

	return err
}

// Delete
func (e *ExampleRepo) DeleteExample(id string) error {
	tx, err := e.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update examples
	set
		deleted_at = $1
	where
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}

// Recover
func (e *ExampleRepo) RecoverExample(id string) error {
	tx, err := e.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update examples
	set
		deleted_at = null
	where
		deleted_at is not null and id = $1`
	_, err = tx.Exec(query, id)

	return err
}
