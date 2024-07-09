package models

type Example struct {
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation"`
}

type ExampleCreate struct {
	ProblemId   string `json:"problem_id"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation"`
}

type ExampleUpdate struct {
	Id          string `json:"id"`
	ProblemId   string `json:"problem_id"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation"`
}
