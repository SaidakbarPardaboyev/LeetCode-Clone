package models

type TopicsProblem struct {
	ProblemId  string   `json:"problem_id"`
	TopicNames []string `json:"topics"`
}

type TopicProblemFilter struct {
	TopicId   *string
	ProblemId *string
}

type Skill struct {
	SkillName         string `json:"skill_name"`
	NumberOfTimesUsed int    `json:"number_of_times_used"`
}
