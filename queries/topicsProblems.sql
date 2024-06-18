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
('Array', DEFAULT, NULL, NULL),
('String', DEFAULT, NULL, NULL),
('Linked List', DEFAULT, NULL, NULL),
('Stack', DEFAULT, NULL, NULL),
('Queue', DEFAULT, NULL, NULL),
('Tree', DEFAULT, NULL, NULL),
('Binary Search', DEFAULT, NULL, NULL),
('Heap', DEFAULT, NULL, NULL),
('Hash Table', DEFAULT, NULL, NULL),
('Two Pointers', DEFAULT, NULL, NULL),
('Depth-First Search (DFS)', DEFAULT, NULL, NULL),
('Breadth-First Search (BFS)', DEFAULT, NULL, NULL),
('Backtracking', DEFAULT, NULL, NULL),
('Dynamic Programming', DEFAULT, NULL, NULL),
('Greedy', DEFAULT, NULL, NULL),
('Design', DEFAULT, NULL, NULL),
('Math', DEFAULT, NULL, NULL),
('Bit Manipulation', DEFAULT, NULL, NULL),
('Sort', DEFAULT, NULL, NULL),
('Recursion', DEFAULT, NULL, NULL),
('Memoization', DEFAULT, NULL, NULL),
('Segment Tree', DEFAULT, NULL, NULL),
('Union Find', DEFAULT, NULL, NULL),
('Trie', DEFAULT, NULL, NULL),
('Sliding Window', DEFAULT, NULL, NULL),
('Binary Indexed Tree (BIT)', DEFAULT, NULL, NULL),
('Topological Sort', DEFAULT, NULL, NULL),
('Minimum Spanning Tree (MST)', DEFAULT, NULL, NULL),
('Suffix Array', DEFAULT, NULL, NULL),
('Geometry', DEFAULT, NULL, NULL),
('Simulation', DEFAULT, NULL, NULL),
('Probability', DEFAULT, NULL, NULL),
('Concurrency', DEFAULT, NULL, NULL),
('Database', DEFAULT, NULL, NULL);
