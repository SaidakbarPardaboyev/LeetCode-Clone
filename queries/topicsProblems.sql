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
