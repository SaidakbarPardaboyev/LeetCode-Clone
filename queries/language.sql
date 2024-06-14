create table language (
	id    uuid primary key not null,
	name  varchar not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);
