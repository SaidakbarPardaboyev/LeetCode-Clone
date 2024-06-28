package postgres

import (
	"database/sql"
	"fmt"
	model "leetcode/models"
	"leetcode/pkg"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProblemRepo struct {
	ExampleRepo      *ExampleRepo
	TopicProblemRepo *TopicProblemRepo
	Db               *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{
		Db:               db,
		ExampleRepo:      NewExampleRepo(db),
		TopicProblemRepo: NewTopicProblemRepo(db),
	}
}

// Create
func (p *ProblemRepo) CreateProblem(problem *model.Problem) (string, error) {

	tx, err := p.Db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Commit()

	newId := uuid.NewString()
	query := `
		insert into problems(
			id, title, difficulty, description, hints, constraints
		) values(
		 	$1, $2, $3, $4, $5, $6
		)`
	_, err = tx.Exec(query, newId, problem.Title,
		problem.Difficulty, problem.Description,
		pq.Array(problem.Hints), pq.Array(problem.Constraints))

	if err != nil {
		return "", err
	}
	return newId, nil
}

// Read
func (p *ProblemRepo) GetProblemById(problemId string) (*model.Problem, error) {
	problem := model.Problem{Id: problemId}
	query := `
	select
		title, problem_number, difficulty, description, constraints, hints,
		created_at, updated_at
	from
		problems
	where
		id = $1 and 
		deleted_at is null
	`
	err := p.Db.QueryRow(query, problemId).Scan(&problem.Title,
		&problem.ProblemNumber, &problem.Difficulty, &problem.Description,
		pq.Array(&problem.Constraints), pq.Array(&problem.Hints),
		&problem.CreatedAt, &problem.UpdatedAt)
	if err != nil {
		return nil, err
	}

	examples, err := p.ExampleRepo.GetExamplesByProblemId(problemId)
	if err != nil {
		return nil, err
	}
	problem.Examples = *examples

	problemTopics, err := p.TopicProblemRepo.GetTopicsByProblemId(problemId)
	if err != nil {
		return nil, err
	}
	problem.Topics = problemTopics.TopicNames

	return &problem, err
}

func (p *ProblemRepo) GetProblems(username string, filter *model.ProblemFilter) (*[]model.ProblemSet, error) {
	params := []interface{}{username}
	paramCount := 2

	// Get defult query for get all
	withQuery, selectQuery, whereQuery, innerJoinQuery, groupByQuery,
		havingQuery, orderByQuery := pkg.GetAllDefaultQueries()

	if filter.Sorting != nil {
		err := pkg.FilterProblemsBySorting(filter, &innerJoinQuery, &groupByQuery, &orderByQuery)
		if err != nil {
			return nil, err
		}
	}
	if filter.Status != nil {
		err := pkg.FilterProblemsByStatus(filter, &withQuery, &whereQuery, &params,
			username, &paramCount)
		if err != nil {
			return nil, err
		}
	}
	if filter.Search != nil {
		pkg.FilterProblemsBySearch(filter, &whereQuery, &params, &paramCount, &orderByQuery)
	}
	if filter.TopicsSlugs != nil {
		pkg.FilterProblemsByTopicsSlugs(filter, &withQuery, &innerJoinQuery, &whereQuery,
			&havingQuery, &params, &paramCount)
	}
	if filter.Difficulty != nil {
		newWhere := fmt.Sprintf("difficulty = $%d", paramCount)
		whereQuery = append([]string{newWhere}, whereQuery...)
		params = append(params, *filter.Difficulty)
		paramCount++
	}

	// join all parts of query
	query := pkg.JoinQueryParts(withQuery, selectQuery, whereQuery, innerJoinQuery, groupByQuery, havingQuery, orderByQuery)

	if filter.Limit != nil {
		query += fmt.Sprintf(" limit $%d", paramCount)
		params = append(params, *filter.Limit)
		paramCount++
	}
	if filter.Offset != nil {
		query += fmt.Sprintf(" offset $%d", paramCount)
		params = append(params, *filter.Offset)
		paramCount++
	}

	// fmt.Println(query)
	// fmt.Println(query)
	rows, err := p.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	problems := []model.ProblemSet{}
	for rows.Next() {
		problem := model.ProblemSet{}
		err = rows.Scan(&problem.Id, &problem.Status, &problem.ProblemNumber,
			&problem.Title, &problem.Acceptence, &problem.Difficulty)
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

func (p *ProblemRepo) GetSubmissionStatisticsByProblemId(problemTitle string) (*model.SubmissionStatisticsOfProblem, error) {
	query := `with submission_stat as (
					select
						count(
							case
								when submission_status = 'Accepted' then 1
							end
						) as accepted,
						count(*) as submissions
					from
						submissions 
					where
						problem_id = $1
				)

				select
					accepted,
					submissions,
					round(accepted::numeric / submissions * 100, 1) as acceptence_rate
				from 
					submission_stat;`

	submissionStatisticsOfProblem := &model.SubmissionStatisticsOfProblem{}
	err := p.Db.QueryRow(query, problemTitle).Scan(&submissionStatisticsOfProblem.Accepted,
		&submissionStatisticsOfProblem.Submissions, &submissionStatisticsOfProblem.AcceptanceRate)
	if err != nil {
		return nil, err
	}
	return submissionStatisticsOfProblem, nil
}

func (p *ProblemRepo) GetAllProblemsId() (*[]string, error) {
	var problemsId []string

	queryToGetProblemsId := `select id from problems where deleted_at is null`
	rows, err := p.Db.Query(queryToGetProblemsId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var problemId string
		err := rows.Scan(&problemId)
		if err != nil {
			return nil, err
		}
		problemsId = append(problemsId, problemId)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &problemsId, nil
}

func (p *ProblemRepo) PickRandomProblem() (*model.Problem, error) {
	problemsId, err := p.GetAllProblemsId()
	if err != nil {
		return nil, err
	}

	randomProblemId := (*problemsId)[rand.Intn(len(*problemsId))]

	problem, err := p.GetProblemById(randomProblemId)
	return problem, err
}

// Update
func (p *ProblemRepo) UpdateProblem(problem *model.ProblemUpdate) error {
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `
		update 
			problems 
		set 
			problem_number = $1,
			title = $2,
			difficulty = $3,
			description = $4,
			constraints = $5,
			hints = $6,
			updated_at = $7
		where 
			deleted_at is null and
			id = $8
	`
	result, err := tx.Exec(query, problem.ProblemNumber, problem.Title,
		problem.Difficulty, problem.Description, pq.Array(problem.Constraints),
		pq.Array(problem.Hints), time.Now(), problem.Id)

	if err != nil {
		return err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return fmt.Errorf("nothing updated")
	}
	return nil
}

// Delete
func (p *ProblemRepo) DeleteProblem(problemId string) error {
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `
		update
			problems 
		set 
			deleted_at = $1
		where 
			deleted_at is null and
			id = $2
	`
	result, err := tx.Exec(query, time.Now(), problemId)

	if err != nil {
		return err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return fmt.Errorf("nothing deleted")
	}
	return nil
}

func (p *ProblemRepo) RecoverProblem(problemId string) error {
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `
		update
			problems 
		set 
			deleted_at = null
		where 
			deleted_at is not null and
			id = $1
	`
	result, err := tx.Exec(query, problemId)

	if err != nil {
		return err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return fmt.Errorf("nothing recovered")
	}
	return nil
}
