package models

type Language struct {
	Id   string
	Name string
	Time
}

type LanguageFilter struct {
	Name *string
}

type UsedLanguage struct {
	Name              string `json:"name"`
	NumberOfTimesUsed int    `json:"number_of_times_used"`
}