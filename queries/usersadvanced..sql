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
	number_solved_problems_by_difficulty as(
		select 
			als.user_username,
			count(case when p.difficulty = 'Easy' then 1 end) as easy_count,
			count(case when p.difficulty = 'Medium' then 1 end) as medium_count,
			count(case when p.difficulty = 'Hard' then 1 end) as hard_count,
			count(*) as total_count
		from
			all_solved_problems as als
		join
			problems as p
		on
			p.title = als.problem_title
		group by
			als.user_username
	),
	quality as(
		select
			user_username,
			easy_count + medium_count * 3 + hard_count * 5 as quality_index
		from
			number_solved_problems_by_difficulty
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
			nsd.user_username,
			nsd.total_count,
			q.quality_index,
			sp.speed,
			a.acceptance_rate,
			dense_rank() over (order by q.quality_index desc, a.acceptance_rate desc, sp.speed asc, nsd.total_count desc) as rank
		from
			number_solved_problems_by_difficulty as nsd
		join
			quality as q
		on
			q.user_username = nsd.user_username
		join
			speed_of_solutions as sp
		on
			sp.user_username = nsd.user_username
		join
			acceptance as a
		on
			a.user_username = nsd.user_username
	)
	select
		rank
	from
		rankings
	where 
		user_username = $1;

insert into 
		submissions(problem_title, user_username, language_name, 
		code, submission_status, runtime, submission_date)
	values('regular-expression-matching', 'stark', 'C', 'E67x316eGXT5PNufKFBKWktgzKb0kpz9vmpUq1oWqIEHUKFpY7',
    'Accepted', 0.12, current_timestamp)


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
			s.user_username = 'rrogers' and
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
			s.user_username = 'rrogers' and
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