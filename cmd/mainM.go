package main

import (
	// "leetcode/generator"

	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"os"
	"reflect"
)

type ResultValues struct {
	Isaccepted bool
	Status     string
	Output     string
	Result     reflect.Value
	TestcaseId string
	RunTime    float64
}

func main() {
	// db, err := postgres.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// generator.GenerateAllMockData(db)

	res := ResultValues{
		Isaccepted: false,
		Status:     "time limit exceeded",
		Output:     "output",
		Result:     reflect.Value{},
		TestcaseId: "",
		RunTime:    0,
	}
	result := bytes.Buffer{}
	json.NewEncoder(&result).Encode(res)

	fmt.Println(result.String())

	// generator.GenerateAllMockData(db)
	// u := postgres.NewUserRepo(db)

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
}

// imagePath := "profile/image.png"
// imageBytes, err := ioutil.ReadFile(imagePath)
//
//	if err != nil {
//		log.Fatal("Error reading file:", err)
//	}
func saveImage(profileImage []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(bytes.NewReader(profileImage))
	if err != nil {
		return err
	}

	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}
