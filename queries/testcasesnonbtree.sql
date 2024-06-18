create type run_submit as enum ('Run', 'Submit');

create table testcases_non_btree (
    id uuid primary key unique default gen_random_uuid() not null,
    problem_title varchar references problems(title) not null,
    function_name varchar NOT NULL,
    run_or_submit run_submit not null,
    arg1 JSONB,
    arg2 JSONB,
    arg3 JSONB,
    arg4 JSONB,
    arg5 JSONB,
    arg6 JSONB,
    answer JSONB,
    arg1_type varchar,
    arg2_type varchar,
    arg3_type varchar,
    arg4_type varchar,
    arg5_type varchar,
    arg6_type varchar,
    answer_type varchar,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Add', '123', '456', '579', 'int', 'int', 'int'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Concat', '"Hello "', '"World"', '"Hello World"', 'string', 'string', 'string'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Two Sum', '[2,7,11,15]', '9', '[0,1]', '[]int', 'int', '[]int'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Two Sum', '[3,3]', '6', '[0,1]', '[]int', 'int', '[]int'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Two Sum', '[3,2,3]', '6', '[0,2]', '[]int', 'int', '[]int'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Two Sum', '[2,5,5,11]', '10', '[1,2]', '[]int', 'int', '[]int'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Two Sum', '[0,4,3,0]', '0', '[0,3]', '[]int', 'int', '[]int'
);
INSERT INTO function_calls (
    function_name, arg1, arg2, answer, arg1_type, arg2_type, answer_type
) VALUES (
    'Two Sum', '[-3,4,3,90]', '0', '[0,2]', '[]int', 'int', '[]int'
);