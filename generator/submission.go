package generator

import (
	"database/sql"
	"leetcode/model"
	"math/rand"
	"time"
)

var status = []string{
	"Accepted",
	"Run Time Error",
	"Compile Error",
	"Wrong Answer",
	"Time Limit Exceeded",
	"Memory Limit Exceeded",
	"Output Limit Exceeded",
}

func getProblemID(db *sql.DB) ([]string, error) {
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

func getUserId(db *sql.DB) ([]string, error) {
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

func getLanguageId(db *sql.DB) ([]string, error) {
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
	problemsId, err := getProblemID(db)
	if err != nil {
		panic(err)
	}

	userUsersId, err := getUserId(db)
	if err != nil {
		panic(err)
	}

	languagesId, err := getLanguageId(db)
	if err != nil {
		panic(err)
	}

	for _, userId := range userUsersId {
		for i := 0; i < 10; i++ {
			submission := model.Submission{
				ProblemId:        problemsId[rand.Intn(len(problemsId))],
				UserId:           userId,
				LanguageId:       languagesId[rand.Intn(len(languagesId))],
				Code:             randomString(50),
				SubmissionStatus: status[rand.Intn(len(status))],
				SubmissionDate:   randomDate(),
			}
			tx, err := db.Begin()
			if err != nil {
				panic(err)
			}
			defer tx.Commit()
			
			query := `insert into submissions(problem_id, user_id, language_id, 
			code, submission_status, submission_date)
			values($1, $2, $3, $4, $5, $6)`
			_, err = tx.Exec(query, submission.ProblemId, submission.UserId,
				submission.LanguageId, submission.Code, submission.SubmissionStatus,
				submission.SubmissionDate)
			if err != nil {
				panic(err)
			}
		}
	}
}
