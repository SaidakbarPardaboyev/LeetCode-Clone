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

-------------------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------------------

-- add full text searching in problems table 
alter 
	table 
		problems 
	add 
		column search_vector tsvector;

UPDATE 
	problems 
SET 
	search_vector = to_tsvector('english', title || ' ' || problem_number::text);

create index 
	problems_search_vector_idx 
		ON problems USING gin(search_vector);

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

--------------------------------------------------------------------------------------------
--------------------------------------------------------------------------------------------

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('two-sum', DEFAULT, 'Easy', 'Find indices of the two numbers such that they add up to a specific target.', ARRAY['Each input would have exactly one solution.', 'You may not use the same element twice.'], ARRAY['Try using a hashmap.', 'Think about the complement of each number.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('add-two-numbers', DEFAULT, 'Medium', 'Add two numbers represented by linked lists.', ARRAY['The digits are stored in reverse order.', 'Each node contains a single digit.'], ARRAY['Use a dummy head to simplify code.', 'Keep track of carry.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('longest-substring-without-repeating-characters', DEFAULT, 'Medium', 'Find the length of the longest substring without repeating characters.', ARRAY['0 <= s.length <= 5 * 10^4', 's consists of English letters, digits, symbols and spaces.'], ARRAY['Use a sliding window technique.', 'Keep track of characters using a set.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('median-of-two-sorted-arrays', DEFAULT, 'Hard', 'Find the median of the two sorted arrays.', ARRAY['You may assume nums1 and nums2 cannot be both empty.'], ARRAY['Try to divide the problem into smaller parts.', 'Consider the use of binary search.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('longest-palindromic-substring', DEFAULT, 'Medium', 'Find the longest palindromic substring in a given string.', ARRAY['1 <= s.length <= 1000', 's consists of only digits and English letters.'], ARRAY['Expand around the center.', 'Dynamic programming can also be used.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('zigzag-conversion', DEFAULT, 'Medium', 'Convert a string to zigzag pattern on a given number of rows.', ARRAY['1 <= s.length <= 1000', '1 <= numRows <= 1000'], ARRAY['Simulate the zigzag pattern row by row.', 'Use a list to build the output string.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('reverse-integer', DEFAULT, 'Medium', 'Reverse digits of an integer.', ARRAY['Assume the environment does not allow storing 64-bit integers (signed or unsigned).'], ARRAY['Handle overflow/underflow cases.', 'Consider the edge cases like zero.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('string-to-integer-(atoi)', DEFAULT, 'Medium', 'Convert a string to an integer.', ARRAY['Implement atoi which converts a string to an integer.', 'The function first discards as many whitespace characters as necessary.'], ARRAY['Check for valid characters.', 'Consider edge cases for overflow.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('palindrome-number', DEFAULT, 'Easy', 'Determine whether an integer is a palindrome.', ARRAY['Do this without converting the integer to a string.'], ARRAY['Consider reversing the integer.', 'Check edge cases such as negative numbers.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('regular-expression-matching', DEFAULT, 'Hard', 'Implement regular expression matching with support for . and *.', ARRAY['The matching should cover the entire input string (not partial).'], ARRAY['Use dynamic programming.', 'Consider each character in the string and the pattern.'], DEFAULT, NULL);
