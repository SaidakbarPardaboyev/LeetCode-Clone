create table topics (
    id uuid primary key unique default gen_random_uuid() not null,
	name varchar unique not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

create table topics_problems (
	id uuid primary key default gen_random_uuid() not null,
	problem_id varchar references problems(id) not null,
	topic_id varchar references topics(id) not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

-- get topics name by problemId
select
    tp.problem_id,
    array_agg(t.name) as topics
from
    topics as t
inner join 
    topics_problems as tp 
        on t.id = tp.topic_id
where
    tp.problem_id = '79cb0553-226c-4368-b3fb-dc2b5f3b74ab' and
    t.deleted_at is null
group by
    tp.problem_id;