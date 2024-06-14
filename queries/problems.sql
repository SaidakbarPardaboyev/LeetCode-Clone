create type enum as difficulty_lavels("Easy", "Medium", "Hard");

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