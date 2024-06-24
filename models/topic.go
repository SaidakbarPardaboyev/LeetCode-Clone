package models

type Topic struct {
	Id   string
	Name string
	Time
}

type CreateUpdateTopic struct {
	Name string
}

type TopicFilter struct {
	Name   *string
	Limit  *int
	Offset *int
}
