package model

type TopicProblem struct {
	Id        string
	TopicId   string
	ProblemId string
	Time
}

type TopicProblemFilter struct {
	TopicId   *string
	ProblemId *string
}

type Skill struct {
	SkillName         string `json:"skill_name"`
	NumberOfTimesUsed int    `json:"number_of_times_used"`
}
