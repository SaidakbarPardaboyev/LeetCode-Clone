package model

import (
	"database/sql"
	"time"
)

type Submission struct {
	Id               int             `json:"id"`
	ProblemTitle     string          `json:"problem_title"`
	UserUsername     string          `json:"user_username"`
	LanguageName     string          `json:"language_name"`
	Code             string          `json:"code"`
	SubmissionStatus string          `json:"submission_status"`
	Runtime          sql.NullFloat64 `json:"runtime"`
	SubmissionDate   time.Time       `json:"submission_date"`
	Time
}

type SubmissionFilter struct {
	ProblemTitle     *string
	UserUsername     *string
	LanguageName     *string
	Code             *string
	SubmissionStatus *string
	Runtime          *sql.NullFloat64
	SubmissionDate   *time.Time
	Limit            *int
	Offset           *int
}

type UserActivity struct {
	TotalSubmissions int             `json:"total_submissions"`
	SubmissionDays   []SubmissionDay `json:"submission_days"`
}

type SubmissionDay struct {
	Date            time.Time `json:"date"`
	SubmissionCount int       `json:"submission_count"`
}

type RecentlyAcceptedSubmission struct {
	ProblemTitle   string    `json:"problem_title"`
	SubmissionDate time.Time `json:"submission_date"`
}

type AllStatisticsOfSolvedProblems struct {
	EasySolved           int     `json:"easy"`
	MediumSolved         int     `json:"medium"`
	HardSolved           int     `json:"hard"`
	TotalSolved          int     `json:"total"`
	EasyUnsolved         int     `json:"easy_unsolved"`
	MediumUnsolved       int     `json:"medium_unsolved"`
	HardUnsolved         int     `json:"hard_unsolved"`
	TotalUnsolved        int     `json:"total_unsolved"`
	EasyAcceptanceRate   float64 `json:"easy_acceptance_rate"`
	MediumAcceptanceRate float64 `json:"medium_acceptance_rate"`
	HardAcceptanceRate   float64 `json:"hard_acceptance_rate"`
	TotalAcceptanceRate  float64 `json:"total_acceptance_rate"`
}
