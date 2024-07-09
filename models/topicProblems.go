package models

type TopicsOfProblem struct {
	ProblemId  string   `json:"problem_id"`
	TopicNames []string `json:"topics"`
}

type TopicProblemCreate struct {
	ProblemId string `json:"problem_id"`
	TopicId   string `json:"topic_id"`
}

type TopicProblemUpdate struct {
	Id        string `json:"id"`
	ProblemId string `json:"problem_id"`
	TopicId   string `json:"topic_id"`
}

type TopicProblemFilter struct {
	TopicId   *string
	ProblemId *string
}

type Skill struct {
	SkillName         string `json:"skill_name"`
	NumberOfTimesUsed int    `json:"number_of_times_used"`
}
