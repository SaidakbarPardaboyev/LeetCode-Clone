package generator

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"leetcode/storage/postgres"
	"math/rand"
	"time"
)

var (
	firstNames = []string{
		"John", "Jane", "Michael", "Emily", "David", "Sarah", "James", "Emma", "Robert", "Olivia",
		"William", "Sophia", "Joseph", "Isabella", "Charles", "Ava", "Thomas", "Mia", "Richard", "Charlotte",
		"Daniel", "Amelia", "Matthew", "Ella", "Andrew", "Grace", "Edward", "Chloe", "George", "Lucas",
		"Frank", "Lily", "Henry", "Madison", "Jack", "Harper", "Alice", "Evelyn", "Samuel", "Ethan",
		"Benjamin", "Alexander", "Leo", "Mason", "Lucas", "Luke", "Dylan", "Noah", "Logan",
	}

	lastNames = []string{
		"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor",
		"Anderson", "Thomas", "Jackson", "White", "Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson",
		"Clark", "Rodriguez", "Lewis", "Lee", "Walker", "Hall", "Allen", "Young", "Hernandez", "King",
		"Wright", "Lopez", "Hill", "Scott", "Green", "Adams", "Baker", "Gonzalez", "Nelson", "Carter",
		"Mitchell", "Perez", "Roberts", "Turner", "Phillips", "Campbell", "Parker", "Evans", "Edwards", "Collins",
	}

	usernames = []string{
		"johndoe", "janedoe", "michael92", "emily84", "david33", "sarah7", "james_smith", "emma_w", "robert1980", "olivia_m",
		"william_22", "sophia_17", "joseph29", "isabella_c", "charles1985", "ava_grace", "thomas_007", "mia_rose", "richard_81", "charlotte88",
		"daniel_m", "amelia_j", "matthew91", "ella.brown", "andrew1995", "grace_29", "edward_m", "chloe.03", "george_87", "lucas_f",
		"frank00", "lily03", "henry05", "madison_12", "jack.09", "harper11", "alice_w", "evelyn_c", "samuel.j", "ethan34",
		"benjamin_67", "alexander_88", "leo_r", "mason_99", "lucas_g", "luke_h", "dylan.m", "noah123", "logan_007",
	}

	bios = []string{
		"Software Engineer passionate about creating impactful applications.",
		"Student exploring the world of algorithms and data structures.",
		"Tech enthusiast with a keen interest in machine learning and AI.",
		"Web developer specializing in front-end technologies.",
		"Backend developer building scalable and robust APIs.",
		"Designer transforming ideas into beautiful user interfaces.",
		"Entrepreneur building the next big thing in fintech.",
		"Artist expressing creativity through digital media.",
		"Musician exploring new sounds and rhythms.",
		"Photographer capturing moments that tell a story.",
	}
)

func GenerateUsers() []model.User {
	rand.Seed(time.Now().UnixNano())

	var users []model.User

	for i := 0; i < 10000; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		username := fmt.Sprintf("%s%d", usernames[rand.Intn(len(usernames))], i)
		bio := bios[rand.Intn(len(bios))]

		fullName := fmt.Sprintf("%s %s", firstName, lastName)

		user := model.User{
			FullName: fullName,
			Username: username,
			Bio:      bio,
		}

		users = append(users, user)
	}

	return users
}

func InsertUsers(db *sql.DB) {
	users := GenerateUsers()
	u := postgres.NewUserRepo(db)
	for _, user := range users {
		u.CreateUser(user)
	}
}