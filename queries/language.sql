create table languages (
    id 		   uuid primary key gen_random_uuid()
	name       varchar not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);