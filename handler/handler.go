package handler

import (
	"database/sql"
	"leetcode/storage/postgres"
)

type Handler struct {
	UserRepo         *postgres.UserRepo
	SubmissionRepo   *postgres.SubmissionRepo
	LanguageRepo     *postgres.LanguageRepo
	TopicRepo        *postgres.TopicRepo
}

func NewHandler(db *sql.DB) *Handler {
	u := postgres.NewUserRepo(db)
	s := postgres.NewSubmissionRepo(db)
	l := postgres.NewLanguageRepo(db)
	t := postgres.NewTopicRepo(db)

	return &Handler{
		UserRepo:         u,
		SubmissionRepo:   s,
		LanguageRepo:     l,
		TopicRepo:        t,
	}
}



