package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"

	"github.com/lib/pq"
)

type TopicProblemRepo struct {
	Db *sql.DB
}

func NewTopicProblemRepo(db *sql.DB) *TopicProblemRepo {
	return &TopicProblemRepo{db}
}

// Create
func (l *TopicProblemRepo) CreateTopicProblem(tp model.TopicProblem) error {

	tx, err := l.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := "insert into topics_problems(topic_id, problem_id) values($1, $2)"
	_, err = tx.Exec(query, tp.TopicId, tp.ProblemId)

	return err
}

// Read
func (l *TopicProblemRepo) GetTopicProblemById(id string) (model.TopicProblem, error) {
	topicProblem := model.TopicProblem{}
	query := `
	select * from topics_problems
	where
		id = $1 and deleted_at is null
	`
	row := l.Db.QueryRow(query, id)
	err := row.Scan(&topicProblem.Id, &topicProblem.TopicId, &topicProblem.ProblemId, &topicProblem.Created_at, &topicProblem.Updated_at, &topicProblem.Deleted_at)
	return topicProblem, err
}

func (l *TopicProblemRepo) GetTopicProblems(filter model.TopicProblemFilter) (*[]model.TopicProblem, error) {
	params := []interface{}{}
	paramcount := 1
	query := `
	select * from topics_problems where deleted_at is null`
	if filter.TopicId != nil {
		query += fmt.Sprintf(" and topic_id=$%d", paramcount)
		params = append(params, *filter.ProblemId)
		paramcount++
	}
	if filter.ProblemId != nil {
		query += fmt.Sprintf(" and problem_id=$%d", paramcount)
		params = append(params, *filter.ProblemId)
		paramcount++
	}

	rows, err := l.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	topicProblems := []model.TopicProblem{}
	for rows.Next() {
		topicProblem := model.TopicProblem{}
		err = rows.Scan(&topicProblem.Id, &topicProblem.TopicId, &topicProblem.ProblemId, &topicProblem.Created_at,
			&topicProblem.Updated_at, &topicProblem.Deleted_at)
		if err != nil {
			return nil, err
		}
		topicProblems = append(topicProblems, topicProblem)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &topicProblems, nil
}
func (t *TopicProblemRepo) GetProblemsByTopicId(topicId string) (*[]model.Problem, error) {
	query := `
	select 
		p.id, p.question_number, p.title, p.difficulty_level, p.description, p.examples, p.hints, p.constraints
	from 
		topics_problems as tp
	join
		topics as t
	on 
		tp.topic_id = t.id and t.deleted_at is null
	join
		problems as p
	on 
		p.id = tp.problem_id and p.deleted_at is null
	where
		tp.topic_id = $1 and tp.deleted_at is null
	`

	rows, err := t.Db.Query(query, topicId)
	if err != nil {
		return nil, err
	}

	problems := []model.Problem{}
	for rows.Next() {
		problem := model.Problem{}
		err := rows.Scan(&problem.Id, &problem.QuestionNumber, &problem.Title, 
			&problem.DifficultyLevel, &problem.Description,
			pq.Array(&problem.Examples), pq.Array(&problem.Hints), pq.Array(&problem.Constraints))
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	err = rows.Err()

	return &problems, err
}

func (t *TopicProblemRepo) GetTopicsByProblemId(problemId string) (*[]model.Topic, error) {
	query := `
	select 
		t.id, t.name
	from 
		topics_problems as tp
	join
		topics as t
	on 
		tp.topic_id = t.id and t.deleted_at is null
	join
		problems as p
	on 
		p.id = tp.problem_id and p.deleted_at is null
	where
		tp.problem_id = 'c81c3b88-6937-47cc-9a8f-32f195911209' and tp.deleted_at is null
	`

	rows, err := t.Db.Query(query, problemId)
	if err != nil {
		return nil, err
	}

	topics := []model.Topic{}
	for rows.Next() {
		topic := model.Topic{}
		err := rows.Scan(&topic.Id, &topic.Name)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	err = rows.Err()

	return &topics, err
}

// Update
func (t *TopicProblemRepo) UpdateTopicProblem(tp model.TopicProblem) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update topics_problems 
	set 
		topic_id = $1,
		problem_id = $2
	where 
		deleted_at is null and id = $3`
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
