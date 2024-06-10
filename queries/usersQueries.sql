create table users(
    id uuid primary key default gen_random_uuid() NOT NULL,
    fullname varchar NOT NULL,
    username varchar unique NOT NULL,
    bio text,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

drop table users;