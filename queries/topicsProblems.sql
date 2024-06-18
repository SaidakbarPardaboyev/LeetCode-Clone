create table topics (
	name varchar primary key unique not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

create table topics_problems (
	id uuid primary key default gen_random_uuid() not null,
	problem_title varchar references problems(title) not null,
	topic_name varchar references topics(name) not null,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

INSERT INTO topics (name, created_at, updated_at, deleted_at)
VALUES
('array', DEFAULT, NULL, NULL),
('string', DEFAULT, NULL, NULL),
('linked-list', DEFAULT, NULL, NULL),
('stack', DEFAULT, NULL, NULL),
('queue', DEFAULT, NULL, NULL),
('tree', DEFAULT, NULL, NULL),
('binary-search', DEFAULT, NULL, NULL),
('heap', DEFAULT, NULL, NULL),
('hash-table', DEFAULT, NULL, NULL),
('two-pointers', DEFAULT, NULL, NULL),
('depth-first-search', DEFAULT, NULL, NULL),
('breadth-first_search', DEFAULT, NULL, NULL),
('backtracking', DEFAULT, NULL, NULL),
('dynamic-programming', DEFAULT, NULL, NULL),
('greedy', DEFAULT, NULL, NULL),
('design', DEFAULT, NULL, NULL),
('math', DEFAULT, NULL, NULL),
('bit_manipulation', DEFAULT, NULL, NULL),
('sort', DEFAULT, NULL, NULL),
('recursion', DEFAULT, NULL, NULL),
('memoization', DEFAULT, NULL, NULL),
('segment_tree', DEFAULT, NULL, NULL),
('union_find', DEFAULT, NULL, NULL),
('trie', DEFAULT, NULL, NULL),
('sliding_window', DEFAULT, NULL, NULL),
('binary-indexed-tree', DEFAULT, NULL, NULL),
('topological-sort', DEFAULT, NULL, NULL),
('minimum-spanning-tree', DEFAULT, NULL, NULL),
('suffix-array', DEFAULT, NULL, NULL),
('geometry', DEFAULT, NULL, NULL),
('simulation', DEFAULT, NULL, NULL),
('probability', DEFAULT, NULL, NULL),
('concurrency', DEFAULT, NULL, NULL),
('database', DEFAULT, NULL, NULL);
