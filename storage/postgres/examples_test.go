package postgres

import (
	"database/sql"
	"leetcode/models"
	"testing"
)

func TestCreateExample(t *testing.T) {
	NewExample := models.ExampleCreate{
		ProblemId:   "2396a6bb-ef10-4176-a16c-6f3183ea0db2",
		Input:       "Saidakbar",
		Output:      "Pardaboyev",
		Explanation: "Saidakbar Pardaboyev",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	_, err = NewExampleRepo(db).CreateExample(&NewExample)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestGetExamplesByProblemId(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	_, err = NewExampleRepo(db).GetExamplesByProblemId(
		"2396a6bb-ef10-4176-a16c-6f3183ea0db2")
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestUpdateExample(t *testing.T) {
	NewExample := models.ExampleUpdate{
		Id:          "bfaa241a-5b02-4cb6-a9ab-68fb6200bad3",
		ProblemId:   "2396a6bb-ef10-4176-a16c-6f3183ea0db2",
		Input:       "Muhammadjon",
		Output:      "Ko'palov",
		Explanation: "Saidakbar Pardaboyev",
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = NewExampleRepo(db).UpdateExample(
		&NewExample)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestDeleteExample(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = NewExampleRepo(db).DeleteExample(
		"bfaa241a-5b02-4cb6-a9ab-68fb6200bad3")
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestRecoverExample(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = NewExampleRepo(db).RecoverExample(
		"bfaa241a-5b02-4cb6-a9ab-68fb6200bad3")
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}