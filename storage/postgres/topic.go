package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/models"
	"time"

	"github.com/google/uuid"
)

type TopicRepo struct {
	Db *sql.DB
}

func NewTopicRepo(db *sql.DB) *TopicRepo {
	return &TopicRepo{db}
}

// Create
func (t *TopicRepo) CreateTopic(topic *models.CreateTopic) (string, error) {

	tx, err := t.Db.Begin()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	query := "insert into topics(id, name) values($1, $2)"
	_, err = tx.Exec(query, id, topic.Name)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()

	return id, err
}

// Read
func (t *TopicRepo) GetTopicById(id string) (models.Topic, error) {
	topic := models.Topic{}
	query := `
	select 
		id, name, created_at, updated_at 
	from 
		topics
	where
		id = $1 and deleted_at is null
	`
	row := t.Db.QueryRow(query, id)
	err := row.Scan(&topic.Name, &topic.CreatedAt, &topic.UpdatedAt)
	return topic, err
}

func (t *TopicRepo) GetTopics(filter *models.TopicFilter) (*[]models.Topic, error) {
	params := []interface{}{}
	query := `
	select 
		id, name, created_at, updated_at 
	from 
		topics 
	where 
		deleted_at is null`

	if filter.Name != nil {
		query += " and name=$1"
		params = append(params, *filter.Name)
	}

	rows, err := t.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	languages := []models.Topic{}
	for rows.Next() {
		topic := models.Topic{}
		err = rows.Scan(&topic.Name, &topic.CreatedAt, &topic.UpdatedAt)
		if err != nil {
			return nil, err
		}
		languages = append(languages, topic)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &languages, nil
}

// Update
func (t *TopicRepo) UpdateTopic(topic *models.UpdateTopic) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 	
		topics 
	set 
		name = $1
	where 
		deleted_at is null and name = $2 `
	result, err := tx.Exec(query, topic.Name, time.Now(), topic.Name)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing updated")
	}

	return err
}

// Delete
func (t *TopicRepo) DeleteTopic(id string) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 
		topics 
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
