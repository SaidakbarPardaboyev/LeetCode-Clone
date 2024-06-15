create table languages (
	name  varchar primary key not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);
