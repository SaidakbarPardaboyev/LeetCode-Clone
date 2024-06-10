package model


type Problem struct {
	Id              string
	QuestionNumber  int
	Title           string
	DifficultyLevel string
	Description     string
	Examples        []string
	Hints           []string
	Constraints     []string
	Time
}

type ProblemFilter struct {
	QuestionNumber  *int
	Title           *string
	DifficultyLevel *string
}
