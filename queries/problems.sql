create type difficulty_lavels as enum('Easy', 'Medium', 'Hard');

create table problems (
	title varchar primary key unique not null,
	problem_number serial unique not null,
	difficulty difficulty_lavels not null,
	description text not null,
	constraints text[] not null,
	hints text[],
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

insert into problems(
	title, difficulty, description, constraints
) values (
	'Two Sum',
	'Easy', 
	'Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.',
	'{
    	2 <= nums.length <= 104,
    	-109 <= nums[i] <= 109,
    	-109 <= target <= 109,
    	Only one valid answer exists.
	}'
);