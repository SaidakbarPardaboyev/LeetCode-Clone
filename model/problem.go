package model

import "database/sql"

type Problem struct {
	Title         string   `json:"title"`
	ProblemNumber int      `json:"problem_number"`
	Difficulty    string   `json:"difficulty"`
	Description   string   `json:"description"`
	Constraints   []string `json:"constraints"`
	Hints         []string `json:"hints"`
	Time
}

type Problems struct {
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
