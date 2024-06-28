package main

import (
	// "leetcode/generator"

	"database/sql"
	"fmt"
	"leetcode/models"
	"leetcode/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// generator.GenerateAllMockData(db)

	// h := handler.NewHandler(db)
	// server := router.CreateServer(h)
	// server.ListenAndServe()

	// ProblemsTesting(db)
}

func ProblemsTesting(db *sql.DB) {
	// problems := postgres.NewProblemRepo(db)

	// FilterProblems(*problems)

	// GetProblemById(*problems)

	// PickRandomProblem(*problems)

	// UpdateProblem(*problems)

	// DeleteProblem(*problems)

	// RecoverProblem(*problems)
}

func RecoverProblem(problems postgres.ProblemRepo) {
	err := problems.RecoverProblem("89adc799-4880-449e-ad42-b4946c396f69")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Recovered successfully")
}

func DeleteProblem(problems postgres.ProblemRepo) {
	err := problems.DeleteProblem("89adc799-4880-449e-ad42-b4946c396f69")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted successfully")
}

func UpdateProblem(problems postgres.ProblemRepo) {
	pro := models.ProblemUpdate{
		Id:            "89adc799-4880-449e-ad42-b4946c396f69",
		Title:         "container-with-most-water",
		ProblemNumber: 11,
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
	err := problems.UpdateProblem(&pro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated successfully")
}

func PickRandomProblem(problems postgres.ProblemRepo) {
	problem, err := problems.PickRandomProblem()
	if err != nil {
		log.Fatal(err)
	}
	PrintProblem(*problem)
}

func GetProblemById(problems postgres.ProblemRepo) {
	problem, err := problems.GetProblemById("79cb0553-226c-4368-b3fb-dc2b5f3b74ab")
	if err != nil {
		log.Fatal(err)
	}
	PrintProblem(*problem)
}

func PrintProblem(problem models.Problem) {
	fmt.Println("ID = ", problem.Id)
	fmt.Println("Problem number = ", problem.ProblemNumber)
	fmt.Println("Title = ", problem.Title)
	fmt.Println("Difficulty = ", problem.Difficulty)
	fmt.Println("Description = ", problem.Description)
	fmt.Println("Examples:")
	for _, example := range problem.Examples {
		fmt.Println("\tinput = ", example.Input)
		fmt.Println("\toutput = ", example.Output)
		fmt.Println("\texamples = ", example.Explanation, "\n")
	}
	fmt.Println("Constraints = ", problem.Constraints)
	fmt.Println("Topics = ", problem.Topics)
	fmt.Println("Hints = ", problem.Hints)
	fmt.Println("CreateAT = ", problem.CreatedAt)
	fmt.Println("DeletedAt = ", problem.UpdatedAt)
}

func GetSubmissionStatisticsByProblemId() {
	// tem, err := problems.GetSubmissionStatisticsByProblemId("79cb0553-226c-4368-b3fb-dc2b5f3b74ab") // zigzag-conversion
	// fmt.Println(err)
	// fmt.Println(tem)
}

func GetProblemByIds() {
	// problem, err := problems.GetProblemById("")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(problem)
}

func FilterProblems(problems postgres.ProblemRepo) {
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
	problemsGr, err := problems.GetProblems("d490e243-22df-4d17-b0bd-13887fda6e59", &models.ProblemFilter{
		// Sorting:     &ascProblemNumber,
		// Search:      &search,
		// Status:      &notStarted,
		// Difficulty:  &difficulty,
		// TopicsSlugs: &topics,
		Limit:  &limit,
		Offset: &offset,
	})
	if err != nil {
		panic(err)
	}
	for _, val := range *problemsGr {
		fmt.Println(val)
	}
}

func CreateProblem() {
	// pro := model.Problem{
	// 	Title:         "contaidfgners-with-most-waterr",
	// 	ProblemNumber: 11,
	// 	Difficulty:    "Medium",
	// 	Description:   `You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).\n\nFind two lines that together with the x-axis form a container, such that the container contains the most water.\n\nReturn the maximum amount of water a container can store.\n\nNotice that you may not slant the container.`,
	// 	Hints: []string{
	// 		"If you simulate the problem, it will be O(n^2) which is not efficient.",
	// 		"Try to use two-pointers. Set one pointer to the left and one to the right of the array. Always move the pointer that points to the lower line.",
	// 		"How can you calculate the amount of water at each step?",
	// 	},
	// 	Constraints: []string{
	// 		"n == height.length",
	// 		"2 <= n <= 105",
	// 		"0 <= height[i] <= 104",
	// 	},
	// }
	// id, err := problems.CreateProblem(&pro)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(id)
}
