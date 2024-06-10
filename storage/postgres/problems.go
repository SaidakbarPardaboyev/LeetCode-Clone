package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"

	"github.com/lib/pq"
)

type ProblemRepo struct {
	Db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{db}
}

// Create
func (p *ProblemRepo) CreateProblem(problem *model.Problem) error {

	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	insert into 
	problems(question_number, title, difficulty_level, description, 
	examples, hints, constraints)
	values($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(query, problem.QuestionNumber, problem.Title,
		problem.DifficultyLevel, problem.Description,
		pq.Array(problem.Examples), pq.Array(problem.Hints),
		pq.Array(problem.Constraints))

	return err
}

// Read
func (p *ProblemRepo) GetProblemById(id string) (model.Problem, error) {
	problem := model.Problem{}
	query := `
	select * from problems
	where
		id = $1 and deleted_at is null
	`
	row := p.Db.QueryRow(query, id)
	err := row.Scan(&problem.Id, &problem.QuestionNumber, &problem.Title,
		&problem.DifficultyLevel, &problem.Description,
		pq.Array(&problem.Examples), pq.Array(&problem.Hints),
		pq.Array(&problem.Constraints), &problem.Created_at,
		&problem.Updated_at, &problem.Deleted_at)

	return problem, err
}

func (p *ProblemRepo) GetProblems(filter *model.ProblemFilter) (*[]model.Problem, error) {
	params := []interface{}{}
	paramCount := 1
	query := `
	select * from problems where deleted_at is null`
	if filter.QuestionNumber != nil {
		query += fmt.Sprintf(" and question_number=$%d", paramCount)
		params = append(params, *filter.QuestionNumber)
		paramCount++
	}
	if filter.Title != nil {
		query += fmt.Sprintf(" and title=$%d", paramCount)
		params = append(params, *filter.Title)
		paramCount++
	}
	if filter.DifficultyLevel != nil {
		query += fmt.Sprintf(" and difficulty_level=$%d", paramCount)
		params = append(params, *filter.DifficultyLevel)
		paramCount++
	}

	rows, err := p.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	problems := []model.Problem{}
	for rows.Next() {
		problem := model.Problem{}
		err = rows.Scan(&problem.Id, &problem.QuestionNumber,
			&problem.Title, &problem.DifficultyLevel, &problem.Description,
			pq.Array(&problem.Examples), pq.Array(&problem.Hints),
			pq.Array(&problem.Constraints), &problem.Created_at,
			&problem.Updated_at, &problem.Deleted_at)
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &problems, nil
}

// Update
func (p *ProblemRepo) UpdateProblem(problem *model.Problem) error {
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update users 
	set 
		title=$1,
		difficulty_level=$2,
		description=$3,
		examples=$4,
		hints=$5,
		hints=$6,
		updated_at=$7
	where 
		deleted_at is null and id = $8 `
	_, err = tx.Exec(query, problem.Title, problem.DifficultyLevel,
		problem.Description, problem.Examples, problem.Hints,
		problem.Constraints, time.Now(), problem.Id)

	return err
}

// Delete
func (p *ProblemRepo) DeleteProblem(id string) error {
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update problems 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
