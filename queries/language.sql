create table languages (
    id 		   uuid primary key gen_random_uuid()
	name       varchar not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);
INSERT INTO languages (name, created_at, updated_at, deleted_at)
VALUES
('Go', DEFAULT, NULL, NULL),
('Python3', DEFAULT, NULL, NULL),
('C', DEFAULT, NULL, NULL),
('C++', DEFAULT, NULL, NULL),
('Rust', DEFAULT, NULL, NULL),
('Java', DEFAULT, NULL, NULL),
('JavaScript', DEFAULT, NULL, NULL),
('Kotlin', DEFAULT, NULL, NULL),
('PHP', DEFAULT, NULL, NULL);
