package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type TopicRepo struct {
	Db *sql.DB
}

func NewTopicRepo(db *sql.DB) *TopicRepo {
	return &TopicRepo{db}
}

// Create
func (t *TopicRepo) CreateTopic(topic *model.Topic) error {

	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := "insert into topics(name) values($1)"
	_, err = tx.Exec(query, topic.Name)

	return err
}

// Read
func (t *TopicRepo) GetTopicByName(name string) (model.Topic, error) {
	topic := model.Topic{}
	query := `
	select * 
	from 
		topics
	where
		name = $1 and deleted_at is null
	`
	row := t.Db.QueryRow(query, name)
	err := row.Scan(&topic.Name, &topic.CreatedAt,
		&topic.UpdatedAt, &topic.DeletedAt)
	return topic, err
}
func (t *TopicRepo) GetTopics(filter *model.TopicFilter) (*[]model.Topic, error) {
	params := []interface{}{}
	query := `
	select * from topics where deleted_at is null`
	if filter.Name != nil {
		query += " and name=$1"
		params = append(params, *filter.Name)
	}

	rows, err := t.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	languages := []model.Topic{}
	for rows.Next() {
		topic := model.Topic{}
		err = rows.Scan(&topic.Name, &topic.CreatedAt,
			&topic.UpdatedAt, &topic.DeletedAt)
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
func (t *TopicRepo) UpdateTopic(topic *model.Topic) error {
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
