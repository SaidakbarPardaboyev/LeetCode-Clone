package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/model"
	"math"
	"time"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Create
func (u *UserRepo) CreateUser(user *model.User) error {

	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	insert into 
		users(username, full_name, email, password, profile_image, 
		gender, location, birthday, summary, website, github, linkedin) 
		values($1, $2, $3,$4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err = tx.Exec(query, user.Username, user.FullName, user.Email, user.Password,
		user.ProfileImage, user.Gender, user.Location, user.Birthday, user.Summary,
		user.Website, user.Github, user.LinkedIn)

	return err
}

// Read
func (u *UserRepo) GetUserByUsername(id string) (model.User, error) {
	user := model.User{}
	query := `
	select 
		username, full_name, email, password, profile_image, 
		gender, location, birthday, summary, website, github, linkedin,
		created_at, updated_at, deleted_at
	from 
		users
	where
		username = $1 and deleted_at is null
	`
	row := u.Db.QueryRow(query, id)
	err := row.Scan(&user.Username, &user.FullName, &user.Email,
		&user.Password, &user.ProfileImage, &user.Gender, &user.Location, &user.Birthday,
		&user.Summary, &user.Website, &user.Github, &user.LinkedIn,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	return user, err
}
func (u *UserRepo) GetUsers(filter *model.UserFilter) (*[]model.User, error) {
	params := []interface{}{}
	paramCount := 1
	query := `
	select 
		username, full_name, email, password, profile_image, 
		gender, location, birthday, summary, website, github, linkedin,
		created_at, updated_at, deleted_at 
	from 
		users 
	where 
		deleted_at is null`
	if filter.FullName != nil {
		query += fmt.Sprintf(" and full_name=$%d", paramCount)
		params = append(params, *filter.FullName)
		paramCount++
	}
	if filter.Email != nil {
		query += fmt.Sprintf(" and email=$%d", paramCount)
		params = append(params, *filter.Email)
		paramCount++
	}
	if filter.Gender != nil {
		query += fmt.Sprintf(" and gender=$%d", paramCount)
		params = append(params, *filter.Gender)
		paramCount++
	}
	if filter.Gender != nil {
		query += fmt.Sprintf(" and gender=$%d", paramCount)
		params = append(params, *filter.Gender)
		paramCount++
	}
	if filter.AgeFrom != nil {
		query += fmt.Sprintf(" and extract(year from age(current_timestamp, birthday))>=$%d", paramCount)
		params = append(params, *filter.AgeFrom)
		paramCount++
	}
	if filter.AgeTo != nil {
		query += fmt.Sprintf(" and extract(year from age(current_timestamp, birthday))<=$%d", paramCount)
		params = append(params, *filter.AgeTo)
		paramCount++
	}
	if filter.Limit != nil {
		query += fmt.Sprintf(" limit $%d", paramCount)
		params = append(params, *filter.Limit)
		paramCount++
	}
	if filter.Offset != nil {
		query += fmt.Sprintf(" offset $%d", paramCount)
		params = append(params, *filter.Offset)
		paramCount++
	}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}

	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Username, &user.FullName, &user.Email,
			&user.Password, &user.ProfileImage, &user.Gender, &user.Location, &user.Birthday,
			&user.Summary, &user.Website, &user.Github, &user.LinkedIn,
			&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserRepo) GetUserRankingByUsername(username string) (int, error) {
	//number of Solved Problems
	//difficulty of problem
	//average speed of solutions
	//acceptance rate
	query := `
	with all_solved_problems as(
		select distinct
			problem_title,
			user_username
		from
			submissions
		where
			submission_status = 'Accepted'
		group by
			problem_title, user_username
	),
	quality as(
		select
			als.user_username,
			count(case when p.difficulty = 'Easy' then 1 end) + 
            count(case when p.difficulty = 'Medium' then 1 end) * 3 + 
            count(case when p.difficulty = 'Hard' then 1 end) * 5 as quality_index,
            count(*) total_count
		from
			all_solved_problems as als
		join
			problems as p
		on
			p.title = als.problem_title
		group by
			als.user_username
	),
	speed_of_solutions as(
		select 
			user_username,
			round(avg(runtime)::numeric, 2) as speed
		from
			submissions
		where
			submission_status = 'Accepted'
		group by
			user_username
	),
	acceptance as (
		select 
			user_username,
			round(
				count(case when submission_status = 'Accepted' then 1 end)::numeric * 100 / count(*), 2
				) as acceptance_rate
		from
			submissions
		group by
			user_username
	),
	rankings as(
		select
			q.user_username,
			q.total_count,
			q.quality_index,
			sp.speed,
			a.acceptance_rate,
			dense_rank() over (order by q.quality_index desc, a.acceptance_rate desc, sp.speed asc, q.total_count desc) as rank
		from
			quality as q
		join
			speed_of_solutions as sp
		on
			sp.user_username = q.user_username
		join
			acceptance as a
		on
			a.user_username = q.user_username
	)
	select
		rank
	from
		rankings
	where 
		user_username = $1;
	`
	rank := 0
	row := u.Db.QueryRow(query, username)
	err := row.Scan(&rank)

	return rank, err
}

func (u *UserRepo) GetNumberOfSolvedProblemsByUsername(username string) (*model.AllStatisticsOfSolvedProblems, error) {
	unsolvedProblemsQuery := `
		select 
			count(case when difficulty = 'Easy' then 1 end) as easy,
            count(case when difficulty = 'Medium' then 1 end) as medium,
            count(case when difficulty = 'Hard' then 1 end) as hard,
			count(*) as total_unsolved
		from
			problems
	`
	query := `
	with solved_problems as (
		select distinct
			s.user_username,
			count(case when p.difficulty = 'Easy' then 1 end) as easy_solved,
			count(case when p.difficulty = 'Medium' then 1 end) as medium_solved,
			count(case when p.difficulty = 'Hard'  then 1 end) as hard_solved,
			count(*) total_solved
		from
			submissions as s
		join
			problems as p
		on
			s.submission_status = 'Accepted' and
			s.user_username = $1 and
			s.problem_title = p.title 
		group by
			s.user_username

	),
	tried_problems as(
		select
			s.user_username,
			count(case when p.difficulty = 'Easy'  then 1 end) as easy_tried,
			count(case when p.difficulty = 'Medium'  then 1 end) as medium_tried,
			count(case when p.difficulty = 'Hard'  then 1 end) as hard_tried,
			count(*) total_tried
		from
			submissions as s	
		join
			problems as p
		on	
			s.user_username = $1 and
			s.problem_title = p.title 
		group by
			s.user_username

	)
	select
		sp.easy_solved,
		sp.medium_solved,
		sp.hard_solved,
		sp.total_solved,
		tp.easy_tried,
		tp.medium_tried,
		tp.hard_tried,
		tp.total_tried
	from
		solved_problems as sp
	join	
		tried_problems as tp
	on	
		sp.user_username = tp.user_username
	`

	stats := model.AllStatisticsOfSolvedProblems{}

	row := u.Db.QueryRow(unsolvedProblemsQuery)

	err := row.Scan(&stats.EasyUnsolved, &stats.MediumUnsolved, &stats.HardUnsolved, &stats.TotalUnsolved)
	if err != nil {
		return nil, err
	}
	row = u.Db.QueryRow(query, username)

	easyTried := 1
	mediumTried := 1
	hardTried := 1
	totalTried := 1
	err = row.Scan(&stats.EasySolved, &stats.MediumSolved, &stats.HardSolved, &stats.TotalSolved,
		&easyTried, &mediumTried, &hardTried, &totalTried)

	if err != nil {
		return nil, err
	}

	acceptanceRate := (float64(stats.EasySolved) / float64(easyTried)) * 100
	stats.EasyAcceptanceRate = math.Floor(acceptanceRate*100) / 100

	acceptanceRate = (float64(stats.MediumSolved) / float64(mediumTried)) * 100
	stats.MediumAcceptanceRate = math.Floor(acceptanceRate*100) / 100

	acceptanceRate = (float64(stats.HardSolved) / float64(hardTried)) * 100
	stats.HardAcceptanceRate = math.Floor(acceptanceRate*100) / 100

	acceptanceRate = (float64(stats.TotalSolved) / float64(totalTried)) * 100
	stats.TotalAcceptanceRate = math.Floor(acceptanceRate*100) / 100

	return &stats, nil
}

func (u *UserRepo) GetSkillsByUsername(username string) (*[]model.Skill, error) {
	query := `
	with topics as(
		select
			topic_name
		from
			topics_problems
		where 
			problem_title in (
			select distinct
				p.title
			from
				submissions as s
			join
				problems as p
			on
				s.submission_status = 'Accepted' and
				s.user_username = $1 and
				s.problem_title = p.title 
			)
	)
	select
		topic_name,
		count(*)
	from
		topics
	group by
		topic_name
	`

	rows, err := u.Db.Query(query, username)
	if err != nil {
		return nil, err
	}

	skills := []model.Skill{}
	for rows.Next() {
		skill := model.Skill{}
		err = rows.Scan(&skill.SkillName, &skill.NumberOfTimesUsed)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return &skills, rows.Err()
}

func (u *UserRepo) GetLanguagesWithNumberOfAcceptedProblemsByUsername(username string) (*[]model.UsedLanguage, error) {
	query := `
	select distinct
			s.language_name,
			count(*)
		from
			submissions as s
		join
			problems as p
		on
			s.submission_status = 'Accepted' and
			s.user_username = $1 and
			s.problem_title = p.title 
		group by
			s.language_name
	
	`
	rows, err := u.Db.Query(query, username)
	if err != nil {
		return nil, err
	}

	languages := []model.UsedLanguage{}
	for rows.Next(){
		language := model.UsedLanguage{}
		err = rows.Scan(&language.Name, &language.NumberOfTimesUsed)
		if err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}

	return &languages, rows.Err()
}

// Update
func (u *UserRepo) UpdateUser(user *model.User) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 
		users 
	set  
		full_name =     $1, 
		email =         $2, 
		password =      $3, 
		profile_image = $4, 
		gender =        $5, 
		location =      $6, 
		birthday =      $7, 
		summary =       $8, 
		website =       $9, 
		github =        $10, 
		linkedin =      $11,
		updated_at =    $12
	where 
		deleted_at is null and username = $13 `

	result, err := tx.Exec(query, user.FullName, user.Email, user.Password,
		user.ProfileImage, user.Gender, user.Location, user.Birthday, user.Summary,
		user.Website, user.Github, user.LinkedIn, time.Now(), user.Username)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing updated")
	}

	return err
}

// Delete
func (u *UserRepo) DeleteUser(username string) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update 
		users 
	set 
		deleted_at = $1
	where 
		deleted_at is null and username = $2 `
	result, err := tx.Exec(query, time.Now(), username)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing deleted")
	}

	return err
}
