create type enum as difficulty_lavels("Easy", "Medium", "Hard");

create table problems {
	id uuid primary key unique,
	title varchar unique,
	problem_number integer increments unique,
	difficulty enum references difficulty_lavels,
	description text,
	constraints text[],
	hints text[]
}

create table topics {
	id uuid primary key unique
	name varchar
}
