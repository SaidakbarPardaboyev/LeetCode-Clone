package postgres

import (
	"fmt"
	"leetcode/models"
	"log"
	"testing"
	"time"
)

func newUserRepoTest() *UserRepo {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	return &UserRepo{Db: db}
}

// func TestCreateUser(t *testing.T) {
// 	u := newUserRepoTest()

// 	user := models.CreateUser{
// 		Username: "qwerty",
// 		Email: "qwerty@hmail.com",
// 		Password: "12345",
// 	}
// 	id, err := u.CreateUser(&user)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = uuid.Parse(id)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestGetUserById(t *testing.T) {
	u := newUserRepoTest()

	user, err := u.GetUserById("ef475ecd-21e2-4d6e-99a7-fa97ade13d92")
	if err != nil {
		panic(err)
	}
	if len(user.Email) < 3 {
		panic(fmt.Errorf("no email or invalid one"))
	}
}

func TestGetUsers(t *testing.T){
	u := newUserRepoTest()

	g := "male"
	filter := models.UserFilter{
		Gender: &g,
	}

	users, err := u.GetUsers(&filter)
	if err != nil {
		panic(err)
	}
	if len(*users) == 0{
		log.Fatal("invalid answer got 0 length ")
	} 
}


func TestGetUserRankingByUserId(t *testing.T){
	u := newUserRepoTest()

	rank, err := u.GetUserRankingByUserId("294925d6-3004-49e1-8b4c-3c7a2bf23530")
	if err != nil {
		panic(err)
	}
	if rank == 0 {
		log.Fatal("got no rank in result")
	}
}

func TestGetNumberOfSolvedProblemsByUserIdWithStats(t *testing.T){
	u := newUserRepoTest()

	stats, err := u.GetNumberOfSolvedProblemsByUserIdWithStats("294925d6-3004-49e1-8b4c-3c7a2bf23530")
	if err != nil {
		panic(err)
	}
	fmt.Println("EasySolved", stats.EasySolved, "\n",
	"MediumSolved",stats.MediumSolved,"\n",
	"HardSolved",stats.HardSolved,"\n",
	"TotalSolved",stats.TotalSolved,"\n",
	"EasyUnsolved",stats.EasyUnsolved,     "\n",
	"MediumUnsolved",stats.MediumUnsolved ,     "\n",
	"HardUnsolved" , stats.EasyUnsolved,     "\n",
	"TotalUnsolved" , stats.TotalUnsolved,     "\n",
	"EasyAcceptanceRate", stats.EasyAcceptanceRate,  "\n",
	"MediumAcceptanceRate",stats.MediumAcceptanceRate,"\n",
	"HardAcceptanceRate",  stats.HardAcceptanceRate,"\n",
	"TotalAcceptanceRate",  stats.TotalAcceptanceRate,)
	if stats.TotalSolved == 0 {
		log.Fatal("got no solved problebs of user in result")
	}
}

func TestGetSkillsByUserId(t *testing.T){
	u := newUserRepoTest()

	skills, err := u.GetSkillsByUserId("294925d6-3004-49e1-8b4c-3c7a2bf23530")
	if err != nil {
		panic(err)
	}

	if len(*skills) == 0 {
		log.Fatal("got no skills of user in result")
	}
}

func TestGetLanguagesWithNumberOfAcceptedProblemsByUserId(t *testing.T){
	u := newUserRepoTest()

	langs, err := u.GetLanguagesWithNumberOfAcceptedProblemsByUserId("294925d6-3004-49e1-8b4c-3c7a2bf23530")
	if err != nil {
		panic(err)
	}

	if len(*langs) == 0 {
		log.Fatal("got no skills of user in result")
	}
}

func TestUpdateUser(t *testing.T){
	u := newUserRepoTest()

	user := models.UpdateUser{
		Id:           "294925d6-3004-49e1-8b4c-3c7a2bf23530",
		Username:     "qale",
		FullName:     "gfd",
		Email:        "hbgfd",
		Password:     "hgfd",
		ProfileImage: []byte{},
		Gender:       "male",
		Location:     "hytr",
		Birthday:     time.Now(),
		Summary:      "iuyt",
		Website:      "iuytr",
		Github:       "iuytr",
		LinkedIn:     "iuyt76yt",
	}

	err := u.UpdateUser(&user)
	if err != nil {
		panic(err)
	}
}

func TestDeleteUser(t *testing.T){
	u := newUserRepoTest()

	err := u.DeleteUser("294925d6-3004-49e1-8b4c-3c7a2bf23530")
	if err != nil {
		panic(err)
	}
}

