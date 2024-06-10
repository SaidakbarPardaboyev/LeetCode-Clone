package postgres

import (
	"database/sql"
	"leetcode/model"
	"time"
)

type LanguageRepo struct {
	Db *sql.DB
}

func NewLanguageRepo(db *sql.DB) *LanguageRepo {
	return &LanguageRepo{db}
}

// Create
func (l *LanguageRepo) CreateLanguage(language model.Language) error {

	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := "insert into languages(name) values($1)"
	_, err = tx.Exec(query, language.Name)

	return err
}

// Read
func (l *LanguageRepo) GetLanguageById(id string) (model.Language, error) {
	language := model.Language{}
	query := `
	select * from languages
	where
		id = $1 and deleted_at is null
	`
	row := l.Db.QueryRow(query, id)
	err := row.Scan(&language.Id, &language.Name, &language.Created_at, &language.Updated_at, &language.Deleted_at)
	return language, err
}
func (l *LanguageRepo) GetLanguages(filter model.LanguageFilter) (*[]model.Language, error) {
	params := []interface{}{}
	query := `
	select * from languages where deleted_at is null`
	if filter.Name != nil {
		query += " and name=$1"
		params = append(params, *filter.Name)
	}

	rows, err := l.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	languages := []model.Language{}
	for rows.Next() {
		language := model.Language{}
		err = rows.Scan(&language.Id, &language.Name, &language.Created_at, &language.Updated_at, &language.Deleted_at)
		if err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &languages, nil
}

// Update
func (l *LanguageRepo) UpdateLanguage(language model.Language) error {
	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update languages 
	set 
		name = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, language.Name, time.Now(), language.Id)

	return err
}

// Delete
func (l *LanguageRepo) DeleteLanguage(id string) error {
	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update languages 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
