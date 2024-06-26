create table examples (
	id uuid primary key unique default gen_random_uuid() not null,
	problem_title varchar references problems(title) not null,
	input text not null,
	output text not null,
	explanation text,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);