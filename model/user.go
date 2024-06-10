package model

type User struct {
	Id       string
	FullName string
	Username string
	Bio      string
	Time
}

type UserFilter struct {
	FullName *string
	Username *string
}
