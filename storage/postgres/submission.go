package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"leetcode/models"
	"time"
)

type SubmissionRepo struct {
	Db *sql.DB
}

func NewSubmissionRepo(db *sql.DB) *SubmissionRepo {
	return &SubmissionRepo{db}
}

// Create
func (s *SubmissionRepo) CreateSubmission(submission *models.CreateSubmission) (string, error) {

	id := uuid.NewString()
	query := `
	insert into 
	submissions(id, problem_id, user_id, language_id, 
	code, submission_status, runtime, submission_date, created_at)
	values($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	tx, err := s.Db.Begin()
	if err != nil {
		return "", err
	}

	_, execErr := tx.Exec(query, id, submission.ProblemId, submission.UserId,
		submission.LanguageId, submission.Code, submission.SubmissionStatus,
		submission.Runtime, time.Now(), time.Now())

	if execErr != nil {
		tx.Rollback()
		return "", execErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return "", commitErr
	}

	return id, nil
}

// Read
func (s *SubmissionRepo) GetSubmissionById(id string) (*models.Submission, error) {
	submission := models.Submission{Id: id}
	query := `
	select 
		problem_id, user_id, language_id, 
		code, submission_status, runtime, submission_date,
		created_at, updated_at 
	from 
		submissions
	where
		id = $1 and deleted_at is null
	`
	row := s.Db.QueryRow(query, id)
	err := row.Scan(
		&submission.ProblemId, &submission.UserId, &submission.LanguageId,
		&submission.Code, &submission.SubmissionStatus, &submission.Runtime,
		&submission.SubmissionDate, &submission.CreatedAt, &submission.UpdatedAt,
	)

	return &submission, err
}

func (s *SubmissionRepo) GetSubmissions(filter *models.SubmissionFilter) (*[]models.Submission, error) {
	params := []interface{}{}
	paramCount := 1
	query := `
	select 
		id, problem_id, user_id, language_id, 
		code, submission_status, runtime, submission_date,
		created_at, updated_at  
	from 
		submissions 
	where 
		deleted_at is null`

	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and problem_id = $%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.UserId != nil {
		query += fmt.Sprintf(" and user_id = $%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.LanguageId != nil {
		query += fmt.Sprintf(" and language_id = $%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.Code != nil {
		query += fmt.Sprintf(" and code = $%d", paramCount)
		params = append(params, *filter.Code)
		paramCount++
	}
	if filter.SubmissionStatus != nil {
		query += fmt.Sprintf(" and submission_status = $%d", paramCount)
		params = append(params, *filter.SubmissionStatus)
		paramCount++
	}
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

	rows, err := s.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	submissions := []models.Submission{}
	for rows.Next() {
		submission := models.Submission{}
		err = rows.Scan(&submission.Id, &submission.ProblemId,
			&submission.UserId, &submission.LanguageId, &submission.Code,
			&submission.SubmissionStatus, &submission.Runtime, &submission.SubmissionDate,
			&submission.CreatedAt, &submission.UpdatedAt)

		if err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &submissions, nil
}

func (s *SubmissionRepo) GetActiveDays(userId string, year int) (*models.UserActivity, error) {
	if year == nil {
		y := time.Now().Year()
		year = &y
	}
	query := `
	select
		submission_date,
		count(*) as number_of_submissions
	from
		submissions
	where
		user_id = $1 and
		extract(year from submission_date) = $2
	group by
		submission_date
	order by
		submission_date
	`
	userActivity := models.UserActivity{}
	rows, err := s.Db.Query(query, userId, year)
	if err != nil {
		return nil, err
	}
	totalSubmissions := 0
	for rows.Next() {
		submissionDay := models.SubmissionDay{}
		err = rows.Scan(&submissionDay.Date, &submissionDay.SubmissionCount)
		if err != nil {
			return nil, err
		}

		totalSubmissions += submissionDay.SubmissionCount
		userActivity.SubmissionDays = append(userActivity.SubmissionDays, submissionDay)
	}
	userActivity.TotalSubmissions = totalSubmissions

	return &userActivity, rows.Err()
}

func (s *SubmissionRepo) GetRecentlyAcceptedSubmissionsByUserId(userId string) (*[]models.RecentlyAcceptedSubmission, error) {
	query := `
	select
		p.title,
		min(submission_date) as recent_submission
	from
		submissions as s
	join
		problems as p
	on
		s.user_id = $1 and p.id = s.problem_id
	group by
		p.title
	order by
		recent_submission desc
	limit 15
	`

	recentAc := []models.RecentlyAcceptedSubmission{}

	rows, err := s.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		sub := models.RecentlyAcceptedSubmission{}
		err = rows.Scan(&sub.ProblemTitle, &sub.SubmissionDate)
		if err != nil {
			return nil, err
		}
		recentAc = append(recentAc, sub)
	}

	return &recentAc, rows.Err()
}

func (s *SubmissionRepo) GetLastSubmittedCodeByUserId(userId, problemId string) (string, error) {
	query := `
	select 
		code
	from
	    submissions
	where
	    user_id = $1 and problem_id = $2
	order by
	    submission_date desc
`
	code := ""
	err := s.Db.QueryRow(query, userId, problemId).Scan(&code)
	if err != nil {
		return "", err
	}

	return code, nil
}

// Update
func (s *SubmissionRepo) UpdateSubmission(submission *models.UpdateSubmission) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 
		submissions 
	set 
		submission_status=$1,
		runtime=$2,
		updated_at=$3
	where 
		deleted_at is null and id = $4 `
	result, err := tx.Exec(query, submission.SubmissionStatus,
		submission.Runtime, time.Now(), submission.Id)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing updated")
	}
	return err
}

// Delete
func (s *SubmissionRepo) DeleteSubmission(id string) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 
		submissions 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	result, err := tx.Exec(query, time.Now(), id)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing deleted")
	}
	return err
}
