package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"leetcode/models"
	"time"
)

type LanguageRepo struct {
	Db *sql.DB
}

func NewLanguageRepo(db *sql.DB) *LanguageRepo {
	return &LanguageRepo{db}
}

// Create
func (l *LanguageRepo) CreateLanguage(language *models.Language) (string, error) {

	tx, err := l.Db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Commit()
	id := uuid.NewString()
	query := "insert into languages(id, name) values($1)"
	_, err = tx.Exec(query, id, language.Name)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return id, err
}

// Read
func (l *LanguageRepo) GetLanguageById(id string) (models.Language, error) {
	language := models.Language{}
	query := `
	select 
		id, name, created_at, updated_at
	from 
	    languages
	where
		id = $1 and deleted_at is null
	`
	row := l.Db.QueryRow(query, id)
	err := row.Scan(&language.Id, &language.Name, &language.CreatedAt,
		&language.UpdatedAt)

	return language, err
}
func (l *LanguageRepo) GetLanguages(filter *models.LanguageFilter) (*[]models.Language, error) {
	params := []interface{}{}
	query := `
	select 
	    id, name, created_at, updated_at 
	from 
	    languages 
	where 
	    deleted_at is null`
	if filter.Name != nil {
		query += " and name=$1"
		params = append(params, *filter.Name)
	}

	rows, err := l.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	languages := []models.Language{}
	for rows.Next() {
		language := models.Language{}
		err = rows.Scan(&language.Id, &language.Name, &language.CreatedAt,
			&language.UpdatedAt)
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
func (l *LanguageRepo) UpdateLanguage(language *models.Language) error {
	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}

	query := `
	update 
	    languages 
	set 
		name = $1
		updatedAt = now()
	where 
		deleted_at is null and id = $2 `
	result, err := tx.Exec(query, language.Name, language.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing updated")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

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
	result, err := tx.Exec(query, time.Now(), id)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing deleted")
	}

	return err
}
