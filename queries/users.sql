create table users {
	username varchar primary key not null,
	full_name varchar not null,
    email varchar not null,
	password varchar not null,
	profile_image bytea[],
	gender varchar,
	location varchar,
	birthday date,
	summary text,
	website varchar,
	github varchar,
	LinkedIn varchar
}