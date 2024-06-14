create table topics {
	id uuid primary key not null,
	name varchar unique not null
}

create table topics_problems {
	id uuid primary key not null,
	problem_title varchar references problems(title) not null,
	topic_id uuid references topics(id) not null
}