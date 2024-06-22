package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"leetcode/pkg"
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
		insert into problems(
			title, difficulty, description, hints, constraints
		) values(
		 	$1, $2, $3, $4, $5
		)`
	_, err = tx.Exec(query, problem.Title,
		problem.Difficulty, problem.Description,
		pq.Array(problem.Hints), pq.Array(problem.Constraints))

	return err
}

// Read
func (p *ProblemRepo) GetProblemByTitle(problemTitle string) (model.Problem, error) {
	problem := model.Problem{}
	query := `
		select
			*
		from
			problems
		where
			title = $1 and 
			deleted_at is null
	`
	row := p.Db.QueryRow(query, problemTitle)
	err := row.Scan(&problem.Title, &problem.ProblemNumber,
		&problem.Difficulty, &problem.Description,
		pq.Array(&problem.Constraints), pq.Array(&problem.Hints), &problem.CreatedAt,
		&problem.UpdatedAt, &problem.DeletedAt)

	return problem, err
}

func (p *ProblemRepo) GetProblems(username string, filter *model.ProblemFilter) (*[]model.Problems, error) {
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
		pkg.FilterProblemsByTopicsSlugs(filter, &havingQuery, &whereQuery, &params,
			&paramCount, &innerJoinQuery)
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
	rows, err := p.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	problems := []model.Problems{}
	for rows.Next() {
		problem := model.Problems{}
		err = rows.Scan(&problem.Status, &problem.ProblemNumber,
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

func (p *ProblemRepo) GetSubmissionStatisticsByProblemTitle(problemTitle string) (*model.SubmissionStatisticsOfProblem, error) {
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
					problem_title = $1
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

// Update
func (p *ProblemRepo) UpdateProblem(problem *model.Problem) error {
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
		update 
			problems 
		set 
			title = $1,
			difficulty = $2,
			description = $3,
			hints = $4,
			constraints = $5,
			updated_at = $6
		where 
			deleted_at is null and
			problem_number = $7
	`
	result, err := tx.Exec(query, problem.Title, problem.Difficulty,
		problem.Description, pq.Array(problem.Hints),
		pq.Array(problem.Constraints), time.Now(), problem.ProblemNumber)

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
func (p *ProblemRepo) DeleteProblem(problemTitle string) error {
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
			title = $2
	`
	result, err := tx.Exec(query, time.Now(), problemTitle)

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
