package handler

import (
	"database/sql"
	"leetcode/storage/postgres"
)

type Handler struct {
	UserRepo         *postgres.UserRepo
	ProblemRepo      *postgres.ProblemRepo
	SubmissionRepo   *postgres.SubmissionRepo
	LanguageRepo     *postgres.LanguageRepo
	TopicRepo        *postgres.TopicRepo
	TopicProblemRepo *postgres.TopicProblemRepo
}

func NewHandler(db *sql.DB) *Handler {
	u := postgres.NewUserRepo(db)
	p := postgres.NewProblemRepo(db)
	s := postgres.NewSubmissionRepo(db)
	l := postgres.NewLanguageRepo(db)
	t := postgres.NewTopicRepo(db)
	tp := postgres.NewTopicProblemRepo(db)

	return &Handler{
		UserRepo:         u,
		ProblemRepo:      p,
		SubmissionRepo:   s,
		LanguageRepo:     l,
		TopicRepo:        t,
		TopicProblemRepo: tp,
	}
}



