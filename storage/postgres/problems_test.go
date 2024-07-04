package postgres

import (
	"database/sql"
	"leetcode/models"
	"testing"

	"github.com/google/uuid"
)

func TestCreateProblem(t *testing.T) {
	newProblem := models.ProblemCreate{
		Title:       "Saidakbar",
		Difficulty:  "Hard",
		Description: "CHiiiiiii",
		Constraints: []string{},
		Hints:       []string{},
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	id, err := NewProblemRepo(db).CreateProblem(&newProblem)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}

	if _, err := uuid.Parse(id); err != nil {
		t.Error(err)
	}
}

func TestGetProblemById(t *testing.T) {
	id := "af7dcc00-1d24-4e16-b1c3-dc6718bdd799"

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewProblemRepo(db).GetProblemById(id)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}
