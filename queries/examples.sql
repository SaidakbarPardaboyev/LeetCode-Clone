create table examples (
	id uuid primary key unique default gen_random_uuid() not null,
	problem_title varchar references problems(title) not null,
	input text not null,
	output text not null,
	explanation text,
    created_at timestamp default now() not null,
    updated_at timestamp,
    deleted_at timestamp
);

-- Two Sum
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('two-sum', 'nums = [2, 7, 11, 15], target = 9', '[0, 1]', 'Because nums[0] + nums[1] == 9', DEFAULT, NULL),
('two-sum', 'nums = [3, 2, 4], target = 6', '[1, 2]', 'Because nums[1] + nums[2] == 6', DEFAULT, NULL),
('two-sum', 'nums = [3, 3], target = 6', '[0, 1]', 'Because nums[0] + nums[1] == 6', DEFAULT, NULL);

-- Add Two Numbers
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('add-two-numbers', 'l1 = [2, 4, 3], l2 = [5, 6, 4]', '[7, 0, 8]', '342 + 465 = 807', DEFAULT, NULL),
('add-two-numbers', 'l1 = [0], l2 = [0]', '[0]', '0 + 0 = 0', DEFAULT, NULL),
('add-two-numbers', 'l1 = [9, 9, 9, 9, 9, 9, 9], l2 = [9, 9, 9, 9]', '[8, 9, 9, 9, 0, 0, 0, 1]', '9999999 + 9999 = 10009998', DEFAULT, NULL);

-- Longest Substring Without Repeating Characters
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('longest-substring-without-repeating-characters', 's = "abcabcbb"', '3', 'The answer is "abc", with the length of 3.', DEFAULT, NULL),
('longest-substring-without-repeating-characters', 's = "bbbbb"', '1', 'The answer is "b", with the length of 1.', DEFAULT, NULL),
('longest-substring-without-repeating-characters', 's = "pwwkew"', '3', 'The answer is "wke", with the length of 3.', DEFAULT, NULL);

-- Median of Two Sorted Arrays
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('median-of-two-sorted-arrays', 'nums1 = [1, 3], nums2 = [2]', '2.0', 'The median is 2.0', DEFAULT, NULL),
('median-of-two-sorted-arrays', 'nums1 = [1, 2], nums2 = [3, 4]', '2.5', 'The median is (2 + 3) / 2 = 2.5', DEFAULT, NULL),
('median-of-two-sorted-arrays', 'nums1 = [0, 0], nums2 = [0, 0]', '0.0', 'The median is 0.0', DEFAULT, NULL);

-- Longest Palindromic Substring
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('longest-palindromic-substring', 's = "babad"', '"bab" or "aba"', 'Both "bab" and "aba" are palindromic substrings.', DEFAULT, NULL),
('longest-palindromic-substring', 's = "cbbd"', '"bb"', 'The longest palindromic substring is "bb".', DEFAULT, NULL),
('longest-palindromic-substring', 's = "a"', '"a"', 'The longest palindromic substring is "a".', DEFAULT, NULL);

-- Zigzag Conversion
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('zigzag-conversion', 's = "PAYPALISHIRING", numRows = 3', '"PAHNAPLSIIGYIR"', 'Zigzag pattern on 3 rows results in "PAHNAPLSIIGYIR".', DEFAULT, NULL),
('zigzag-conversion', 's = "PAYPALISHIRING", numRows = 4', '"PINALSIGYAHRPI"', 'Zigzag pattern on 4 rows results in "PINALSIGYAHRPI".', DEFAULT, NULL),
('zigzag-conversion', 's = "A", numRows = 1', '"A"', 'Zigzag pattern with 1 row is just "A".', DEFAULT, NULL);

-- Reverse Integer
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('reverse-integer', 'x = 123', '321', 'Reversed integer of 123 is 321.', DEFAULT, NULL),
('reverse-integer', 'x = -123', '-321', 'Reversed integer of -123 is -321.', DEFAULT, NULL),
('reverse-integer', 'x = 120', '21', 'Reversed integer of 120 is 21.', DEFAULT, NULL);

-- String to Integer (atoi)
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('string-to-integer-(atoi)', 's = "42"', '42', 'String "42" is converted to integer 42.', DEFAULT, NULL),
('string-to-integer-(atoi)', 's = "   -42"', '-42', 'String "   -42" is converted to integer -42.', DEFAULT, NULL),
('string-to-integer-(atoi)', 's = "4193 with words"', '4193', 'String "4193 with words" is converted to integer 4193.', DEFAULT, NULL);

-- Palindrome Number
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('palindrome-number', 'x = 121', 'true', '121 reads as 121 from left to right and from right to left.', DEFAULT, NULL),
('palindrome-number', 'x = -121', 'false', '-121 reads as 121- from left to right and from right to left.', DEFAULT, NULL),
('palindrome-number', 'x = 10', 'false', '10 reads as 01 from right to left.', DEFAULT, NULL);

-- Regular Expression Matching
INSERT INTO examples (problem_title, input, output, explanation, updated_at, deleted_at)
VALUES
('regular-expression-matching', 's = "aa", p = "a*"', 'true', '"a*" matches any sequence of "a"s.', DEFAULT, NULL),
('regular-expression-matching', 's = "aa", p = "a"', 'false', '"a" does not match the string "aa".', DEFAULT, NULL),
('regular-expression-matching', 's = "ab", p = ".*"', 'true', '".*" matches any string.', DEFAULT, NULL);
