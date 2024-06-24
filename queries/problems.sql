create type difficulty_lavels as enum('Easy', 'Medium', 'Hard');

create table problems (
	id uuid primary key unique default gen_random_uuid() not null,
	title varchar unique not null,
	problem_number serial unique not null,
	difficulty difficulty_lavels not null,
	description text not null,
	constraints text[] not null,
	hints text[],
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

-- GetSubmissionStatisticsByProblemTitle
with submission_stat as (
	select
		count(
			case
				when submission_status = 'Accepted' then 1
			end
		) as accepted,
		count(*) as submited
	from
		submissions 
	where
		problem_title = 'longest-palindromic-substring'
) 

select
	accepted,
	submited,
	round(accepted::numeric / submited * 100, 1) as acceptence_rate
from 
	submission_stat;

-- select by searching in problem_number and title
SELECT problem_number, title
FROM problems
WHERE search_vector @@ plainto_tsquery('english', 'longest substring')
ORDER BY ts_rank(search_vector, plainto_tsquery('english', 'longest substring')) DESC;

SELECT problem_number, title
FROM problems
WHERE search_vector @@ plainto_tsquery('english', '1 sum')
ORDER BY ts_rank(search_vector, plainto_tsquery('english', '1 sum')) DESC;

-- select problems by status(easy, medium, hard)
select
	problem_number,
	title,
	difficulty
from
	problems 
where
	difficulty = 'Medium';

-- select topics
select
	p.problem_number,
	title
from 
	problems as p
inner join
	topics_problems as t
		on t.problem_title = p.title
where
	t.topic_name in ('string')
group by
	p.problem_number, title
having
	count(distinct t.topic_name) = 1
order by
	p.problem_number;

-- select by acceptence rate
select
	p.problem_number,
	title
	round(count(case when submission_status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2) as acceptenceRate
from
	problems as p
inner join
	submissions as s
		on p.title = s.problem_title
group by
	p.problem_number, p.title
order by
	round(count(case when submission_status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2) asc;

-- select by difficulty (from begin to end/from end to begin)
select
	title,
	difficulty
from
	problems
order by
	difficulty desc;

-- select by problem_number (asc/desc)
select
	problem_number,
	title
from
	problems
order by
	problem_number desc;

-- select by status (NOT_STARTED)
with AcceptedProblemsTitle as (
	select 
		distinct problem_title 
	from 
		submissions 
	where
		user_username='jdoe'
)

select
	p.problem_number,
	p.title
from
	problems as p
where
	p.title not in (
		select 
			distinct problem_title 
		from 
			AcceptedProblemsTitle
	);

-- select by status (AC)
with AcceptedProblemsTitle as (
	select 
		distinct problem_title 
	from 
		submissions 
	where
		user_username='jdoe' and 
		submission_status='Accepted'
)

select
	p.problem_number,
	p.title
from
	problems as p
where
	p.title in (
		select 
			distinct problem_title 
		from 
			AcceptedProblemsTitle
	);

-- select by status (TRIED but not AC)
with AcceptedProblemsTitle as (
	select 
		problem_title,
		array_agg(submission_status)
	from 
		submissions 
	where
		user_username='jdoe' 
	group by
		problem_title
	having
		not ('Accepted' = ANY(array_agg(submission_status)))
)

select
	p.problem_number,
	p.title
from
	problems as p
where
	p.title in (
		select 
			problem_title 
		from 
			AcceptedProblemsTitle
	);

-- select by topics and acceptenceRate of sorting
select
	p.problem_number,
	title
from 
	problems as p
inner join
	topics_problems as t
		on t.problem_title = p.title
inner join
	submissions as s
		on p.title = s.problem_title
where
	t.topic_name in ('string', 'dynamic-programming') and 
	search_vector @@ plainto_tsquery('english', 'substring') and
	p.difficulty = 'Medium'
group by
	p.problem_number, title
having
	count(distinct t.topic_name) = 2
order by
	round(count(case when submission_status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2) desc,
	ts_rank(search_vector, plainto_tsquery('english', '1 sum')) DESC;

-- beginning of get all by fulter 
with acceptenceRatesOfProblems as (
	select
		title as problem_title,
		round(count(case when submission_status = 'Accepted' then 1 end)::numeric / count(*) * 100, 2)as acceptence
	from
		problems as p
	inner join
		submissions as s
			on p.title = s.problem_title
	group by
		p.problem_number, title
	order by
		p.problem_number
)

select
	case
		when 'Accepted' = ANY(array_agg(sub.submission_status)) then 'Accepted'
		when count(sub.submission_status) > 0 then 'Attempted'
		else 'NOT-TRIED'
	end as status,
	p.problem_number as problem_number,
	p.title as problem_title,
	a.acceptence as acceptence,
	p.difficulty as difficulty
from
	problems as p
left join
	submissions as sub
		on sub.problem_title = p.title and
		'jdoe' = sub.user_username
left join
	acceptenceRatesOfProblems as a
		on a.problem_title = p.title
group by
	p.problem_number, p.title, a.acceptence
order by
	p.problem_number;