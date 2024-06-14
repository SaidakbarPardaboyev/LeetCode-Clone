create table testcases_btree (
	id              serial primary key unique not null,
	problem_title   varchar references problems(title) not null,
	val             integer,
	left_child      integer references testcases_btree(id),
	right_child     integer references testcases_btree(id),
    created_at 		timestamp default now() not null,
    updated_at 		timestamp,
    deleted_at 		timestamp
);