package generator

import (
	"database/sql"
	"leetcode/model"
	"leetcode/storage/postgres"
	"math/rand"
	"time"
)

var status = []string{
	"Passed",
	"Run time Error",
	"Compile Error",
	"Wrong Answer",
	"Time Limit Exceeded",
	"Memory Limit Exceeded",
	"Output Limit Exceeded",
}

func getProblemIDs(db *sql.DB) ([]string, error) {
	query := `SELECT id FROM problems`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		problems = append(problems, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return problems, nil
}

func getUserIDs(db *sql.DB) ([]string, error) {
	query := `SELECT id FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		users = append(users, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func getLanguageIDs(db *sql.DB) ([]string, error) {
	query := `SELECT id FROM languages`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		languages = append(languages, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return languages, nil
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomDate() time.Time {
	min := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func InsertSubmissions(db *sql.DB) {
	problemIds, err := getProblemIDs(db)
	if err != nil {
		panic(err)
	}

	userIds, err := getUserIDs(db)
	if err != nil {
		panic(err)
	}

	languageIds, err := getLanguageIDs(db)
	if err != nil {
		panic(err)
	}

	s := postgres.NewSubmissionRepo(db)

	for i := 0; i < 10000; i++ {
		newSub := model.Submission{
			ProblemId:        problemIds[rand.Intn(len(problemIds))],
			UserId:           userIds[rand.Intn(len(userIds))],
			LanguageId:       languageIds[rand.Intn(len(languageIds))],
			Code:             randomString(50), 
			SubmissionStatus: status[rand.Intn(len(status))],
			SubmissionDate:   randomDate(),
		}
		if err := s.CreateSubmission(newSub); err != nil {
			panic(err)
		}
	}
}
