create type enum as gender("Male", "Female");

create table users (
	username 		varchar primary key not null,
	full_name 		varchar not null,
    email 			varchar not null,
	password 		varchar not null,
	profile_image 	bytea[],
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