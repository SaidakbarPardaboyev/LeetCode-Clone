package model

type Language struct {
	Id   string
	Name string
	Time
}

type LanguageFilter struct {
	Name *string
}
