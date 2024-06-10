package model

type Topic struct {
	Id   string
	Name string
	Time
}

type TopicFilter struct {
	Name *string
}
