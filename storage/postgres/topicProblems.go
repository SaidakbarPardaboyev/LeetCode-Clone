package postgres

import (
	"database/sql"
	"leetcode/models"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type TopicProblemRepo struct {
	Db *sql.DB
}

func NewTopicProblemRepo(db *sql.DB) *TopicProblemRepo {
	return &TopicProblemRepo{db}
}

// Create
func (l *TopicProblemRepo) AddTopicToProblem(tp *models.TopicProblemCreate) (*string, error) {
	tx, err := l.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	newId := uuid.NewString()
	query := `insert into topics_problems(id, topic_id, problem_id)
	values($1, $2)`
	_, err = tx.Exec(query, newId, tp.TopicId, tp.ProblemId)
	if err != nil {
		return nil, err
	}

	return &newId, nil
}

// Create
func (l *TopicProblemRepo) AddTopicsToProblem(tp *models.TopicsOfProblem) error {
	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	for _, topic := range tp.TopicNames {
		newId := uuid.NewString()
		query := `insert into topics_problems(id, topic_id, problem_id)
		values($1, $2)`
		_, err = tx.Exec(query, newId, topic, tp.ProblemId)
		if err != nil {
			return err
		}
	}

	return nil
}

// Read
func (l *TopicProblemRepo) GetTopicsByProblemId(problemId string) (models.TopicsOfProblem, error) {
	topicProblem := models.TopicsOfProblem{}
	query := `
		select
			tp.problem_id,
			array_agg(t.name) as topics
		from
			topics as t
		inner join 
			topics_problems as tp 
				on t.id = tp.topic_id
		where
			tp.problem_id = $1 and
			t.deleted_at is null
		group by
			tp.problem_id;
	`
	err := l.Db.QueryRow(query, problemId).Scan(&topicProblem.ProblemId,
		pq.Array(&topicProblem.TopicNames))

	return topicProblem, err
}

// Update
func (t *TopicProblemRepo) UpdateTopicProblem(tp *models.TopicProblemUpdate) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics_problems
	set
		topic_id = $1,
		problem_id = $2,
		updated_at = $3
	where
		deleted_at is null and id = $4`
	_, err = tx.Exec(query, tp.TopicId, tp.ProblemId, time.Now(), tp.Id)

	return err
}

// Delete
func (t *TopicProblemRepo) DeleteTopicProblem(id string) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics_problems
	set
		deleted_at = $1
	where
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}

// Recover
func (t *TopicProblemRepo) RecoverTopicProblem(id string) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics_problems
	set
		deleted_at = null
	where
		deleted_at is not null and id = $1`
	_, err = tx.Exec(query, id)

	return err
}
