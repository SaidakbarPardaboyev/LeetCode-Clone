package models

type Language struct {
	Id   string
	Name string
	Time
}

type CreateUpdateLanguage struct {
	Name string
}

type LanguageFilter struct {
	Name   *string
	Limit  *int
	Offset *int
}

type UsedLanguage struct {
	Name              string `json:"name"`
	NumberOfTimesUsed int    `json:"number_of_times_used"`
}
