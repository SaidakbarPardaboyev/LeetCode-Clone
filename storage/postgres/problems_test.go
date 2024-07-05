package postgres

import (
	"database/sql"
	"leetcode/models"
	"testing"

	"github.com/google/uuid"
)

func TestCreateProblem(t *testing.T) {
	newProblem := models.ProblemCreate{
		Title:       "contaidfgners-with-most-waterr",
		Difficulty:  "Medium",
		Description: `You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).\n\nFind two lines that together with the x-axis form a container, such that the container contains the most water.\n\nReturn the maximum amount of water a container can store.\n\nNotice that you may not slant the container.`,
		Hints: []string{
			"If you simulate the problem, it will be O(n^2) which is not efficient.",
			"Try to use two-pointers. Set one pointer to the left and one to the right of the array. Always move the pointer that points to the lower line.",
			"How can you calculate the amount of water at each step?",
		},
		Constraints: []string{
			"n == height.length",
			"2 <= n <= 105",
			"0 <= height[i] <= 104",
		},
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

func TestGetProblems(t *testing.T) {
	// get all by Filter (sorting)
	// acASC := "W3sic29ydE9yZGVyIjoiREVTQ0VORElORyIsIm9yZGVyQnkiOiJBQ19SQVRFIn1d"
	// acDesc := "W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZXJCeSI6IkFDX1JBVEUifV0%3D"
	// difHardToEasy := "W3sic29ydE9yZGVyIjoiREVTQ0VORElORyIsIm9yZGVyQnkiOiJESUZGSUNVTFRZIn1d"
	// difEasyToHard := "W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZXJCeSI6IkRJRkZJQ1VMVFkifV0%3D"
	// ascProblemNumber := "W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZXJCeSI6IkZST05URU5EX0lEIn1d"
	// descProblemNumber := "W3sic29ydE9yZGVyIjoiREVTQ0VORElORyIsIm9yZGVyQnkiOiJGUk9OVEVORF9JRCJ9XQ%3D%3D"

	// get all by Filter (user's problems status [NOT_STARTED/AC/TRIED])
	// notStarted := "NOT_STARTED"
	// ac := "AC"
	// tried := "TRIED"

	// get all by Filter (searching)
	// search := "palindromic"

	// get all by Filter (problem status [only hard/medium/easy])
	// difficulty := "Hard"

	// get all by Filter (topics)
	// topics := "string%2Cstack"
	// topics := "string%2Cdynamic-programming"

	// withErrorQuery := "dfghbf"
	// nonsort := "W3t9XQ%3D%3D"
	limit := 10
	offset := 0

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewProblemRepo(db).GetProblems("d490e243-22df-4d17-b0bd-13887fda6e59", &models.ProblemFilter{
		// Sorting:     &ascProblemNumber,
		// Search:      &search,
		// Status:      &notStarted,
		// Difficulty:  &difficulty,
		// TopicsSlugs: &topics,
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestGetSubmissionStatisticsByProblemId(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewProblemRepo(db).GetSubmissionStatisticsByProblemId(
		"79cb0553-226c-4368-b3fb-dc2b5f3b74ab") // zigzag-conversion

	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestGetAllProblemsId(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewProblemRepo(db).GetProblemById(
		"79cb0553-226c-4368-b3fb-dc2b5f3b74ab")
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestPickRandomProblem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	_, err = NewProblemRepo(db).PickRandomProblem()
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestUpdateProblem(t *testing.T) {
	pro := models.ProblemUpdate{
		Id:            "2396a6bb-ef10-4176-a16c-6f3183ea0db2",
		Title:         "container-with-most-water-Saidakbar",
		ProblemNumber: 61,
		Difficulty:    "Medium",
		Description:   `You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).\n\nFind two lines that together with the x-axis form a container, such that the container contains the most water.\n\nReturn the maximum amount of water a container can store.\n\nNotice that you may not slant the container.`,
		Hints: []string{
			"If you simulate the problem, it will be O(n^2) which is not efficient.",
			"Try to use two-pointers. Set one pointer to the left and one to the right of the array. Always move the pointer that points to the lower line.",
			"How can you calculate the amount of water at each step?",
		},
		Constraints: []string{
			"n == height.length",
			"2 <= n <= 105",
			"0 <= height[i] <= 104",
		},
	}

	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	err = NewProblemRepo(db).UpdateProblem(&pro)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestDeleteProblem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	err = NewProblemRepo(db).DeleteProblem(
		"2396a6bb-ef10-4176-a16c-6f3183ea0db2")
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestRecoverProblem(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	err = NewProblemRepo(db).RecoverProblem(
		"2396a6bb-ef10-4176-a16c-6f3183ea0db2")
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}
