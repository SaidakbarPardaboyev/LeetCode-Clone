package postgres

import (
	"database/sql"
	"fmt"
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
func (s *SubmissionRepo) CreateSubmission(submission *models.CreateSubmission) error {
	
	query := `
	insert into 
	submissions(problem_id, user_id, language_id, 
	code, submission_status, runtime, submission_date)
	values($1, $2, $3, $4, $5, $6, $7)`
	
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	_, execErr := tx.Exec(query, submission.ProblemId, submission.UserId,
		submission.LanguageId, submission.Code, submission.SubmissionStatus,
		submission.Runtime, submission.SubmissionDate)

	if execErr != nil {
		tx.Rollback() // Rollback the transaction in case of error
		return execErr
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return commitErr
	}

	return nil
}

// Read
func (s *SubmissionRepo) GetSubmissionById(id int) (*models.Submission, error) {
	submission := models.Submission{Id: id}
	query := `
	select 
		problem_id, user_id, language_id, 
		code, submission_status, runtime, submission_date,
		created_at, updated_at, deleted_at 
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
		&submission.DeletedAt,
	)

	return &submission, err
}

func (s *SubmissionRepo) GetSubmissions(filter *models.SubmissionFilter) (*[]models.Submission, error) {
	params := []interface{}{}
	paramCount := 1
	query := `
	select 
		problem_id, user_id, language_id, 
		code, submission_status, runtime, submission_date,
		created_at, updated_at, deleted_at  
	from 
		submissions 
	where 
		deleted_at is null`
	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and problem_title=$%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and user_username=$%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and language_name=$%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.Code != nil {
		query += fmt.Sprintf(" and code=$%d", paramCount)
		params = append(params, *filter.Code)
		paramCount++
	}
	if filter.SubmissionStatus != nil {
		query += fmt.Sprintf(" and submission_status=$%d", paramCount)
		params = append(params, *filter.SubmissionStatus)
		paramCount++
	}
	if filter.Code != nil {
		query += fmt.Sprintf(" and code=$%d", paramCount)
		params = append(params, *filter.Code)
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
			&submission.CreatedAt, &submission.UpdatedAt, &submission.DeletedAt)

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

func (s *SubmissionRepo) GetActiveDays(userId string, year *int) (*models.UserActivity, error) {
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
	rows, err := s.Db.Query(query, userId, *year)
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
		problem_id,
		min(submission_date) as recent_submission
	from
		submissions
	where
		user_username = $1
	group by
		problem_title
	order by
		recent_submission desc
	limit 15
	`

	recentAc := []models.RecentlyAcceptedSubmission{}

	rows, err := s.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		sub := models.RecentlyAcceptedSubmission{}
		err = rows.Scan(&sub.ProblemTitle, &sub.SubmissionDate)
		if err != nil {
			return nil, err
		}
		recentAc = append(recentAc, sub)
	}

	return &recentAc, rows.Err()
}

// Update
func (s *SubmissionRepo) UpdateSubmission(submission *models.Submission) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 
		submissions 
	set 
		code=$1,
		submission_status=$2,
		submission_date=$3,
		updated_at=$4
	where 
		deleted_at is null and id = $5 `
	result, err := tx.Exec(query, submission.Code, submission.SubmissionStatus,
		submission.SubmissionDate, time.Now(), submission.Id)

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
func (s *SubmissionRepo) DeleteSubmission(id int) error {
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
