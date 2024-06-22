package main

import (
	// "leetcode/generator"

	"fmt"
	"leetcode/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// generator.GenerateAllMockData(db)

	// mockUser := model.User{
	// 	Username:     "john_doe",
	// 	FullName:     "John Doe",
	// 	Email:        "john@example.com",
	// 	Password:     "securepassword",
	// 	ProfileImage: []byte{},
	// 	Gender:       "Male",
	// 	Location:     "New York",
	// 	Birthday:     time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	// 	Summary:      "A passionate software developer.",
	// 	Website:      "https://johndoe.com",
	// 	Github:       "https://github.com/johndoe",
	// 	LinkedIn:     "https://linkedin.com/in/johndoe",
	// }
	// fmt.Println("EasySolved", stats.EasySolved, "\n",
	// "MediumSolved",stats.MediumSolved,"\n",
	// "HardSolved",stats.HardSolved,"\n",
	// "TotalSolved",stats.TotalSolved,"\n",
	// "EasyUnsolved",stats.EasyUnsolved,     "\n",
	// "MediumUnsolved",stats.MediumUnsolved ,     "\n",
	// "HardUnsolved" , stats.EasyUnsolved,     "\n",
	// "TotalUnsolved" , stats.TotalUnsolved,     "\n",
	// "EasyAcceptanceRate", stats.EasyAcceptanceRate,  "\n",
	// "MediumAcceptanceRate",stats.MediumAcceptanceRate,"\n",
	// "HardAcceptanceRate",  stats.HardAcceptanceRate,"\n",
	// "TotalAcceptanceRate",  stats.TotalAcceptanceRate,)

	// skills, err := u.GetLanguagesWithNumberOfAcceptedProblemsByUsername("stark")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(skills)
	// h := handler.NewHandler(db)
	// server := router.CreateServer(h)
	// server.ListenAndServe()

	problems := postgres.NewProblemRepo(db)

	tem, err := problems.GetSubmissionStatisticsByProblemTitle("zigzag-conversion")
	fmt.Println(err)
	fmt.Println(tem)

	// pro := model.Problem{
	// 	Title:         "containers-with-most-water",
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
	// if err := problems.DeleteProblem("container-with-most-water"); err != nil {
	// 	panic(err)
	// }

	// problem, err := problems.GetProblemByTitle("container-with-most-water")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(problem)

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
	// difficulty := "Medium"

	// get all by Filter (topics)
	// topics := "string"
	// topics := "string%2Cdynamic-programming"

	// withErrorQuery := "dfghbf"
	// nonsort := "W3t9XQ%3D%3D"
	// limit := 11
	// offset := 0
	// problemsGr, err := problems.GetProblems("jdoe", &model.ProblemFilter{
	// 	Sorting:     &acASC,
	// 	Search:      &search,
	// 	Status:      &tried,
	// 	Difficulty:  &difficulty,
	// 	TopicsSlugs: &topics,
	// 	Limit:       &limit,
	// 	Offset:      &offset,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// for _, val := range *problemsGr {
	// 	fmt.Println(val)
	// }
}
