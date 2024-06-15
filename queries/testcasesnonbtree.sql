create table testcases_non_btree (
	id                  uuid primary key unique default gen_random_uuid() not null,
	problem_title       varchar references problems(title) not null,
	param1				JSONB,
	param2				JSONB,
	param3				JSONB,
	param4				JSONB,
	param5				JSONB,
	param6				JSONB,
	answer 				JSONB
);

insert into testcases_non_btree(
	param1, param2, problem_title, answer
) values(
	'{"val": [2,7,11,15]}', '{"val":9}', 'Two Sum', '{"answer": [0,1]}'
);

-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{2,7,11,15}', 9, '5bbaa8cb-5aa0-493f-ad8a-b786575a7840' -- [0,1]
-- );
-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{3,2,4}', 6, '076a2953-349e-47ea-bb54-a239792a9f5d' -- [1.2]
-- );
-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{3,3}', 6, '5bbaa8cb-5aa0-493f-ad8a-b786575a7840' -- [0,1]
-- );
-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{3,2,3}', 6, 'cd500ad9-56a0-4c1e-b023-cd7f65684f63' --[0,2]
-- );
-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{2,5,5,11}', 10, '076a2953-349e-47ea-bb54-a239792a9f5d' --[1,2]
-- );
-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{0,4,3,0}', 0, 'a3907ab1-1041-4457-9c34-c51db9c7b9dd' -- [0,3]
-- );
-- INSERT INTO testcases_non_btree (
-- 	problem_title, int_array, int_val, answer_id
-- ) VALUES (
-- 	'Two Sum', '{-3,4,3,90}', 0, 'cd500ad9-56a0-4c1e-b023-cd7f65684f63' -- [0,2]
-- );

-- with problem_title_of_question as (
-- 	select
-- 		problem_title
-- 	from
-- 		testcases_non_btree
-- 	WHERE 
-- 		problem_title = 'Two Sum'
-- 	limit
-- 		1
-- )

-- select
-- 	answer_id, a.problem_title, int_2d_array_answer, float_2d_array_answer, string_2d_array_answer, byte_2d_array_answer, bool_2d_array_answer, int_array_answer, float_array_answer, string_array_answer, byte_array_answer, bool_array_answer, int_val_answer, float_val_answer, string_val_answer, byte_val_answer, bool_val_answer
-- from
-- 	answers as a
-- inner join
-- 	problem_title_of_question as p 
-- 		on p.problem_title = a.problem_title;



-- select
-- 	*
-- from 
-- 	testcases_non_btree
-- WHERE 
-- 	problem_title = 'Two Sum';


-- create table answers (
-- 	answer_id               uuid primary key unique default gen_random_uuid() not null,
-- 	problem_title       	varchar references problems(title) not null,
-- 	int_2d_array_answer		bigint[][],
-- 	float_2d_array_answer 	double precision[][],
-- 	string_2d_array_answer  varchar[][],
-- 	byte_2d_array_answer    bytea[][],
-- 	bool_2d_array_answer    boolean[][],
-- 	int_array_answer        bigint[],
-- 	float_array_answer      double precision[],
-- 	string_array_answer     varchar[],
-- 	byte_array_answer       bytea[],
-- 	bool_array_answer       boolean[],
-- 	int_val_answer          bigint,
-- 	float_val_answer        double precision,
-- 	string_val_answer       varchar,
-- 	byte_val_answer         bytea,
-- 	bool_val_answer         boolean
-- );

-- create table testcases_non_btree (
-- 	id                  uuid primary key unique default gen_random_uuid() not null,
-- 	problem_title       varchar references problems(title) not null,
-- 	answer_id			uuid references answers(answer_id) not null,
-- 	int_2d_array        bigint[][],
-- 	int_2d_array2       bigint[][],
-- 	float_2d_array      double precision[][],
-- 	float_2d_array2     double precision[][],
-- 	string_2d_array     varchar[][],
-- 	string_2d_array2    varchar[][],
-- 	byte_2d_array       bytea[][],
-- 	byte_2d_array2      bytea[][],
-- 	bool_2d_array       boolean[][],
-- 	bool_2d_array2      boolean[][],
-- 	int_array           bigint[],
-- 	int_array2          bigint[],
-- 	float_array         double precision[],
-- 	float_array2        double precision[],
-- 	string_array        varchar[],
-- 	string_array2       varchar[],
-- 	byte_array          bytea[],
-- 	byte_array2         bytea[],
-- 	bool_array          boolean[],
-- 	bool_array2         boolean[],
-- 	int_val             bigint,
-- 	int_val2            bigint,
-- 	int_val3            bigint,
-- 	float_val           double precision,
-- 	float_val2          double precision,
-- 	float_val3          double precision,
-- 	string_val          varchar,
-- 	string_val2         varchar,
-- 	string_val3         varchar,
-- 	byte_val            bytea,
-- 	byte_val2           bytea,
-- 	byte_val3           bytea,
-- 	bool_val            boolean,
-- 	bool_val2           boolean,
-- 	bool_val3           boolean,
--     created_at 			timestamp default now() not null,
--     updated_at 			timestamp,
--     deleted_at 			timestamp
-- );