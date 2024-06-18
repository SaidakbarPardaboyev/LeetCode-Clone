package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type SubmissionRepo struct {
	Db *sql.DB
}

func NewSubmissionRepo(db *sql.DB) *SubmissionRepo {
	return &SubmissionRepo{db}
}

// Create
func (u *SubmissionRepo) CreateSubmission(submission *model.Submission) error {

	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	insert into 
		submissions(problem_title, user_username, language_name, 
		code, submission_status, submission_date)
	values($1, $2, $3, $4, $5, $6)`
	_, err = tx.Exec(query, submission.ProblemTitle, submission.UserUsername,
		submission.LanguageName, submission.Code, submission.SubmissionStatus,
		submission.SubmissionDate)

	return err
}

// Read
func (u *SubmissionRepo) GetSubmissionById(id string) (*model.Submission, error) {
	submission := model.Submission{}
	query := `
	select * 
	from 
		submissions
	where
		id = $1 and deleted_at is null
	`
	row := u.Db.QueryRow(query, id)
	err := row.Scan(&submission.Id, &submission.ProblemTitle,
		&submission.UserUsername, &submission.LanguageName, &submission.Code,
		&submission.SubmissionStatus, &submission.SubmissionDate,
		&submission.CreatedAt, &submission.UpdatedAt,
			&submission.DeletedAt)

	return &submission, err
}

func (u *SubmissionRepo) GetSubmissions(filter *model.SubmissionFilter) (*[]model.Submission, error) {
	params := []interface{}{}
	paramCount := 1
	query := `
	select * 
	from 
		submissions 
	where 
		deleted_at is null`
	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and problem_title=$%d", paramCount)
		params = append(params, *filter.ProblemId)
		paramCount++
	}
	if filter.UserId != nil {
		query += fmt.Sprintf(" and user_username=$%d", paramCount)
		params = append(params, *filter.UserId)
		paramCount++
	}
	if filter.LanguageId != nil {
		query += fmt.Sprintf(" and language_name=$%d", paramCount)
		params = append(params, *filter.LanguageId)
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

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	submissions := []model.Submission{}
	for rows.Next() {
		submission := model.Submission{}
		err = rows.Scan(&submission.Id, &submission.ProblemTitle,
			&submission.UserUsername, &submission.LanguageName, &submission.Code,
			&submission.SubmissionStatus, &submission.SubmissionDate,
			&submission.CreatedAt, &submission.UpdatedAt,
			&submission.DeletedAt)

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

func (u *SubmissionRepo) GetRecentAcceptedSubmissions(user_id string) (*[]model.Submission, error) {
	submissions := []model.Submission{}
	query := `
	select * 
	from 
		submissions
	where
		deleted_at is null and user_id = $1 and submission_status = 'Accepted'
	`
	rows, err := u.Db.Query(query, user_id)
	for rows.Next() {
		submission := model.Submission{}
		err := rows.Scan(&submission.Id, &submission.ProblemTitle,
			&submission.UserUsername, &submission.LanguageName, &submission.Code,
			&submission.SubmissionStatus, &submission.SubmissionDate,
			&submission.CreatedAt, &submission.UpdatedAt,
			&submission.DeletedAt)

		if err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}

	return &submissions, err
}

// Update
func (u *SubmissionRepo) UpdateSubmission(submission *model.Submission) error {
	tx, err := u.Db.Begin()
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
		submission_date=$3
		updated_at=$4
	where 
		deleted_at is null and id = $5 `
	_, err = tx.Exec(query, submission.Code, submission.SubmissionStatus,
		submission.SubmissionDate, time.Now(), submission.Id)

	return err
}

// Delete
func (u *SubmissionRepo) DeleteSubmission(id string) error {
	tx, err := u.Db.Begin()
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
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
