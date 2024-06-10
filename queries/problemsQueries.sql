create type problem_status as enum('Easy', 'Medium', 'Hard');

create table topics(
    id uuid primary key default gen_random_uuid() NOT NULL,
    name varchar unique NOT NULL,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

create table problems(
    id uuid primary key default gen_random_uuid() NOT NULL,
    problem_num serial unique NOT NULL,
    title varchar NOT NULL,
    status problem_status NOT NULL,
    description text NOT NULL,
    examples text[] NOT NULL,
    constraints varchar[] NOT NULL,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

create table problems_topics(
    id uuid primary key default gen_random_uuid() NOT NULL,
    problem_id uuid references problems(id) NOT NULL,
    topic_id uuid references topics(id) NOT NULL,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);  

drop type problem_status;
drop table topics;
drop table problems;
drop table problems_topics;