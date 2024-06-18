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

-- Two Sum
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('two-sum', 'twoSum', 'Run', '[2, 7, 11, 15]', '9', '[0, 1]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Run', '[3, 2, 4]', '6', '[1, 2]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Run', '[3, 3]', '6', '[0, 1]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[1, 2, 3]', '4', '[0, 2]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[5, 6, 7]', '11', '[0, 2]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[1, 1, 1, 1]', '2', '[0, 1]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[1, 2]', '3', '[0, 1]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[2, 5, 5, 11]', '10', '[1, 2]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[3, 6, 9, 12]', '15', '[0, 3]', '[]int', 'int', '[]int'),
('two-sum', 'twoSum', 'Submit', '[1, 3, 5, 7]', '8', '[1, 2]', '[]int', 'int', '[]int');

-- Add Two Numbers
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('add-two-numbers', 'addTwoNumbers', 'Run', '[2, 4, 3]', '[5, 6, 4]', '[7, 0, 8]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Run', '[0]', '[0]', '[0]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Run', '[9, 9, 9, 9, 9, 9, 9]', '[9, 9, 9, 9]', '[8, 9, 9, 9, 0, 0, 0, 1]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[2, 4, 9]', '[5, 6, 4, 9]', '[7, 0, 4, 0, 1]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[0, 1]', '[0, 1, 2]', '[0, 2, 2]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[5]', '[5]', '[0, 1]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[1, 8]', '[0]', '[1, 8]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[9, 9]', '[1]', '[0, 0, 1]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[1]', '[9, 9, 9]', '[0, 0, 0, 1]', '[]int', '[]int', '[]int'),
('add-two-numbers', 'addTwoNumbers', 'Submit', '[9, 9, 9]', '[1]', '[0, 0, 0, 1]', '[]int', '[]int', '[]int');

-- Longest Substring Without Repeating Characters
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Run', '"abcabcbb"', '3', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Run', '"bbbbb"', '1', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Run', '"pwwkew"', '3', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '""', '0', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '" "', '1', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '"au"', '2', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '"dvdf"', '3', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '"anviaj"', '5', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '"abccabb"', '3', 'string', 'int'),
('longest-substring-without-repeating-characters', 'lengthOfLongestSubstring', 'Submit', '"abacabcbb"', '3', 'string', 'int');

-- Median of Two Sorted Arrays
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Run', '[1, 3]', '[2]', '2.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Run', '[1, 2]', '[3, 4]', '2.5', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Run', '[0, 0]', '[0, 0]', '0.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Run', '[]', '[1]', '1.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Submit', '[2]', '[]', '2.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Submit', '[3, 4]', '[1, 2]', '2.5', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Submit', '[1, 1, 1, 1]', '[1, 1, 1, 1]', '1.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Submit', '[1]', '[1, 1, 1, 1]', '1.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Submit', '[1, 2, 3, 4]', '[5]', '3.0', '[]int', '[]int', 'float'),
('median-of-two-sorted-arrays', 'findMedianSortedArrays', 'Submit', '[5]', '[1, 2, 3, 4]', '3.0', '[]int', '[]int', 'float');

-- Longest Palindromic Substring
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('longest-palindromic-substring', 'longestPalindrome', 'Run', '"babad"', '"bab"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Run', '"cbbd"', '"bb"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Run', '"a"', '"a"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"ac"', '"a"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"racecar"', '"racecar"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"abb"', '"bb"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"aabb"', '"aa"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"aaaa"', '"aaaa"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"bananas"', '"anana"', 'string', 'string'),
('longest-palindromic-substring', 'longestPalindrome', 'Submit', '"abcda"', '"a"', 'string', 'string');

-- Zigzag Conversion
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('zigzag-conversion', 'convert', 'Run', '"PAYPALISHIRING"', '3', '"PAHNAPLSIIGYIR"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Run', '"PAYPALISHIRING"', '4', '"PINALSIGYAHRPI"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"A"', '1', '"A"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"AB"', '1', '"AB"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"ABCD"', '2', '"ACBD"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"ABCDE"', '3', '"AEBDC"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"ABCDEFG"', '4', '"AGBFCED"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"A"', '2', '"A"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"ABCDE"', '5', '"ABCDE"', 'string', 'int', 'string'),
('zigzag-conversion', 'convert', 'Submit', '"ABCDEFGHIJKL"', '4', '"AGMBFHLCEIKDJ"', 'string', 'int', 'string');

-- Reverse Integer
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('reverse-integer', 'reverse', 'Run', '123', '321', 'int', 'int'),
('reverse-integer', 'reverse', 'Run', '-123', '-321', 'int', 'int'),
('reverse-integer', 'reverse', 'Run', '120', '21', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '0', '0', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '1534236469', '0', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '-1534236469', '0', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '100', '1', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '-100', '-1', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '101', '101', 'int', 'int'),
('reverse-integer', 'reverse', 'Submit', '123456789', '987654321', 'int', 'int');

-- String to Integer (atoi)
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('string-to-integer-(atoi)', 'myAtoi', 'Run', '"42"', '42', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Run', '"   -42"', '-42', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Run', '"4193 with words"', '4193', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Run', '"words and 987"', '0', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Submit', '"-91283472332"', '-2147483648', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Submit', '"3.14159"', '3', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Submit', '"-91283472332"', '-2147483648', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Submit', '"-13+8"', '-13', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Submit', '"00000-42a1234"', '0', 'string', 'int'),
('string-to-integer-(atoi)', 'myAtoi', 'Submit', '"-5-"', '-5', 'string', 'int');

-- Palindrome Number
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('palindrome-number', 'isPalindrome', 'Run', '121', 'true', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Run', '-121', 'false', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Run', '10', 'false', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '-101', 'false', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '0', 'true', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '12321', 'true', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '1000021', 'false', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '1', 'true', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '11', 'true', 'int', 'boolean'),
('palindrome-number', 'isPalindrome', 'Submit', '22', 'true', 'int', 'boolean');

-- Regular Expression Matching
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('regular-expression-matching', 'isMatch', 'Run', '"aa"', '"a"', 'false', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Run', '"aa"', '"a*"', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"ab"', '".*"', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"aab"', '"c*a*b"', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"mississippi"', '"mis*is*p*."', 'false', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"mississippi"', '"mis*is*ip*."', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"ab"', '"a*b"', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"abcd"', '".*d"', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"abcd"', '"abcd*"', 'true', 'string', 'string', 'boolean'),
('regular-expression-matching', 'isMatch', 'Submit', '"a"', '"ab*"', 'true', 'string', 'string', 'boolean');