package models

<<<<<<< HEAD
type TopicProblem struct {
	Id        string
	TopicId   string
	ProblemId string
	Time
=======
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
>>>>>>> origin/Saidakbar
}

type TopicProblemFilter struct {
	TopicId   *string
	ProblemId *string
}

type Skill struct {
	SkillName         string `json:"skill_name"`
	NumberOfTimesUsed int    `json:"number_of_times_used"`
}
