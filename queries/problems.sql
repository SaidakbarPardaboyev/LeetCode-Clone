create type difficulty_lavels as enum('Easy', 'Medium', 'Hard');

create table problems (
	title varchar primary key unique not null,
	problem_number serial unique not null,
	difficulty difficulty_lavels not null,
	description text not null,
	constraints text[] not null,
	hints text[],
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

insert into problems(
	title, difficulty, description, constraints
) values (
	'two-sum',
	'Easy', 
	'Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.\n\nYou may assume that each input would have exactly one solution, and you may not use the same element twice.\n\nYou can return the answer in any order.',
	'{
    	2 <= nums.length <= 104,
    	-109 <= nums[i] <= 109,
    	-109 <= target <= 109,
    	Only one valid answer exists.
	}'
);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('two-sum', DEFAULT, 'Easy', 'Find indices of the two numbers such that they add up to a specific target.', ARRAY['Each input would have exactly one solution.', 'You may not use the same element twice.'], ARRAY['Try using a hashmap.', 'Think about the complement of each number.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('add-two-numbers', DEFAULT, 'Medium', 'Add two numbers represented by linked lists.', ARRAY['The digits are stored in reverse order.', 'Each node contains a single digit.'], ARRAY['Use a dummy head to simplify code.', 'Keep track of carry.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('longest-substring-without-repeating-characters', DEFAULT, 'Medium', 'Find the length of the longest substring without repeating characters.', ARRAY['0 <= s.length <= 5 * 10^4', 's consists of English letters, digits, symbols and spaces.'], ARRAY['Use a sliding window technique.', 'Keep track of characters using a set.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('median-of-two-sorted-arrays', DEFAULT, 'Hard', 'Find the median of the two sorted arrays.', ARRAY['You may assume nums1 and nums2 cannot be both empty.'], ARRAY['Try to divide the problem into smaller parts.', 'Consider the use of binary search.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('longest-palindromic-substring', DEFAULT, 'Medium', 'Find the longest palindromic substring in a given string.', ARRAY['1 <= s.length <= 1000', 's consists of only digits and English letters.'], ARRAY['Expand around the center.', 'Dynamic programming can also be used.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('zigzag-conversion', DEFAULT, 'Medium', 'Convert a string to zigzag pattern on a given number of rows.', ARRAY['1 <= s.length <= 1000', '1 <= numRows <= 1000'], ARRAY['Simulate the zigzag pattern row by row.', 'Use a list to build the output string.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('reverse-integer', DEFAULT, 'Medium', 'Reverse digits of an integer.', ARRAY['Assume the environment does not allow storing 64-bit integers (signed or unsigned).'], ARRAY['Handle overflow/underflow cases.', 'Consider the edge cases like zero.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('string-to-integer (atoi)', DEFAULT, 'Medium', 'Convert a string to an integer.', ARRAY['Implement atoi which converts a string to an integer.', 'The function first discards as many whitespace characters as necessary.'], ARRAY['Check for valid characters.', 'Consider edge cases for overflow.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('palindrome-number', DEFAULT, 'Easy', 'Determine whether an integer is a palindrome.', ARRAY['Do this without converting the integer to a string.'], ARRAY['Consider reversing the integer.', 'Check edge cases such as negative numbers.'], DEFAULT, NULL);

INSERT INTO problems (title, problem_number, difficulty, description, constraints, hints, updated_at, deleted_at)
VALUES
('regular-expression-matching', DEFAULT, 'Hard', 'Implement regular expression matching with support for . and *.', ARRAY['The matching should cover the entire input string (not partial).'], ARRAY['Use dynamic programming.', 'Consider each character in the string and the pattern.'], DEFAULT, NULL);
