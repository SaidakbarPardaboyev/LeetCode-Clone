package models

import (
	"database/sql"
)

type Problem struct {
	Id            string `json:"id"`
	ProblemNumber int    `json:"problem_number"`
	Title         string `json:"title"`
	Difficulty    string `json:"difficulty"`
	Description   string `json:"description"`
	Examples      []Example
	Constraints   []string `json:"constraints"`
	Topics        []string
	Hints         []string `json:"hints"`
	Time
}

type ProblemUpdate struct {
	Id            string `json:"id"`
	ProblemNumber int    `json:"problem_number"`
	Title         string `json:"title"`
	Difficulty    string `json:"difficulty"`
	Description   string `json:"description"`
	Constraints   []string `json:"constraints"`
	Hints         []string `json:"hints"`
}

type ProblemSet struct {
	Id            string         `json:"id"`
	Status        string         `json:"status"`
	ProblemNumber int            `json:"problem_number"`
	Title         string         `json:"title"`
	Acceptence    sql.NullString `json:"acceptence"`
	Difficulty    string         `json:"difficulty"`
}

type SubmissionStatisticsOfProblem struct {
	Accepted       int     `json:"accepted"`
	Submissions    int     `json:"submissions"`
	AcceptanceRate float32 `json:"acceptence_rate"`
}

type ProblemFilter struct {
	Sorting     *string
	Status      *string
	Search      *string
	TopicsSlugs *string
	Difficulty  *string
	Limit       *int
	Offset      *int
}
