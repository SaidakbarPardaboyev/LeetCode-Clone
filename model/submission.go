package model

import "time"

type Submission struct {
	Id               int
	ProblemTitle     string
	UserUsername     string
	LanguageName     string
	Code             string
	SubmissionStatus string
	SubmissionDate   time.Time
	Time
}

type SubmissionFilter struct {
	ProblemId        *string
	UserId           *string
	LanguageId       *string
	Code             *string
	SubmissionStatus *string
}
