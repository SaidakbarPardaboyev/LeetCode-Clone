create table languages(
    id uuid primary key default gen_random_uuid() NOT NULL,
    name varchar NOT NULL,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

create type submission_statuses as enum('Eccepted', 'Wrong Answer', 'Time Limit Exceeded', 'Memory Limit Exceeded', 'Runtime Error', 'Compile Error');

create table submissions(
    id uuid primary key default gen_random_uuid() NOT NULL,
    user_id uuid references users(id) NOT NULL,
    problem_id uuid references problems(id) NOT NULL,
    language uuid references languages(id) NOT NULL,
    submission_status submission_statuses NOT NULL,
    code text NOT NULL,
    submission_date timestamp default current_timestamp NOT NULL,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

drop table languages;
drop table submissions;
drop type submission_status;