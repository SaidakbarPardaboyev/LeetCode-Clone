package postgres

import (
	"database/sql"
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
func (l *TopicRepo) CreateTopic(topic model.Topic) error {

	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := "insert into topics(name) values($1)"
	_, err = tx.Exec(query, topic.Name)

	return err
}

// Read
func (l *TopicRepo) GetTopicById(id string) (model.Topic, error) {
	topic := model.Topic{}
	query := `
	select * from topics
	where
		id = $1 and deleted_at is null
	`
	row := l.Db.QueryRow(query, id)
	err := row.Scan(&topic.Id, &topic.Name, &topic.Created_at, &topic.Updated_at, &topic.Deleted_at)
	return topic, err
}
func (l *TopicRepo) GetTopics(filter model.TopicFilter) (*[]model.Topic, error) {
	params := []interface{}{}
	query := `
	select * from topics where deleted_at is null`
	if filter.Name != nil {
		query += " and name=$1"
		params = append(params, *filter.Name)
	}

	rows, err := l.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	languages := []model.Topic{}
	for rows.Next() {
		topic := model.Topic{}
		err = rows.Scan(&topic.Id, &topic.Name, &topic.Created_at, &topic.Updated_at, &topic.Deleted_at)
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
func (t *TopicRepo) UpdateTopic(topic model.Topic) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics 
	set 
		name = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, topic.Name, time.Now(), topic.Id)

	return err
}

// Delete
func (t *TopicRepo) DeleteTopic(id string) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
