package postgres

import (
	"database/sql"
	"fmt"
	"leetcode/models"
	"math"
	"time"

	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Create
func (u *UserRepo) CreateUser(user *models.CreateUser) (string, error) {

	tx, err := u.Db.Begin()
	if err != nil {
		return "", err
	}

	query := `
	insert into 
		users(id, username, email, password) 
		values($1, $2, $3, $4)`
	_, err = tx.Exec(query, uuid.NewString(), user.Username, user.Email, user.Password)

	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()

	return "", err
}

// Read
func (u *UserRepo) GetUserByUsername(id string) (models.User, error) {
	user := models.User{Id: id}
	query := `
	select 
		username, full_name, email, password, profile_image, 
		gender, location, birthday, summary, website, github, linkedin,
		created_at, updated_at
	from 
		users
	where
		username = $1 and deleted_at is null
	`
	row := u.Db.QueryRow(query, id)
	err := row.Scan(&user.Username, &user.FullName, &user.Email,
		&user.Password, &user.ProfileImage, &user.Gender, &user.Location, &user.Birthday,
		&user.Summary, &user.Website, &user.Github, &user.LinkedIn,
		&user.CreatedAt, &user.UpdatedAt)

	return user, err
}
func (u *UserRepo) GetUsers(filter *models.UserFilter) (*[]models.User, error) {
	params := []interface{}{}
	paramCount := 1
	query := `
	select 
		username, full_name, email, password, profile_image, 
		gender, location, birthday, summary, website, github, linkedin,
		created_at, updated_at 
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

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Username, &user.FullName, &user.Email,
			&user.Password, &user.ProfileImage, &user.Gender, &user.Location, &user.Birthday,
			&user.Summary, &user.Website, &user.Github, &user.LinkedIn,
			&user.CreatedAt, &user.UpdatedAt)
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

func (u *UserRepo) GetUserRankingByUserId(userId string) (int, error) {
	//number of Solved Problems
	//difficulty of problem
	//average speed of solutions
	//acceptance rate
	query := `
	with all_solved_problems as(
		select distinct
			problem_id,
			user_id
		from
			submissions
		where
			submission_status = 'Accepted'
		group by
			problem_id, user_id
	),
	quality as(
		select
			als.user_id,
			count(case when p.difficulty = 'Easy' then 1 end) + 
            count(case when p.difficulty = 'Medium' then 1 end) * 3 + 
            count(case when p.difficulty = 'Hard' then 1 end) * 5 as quality_index,
            count(*) total_count
		from
			all_solved_problems as als
		join
			problems as p
		on
			p.title = als.problem_id
		group by
			als.user_id
	),
	speed_of_solutions as(
		select 
			user_id,
			round(avg(runtime)::numeric, 2) as speed
		from
			submissions
		where
			submission_status = 'Accepted'
		group by
			user_id
	),
	acceptance as (
		select 
			user_id,
			round(
				count(case when submission_status = 'Accepted' then 1 end)::numeric * 100 / count(*), 2
				) as acceptance_rate
		from
			submissions
		group by
			user_id
	),
	rankings as(
		select
			q.user_id,
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
			sp.user_id = q.user_id
		join
			acceptance as a
		on
			a.user_id = q.user_id
	)
	select
		rank
	from
		rankings
	where 
		user_id = $1;
	`
	rank := 0
	row := u.Db.QueryRow(query, userId)
	err := row.Scan(&rank)

	return rank, err
}

func (u *UserRepo) GetNumberOfSolvedProblemsByUserId(userId string) (*models.AllStatisticsOfSolvedProblems, error) {
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
			s.user_id,
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
			s.user_id = $1 and
			s.problem_id = p.id 
		group by
			s.user_id

	),
	tried_problems as(
		select
			s.user_id,
			count(case when p.difficulty = 'Easy'  then 1 end) as easy_tried,
			count(case when p.difficulty = 'Medium'  then 1 end) as medium_tried,
			count(case when p.difficulty = 'Hard'  then 1 end) as hard_tried,
			count(*) total_tried
		from
			submissions as s	
		join
			problems as p
		on	
			s.user_id = $1 and
			s.problem_id = p.id 
		group by
			s.user_id

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
		sp.user_id = tp.user_id
	`

	stats := models.AllStatisticsOfSolvedProblems{}

	row := u.Db.QueryRow(unsolvedProblemsQuery)

	err := row.Scan(&stats.EasyUnsolved, &stats.MediumUnsolved, &stats.HardUnsolved, &stats.TotalUnsolved)
	if err != nil {
		return nil, err
	}
	row = u.Db.QueryRow(query, userId)

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

func (u *UserRepo) GetSkillsByUserId(userId string) (*[]models.Skill, error) {
	query := `
	with distinct_topics as(
		select
			topic_id
		from
			topics_problems
		where 
			problem_id in (
				select distinct
					p.id
				from
					submissions as s
				join
					problems as p
				on
					s.submission_status = 'Accepted' and
					s.user_id = $1 and
					s.problem_id = p.id 
		)
	)
	select
		t.name,
		count(*)
	from
		distinct_topics as dt
	join
		topics as t
	on
		t.id = dt.topic_id
	group by
		t.name
	`

	rows, err := u.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	skills := []models.Skill{}
	for rows.Next() {
		skill := models.Skill{}
		err = rows.Scan(&skill.SkillName, &skill.NumberOfTimesUsed)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return &skills, rows.Err()
}

func (u *UserRepo) GetLanguagesWithNumberOfAcceptedProblemsByUserId(userId string) (*[]models.UsedLanguage, error) {
	query := `
	with distict_langs as(
		select distinct
				s.language_id,
				count(*)
			from
				submissions as s
			join
				problems as p
			on
				s.submission_status = 'Accepted' and
				s.user_id = $1 and
				s.problem_id = p.id
			group by
				s.language_id
	)
	select 
		l.name,
		dl.count
	from
		languages as l
	join
		distict_langs as dl
	on
		l.id == dl.language_id	
	`
	rows, err := u.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	languages := []models.UsedLanguage{}
	for rows.Next() {
		language := models.UsedLanguage{}
		err = rows.Scan(&language.Name, &language.NumberOfTimesUsed)
		if err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}

	return &languages, rows.Err()
}

// Update
func (u *UserRepo) UpdateUser(user *models.UpdateUser) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	query := `
	update 
		users 
	set  
		username =      $1
		full_name =     $2, 
		email =         $3, 
		password =      $4, 
		profile_image = $5, 
		gender =        $6, 
		location =      $7, 
		birthday =      $8, 
		summary =       $9, 
		website =       $10, 
		github =        $11, 
		linkedin =      $12,
		updated_at =    $13
	where 
		deleted_at is null and id = $14 `

	result, err := tx.Exec(query, user.Username, user.FullName, user.Email, user.Password,
		user.ProfileImage, user.Gender, user.Location, user.Birthday, user.Summary,
		user.Website, user.Github, user.LinkedIn, time.Now(), user.Id)

	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing updated")
	}
	err = tx.Commit()

	return err
}

// Delete
func (u *UserRepo) DeleteUser(userId string) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	query := `
	update 
		users 
	set 
		deleted_at = $1
	where 
		deleted_at is null and id = $2 `
	result, err := tx.Exec(query, time.Now(), userId)

	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("nothing deleted")
	}
	err = tx.Commit()

	return err
}
