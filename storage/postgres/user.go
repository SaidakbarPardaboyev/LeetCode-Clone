package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Create
func (u *UserRepo) CreateUser(user model.User) error{
	
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := "insert into users(full_name, username, bio) values($1, $2, $3)"
	_, err = tx.Exec(query, user.FullName, user.Username, user.Bio)

	return err
}

// Read
func (u *UserRepo) GetUserById(id string) (model.User, error){
	user := model.User{}
	query := `
	select * from users
	where
		id = $1 and deleted_at is null
	`
	row := u.Db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.FullName, &user.Username, &user.Bio, &user.Created_at, &user.Updated_at, &user.Deleted_at)
	return user, err
}
func (u *UserRepo) GetUsers(filter model.UserFilter) (*[]model.User, error){
	params := []interface{}{}
	paramCount := 1
	query := `
	select * from users where deleted_at is null`
	if filter.FullName != nil{
		query += fmt.Sprintf(" and full_name=$%d", paramCount)
		params = append(params, *filter.FullName)
		paramCount++
	}
	if filter.Username != nil{
		query += fmt.Sprintf(" and username=$%d", paramCount)
		params = append(params, *filter.Username)
		paramCount++
	}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	users := []model.User{}
	for rows.Next(){
		user := model.User{}
		err = rows.Scan(&user.Id, &user.FullName, &user.Username, &user.Bio, &user.Created_at, &user.Updated_at, &user.Deleted_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil{
		return nil, err
	}

	return &users, nil
}

// Update
func (u *UserRepo) UpdateUser(user model.User) error{
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update users 
	set 
		full_name = $1, 
		username = $2,
		bio = $3,
		updated_at = $4
	where 
		deleted_at is null and id = $5 `
	_, err = tx.Exec(query, user.FullName, user.Username, user.Bio, time.Now(), user.Id)

	return err
}

// Delete
func (u *UserRepo) DeleteUser(id string) error{
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `update users 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	_, err = tx.Exec(query, time.Now(), id)

	return err
}
