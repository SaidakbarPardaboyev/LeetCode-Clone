create type gender as enum('Male', 'Female');

create table users (
	id 				uuid primary key gen_random_uuid()
	username 		varchar not null,
    email 			varchar not null,
	password 		varchar not null,
	full_name 		varchar,
	profile_image 	bytea,
	gender 			gender,
	location 		varchar,
	birthday 		date,
	summary 		text,
	website 		varchar,
	github 			varchar,
	linkedin 		varchar,
    created_at 		timestamp default now() not null,
    updated_at 		timestamp,
    deleted_at 		timestamp
);