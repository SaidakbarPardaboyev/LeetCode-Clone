create table testcases_btree (
	id              serial primary key unique,
	problem_title   varchar references problems(title),
	val             integer,
	left_child      integer references testcases_btree(id),
	right_child     integer references testcases_btree(id)
)