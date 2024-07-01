package models

type Topic struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Time
}

type CreateTopic struct {
	Name string `json:"name"`
}

type UpdateTopic struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TopicFilter struct {
	Name   *string `json:"name,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
	Offset *int    `json:"offset,omitempty"`
}
