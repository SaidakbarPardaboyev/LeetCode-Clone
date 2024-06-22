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
('bit-manipulation', DEFAULT, NULL, NULL),
('sort', DEFAULT, NULL, NULL),
('recursion', DEFAULT, NULL, NULL),
('memoization', DEFAULT, NULL, NULL),
('segment_tree', DEFAULT, NULL, NULL),
('union-find', DEFAULT, NULL, NULL),
('trie', DEFAULT, NULL, NULL),
('sliding-window', DEFAULT, NULL, NULL),
('binary-indexed-tree', DEFAULT, NULL, NULL),
('topological-sort', DEFAULT, NULL, NULL),
('minimum-spanning-tree', DEFAULT, NULL, NULL),
('suffix-array', DEFAULT, NULL, NULL),
('geometry', DEFAULT, NULL, NULL),
('simulation', DEFAULT, NULL, NULL),
('probability', DEFAULT, NULL, NULL),
('concurrency', DEFAULT, NULL, NULL),
('database', DEFAULT, NULL, NULL);

-- Two Sum
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('two-sum', 'array'),
('two-sum', 'hash-table');

-- Add Two Numbers
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('add-two-numbers', 'linked-list'),
('add-two-numbers', 'math');

-- Longest Substring Without Repeating Characters
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('longest-substring-without-repeating-characters', 'string'),
('longest-substring-without-repeating-characters', 'two-pointers'),
('longest-substring-without-repeating-characters', 'sliding-window');

-- Median of Two Sorted Arrays
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('median-of-two-sorted-arrays', 'array'),
('median-of-two-sorted-arrays', 'binary-search');

-- Longest Palindromic Substring
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('longest-palindromic-substring', 'string'),
('longest-palindromic-substring', 'dynamic-programming');

-- ZigZag Conversion
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('zigzag-conversion', 'string');

-- Reverse Integer
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('reverse-integer', 'math');

-- String to Integer (atoi)
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('string-to-integer-(atoi)', 'string');

-- Palindrome Number
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('palindrome-number', 'math');

-- Regular Expression Matching
INSERT INTO topics_problems (problem_title, topic_name)
VALUES
('regular-expression-matching', 'string'),
('regular-expression-matching', 'dynamic-programming');