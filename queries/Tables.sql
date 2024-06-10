create type difficulty as enum(
  'Easy',
  'Medium',
  'Hard'
);

create type status as enum(
  'Passed',
  'Run time Error',
  'Compile Error',
  'Wrong Answer',
  'Time Limit Exceeded',
  'Memory Limit Exceeded',
  'Output Limit Exceeded'
);

create table users(
    id uuid primary key default gen_random_uuid(),
    full_name varchar not null,
    username varchar unique not null,
    bio varchar not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table problems(
    id uuid primary key default gen_random_uuid(),
    question_number int unique not null,
    title varchar not null,
    difficulty_level difficulty not null,
    description text not null,
    examples text[] not null,
    hints text[] not null,
    constraints VARCHAR[],
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table languages(
    id uuid primary key default gen_random_uuid(),
    name varchar unique not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table submissions(
    id uuid primary key default gen_random_uuid(),
    problem_id uuid references problems(id) not null,
    user_id uuid references users(id) not null,
    language_id uuid references languages(id) not null,
    code text not null,
    submission_status status not null,
    submission_date timestamp not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);


create table topics(
    id uuid primary key default gen_random_uuid(),
    name varchar unique not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table topics_problems(
    id uuid primary key default gen_random_uuid(),
    topics_id uuid references topics(id),
    problem_id uuid references problems(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);



