package generator

import (
	"database/sql"
)

// for creating all mock data
func GenerateAllMockData(db *sql.DB) {
	// fmt.Println(GenerateUsers())
	// InsertLanguages(db)
	// InsertTopics(db)
	// InsertUsers(db)
	// InsertProblems(db)
	InsertSubmissions(db)
	// InsertTopicProblems(db)
}
