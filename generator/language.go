package generator

import (
	"database/sql"
	"leetcode/model"
	"leetcode/storage/postgres"
)

var languages = []string{"Go", "Python3", "C", "C++", "C#", "Rust", "Java", "JavaScript", "Swift", "Kotlin", "PHP"}

func InsertLanguages(db *sql.DB) {
	l := postgres.NewLanguageRepo(db)
	for _, lang := range languages {
		lm := model.Language{Name: lang}
		l.CreateLanguage(&lm)
	}
}
