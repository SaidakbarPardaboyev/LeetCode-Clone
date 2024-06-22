package model

import "time"

type User struct {
	Username     string    `json:"username"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProfileImage []byte    `json:"profile_image"`
	Gender       string    `json:"gender"`
	Location     string    `json:"location"`
	Birthday     time.Time `json:"birthday"`
	Summary      string    `json:"summary"`
	Website      string    `json:"website,omitempty"`
	Github       string    `json:"github,omitempty"`
	LinkedIn     string    `json:"linkedin,omitempty"`
	Time
}
type UserFilter struct {
	FullName *string
	Email    *string
	Gender   *string
	Location *string
	AgeFrom  *int
	AgeTo    *int
	Limit    *int
	Offset   *int
}
