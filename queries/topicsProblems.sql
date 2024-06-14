create table topics (
	name varchar primary key unique not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

create table topics_problems (
	id uuid primary key default gen_random_uuid() not null,
	problem_title varchar references problems(title) not null,
	topic_name varchar references topics(name) not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);