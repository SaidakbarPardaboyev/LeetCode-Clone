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

-- right

-- Container With Most Water
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('container-with-most-water', 'maxArea', 'Run', '[1,8,6,2,5,4,8,3,7]', '49', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Run', '[1,1]', '1', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[4,3,2,1,4]', '16', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[1,2,1]', '2', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[1,8,6,2,5,4,8,3,7]', '49', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[1,2,3,4,5,6,7,8,9,10]', '25', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[1,8,6,2,5,4,8,3,7]', '49', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[10,10,10,10,10]', '40', '[]int', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[5,5,5,5,5]', '20', 'array', 'integer'),
('container-with-most-water', 'maxArea', 'Submit', '[1,2,1]', '2', 'array', 'integer');

-- Integer to Roman
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('integer-to-roman', 'intToRoman', 'Run', '3', 'III', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Run', '4', 'IV', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '9', 'IX', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '58', 'LVIII', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '1994', 'MCMXCIV', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '42', 'XLII', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '14', 'XIV', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '444', 'CDXLIV', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '890', 'DCCCXC', 'integer', 'string'),
('integer-to-roman', 'intToRoman', 'Submit', '2019', 'MMXIX', 'integer', 'string');


-- Longest Common Prefix
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('longest-common-prefix', 'longestCommonPrefix', 'Run', '["flower","flow","flight"]', '"fl"', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Run', '["dog","racecar","car"]', '""', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["flower","flow","flight"]', '"fl"', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["dog","racecar","car"]', '""', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["cir","car"]', '"c"', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["ab","a"]', '"a"', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["a","a","b"]', '""', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["flower","flower","flower","flower"]', '"flower"', 'array', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["aca","cba"]', '""', '[]string', 'string'),
('longest-common-prefix', 'longestCommonPrefix', 'Submit', '["a","b"]', '""', '[]string', 'string');

-- 3Sum
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('3sum', 'threeSum', 'Run', '[-1, 0, 1, 2, -1, -4]', '[[-1, -1, 2], [-1, 0, 1]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Run', '[]', '[]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[-1, 0, 1, 2, -1, -4]', '[[-1, -1, 2], [-1, 0, 1]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[0, 0, 0, 0]', '[[0, 0, 0]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[-2, 0, 1, 1, 2]', '[[-2, 0, 2], [-2, 1, 1]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[1, 1, -2]', '[[-2, 1, 1]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[-1, 0, 1, 2, -1, -4]', '[[-1, -1, 2], [-1, 0, 1]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[-1, 0, 1, 2, -1, -4]', '[[-1, -1, 2], [-1, 0, 1]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[0, 0, 0, 0]', '[[0, 0, 0]]', '[]int', '[][]int'),
('3sum', 'threeSum', 'Submit', '[-2, 0, 1, 1, 2]', '[[-2, 0, 2], [-2, 1, 1]]', '[]int', '[][]int');

-- 3Sum Closest
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('3sum-closest', 'threeSumClosest', 'Run', '[-1, 2, 1, -4]', '1', '2', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Run', '[0, 0, 0]', '1', '0', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[-1, 2, 1, -4]', '1', '2', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[0, 0, 0]', '1', '0', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[-1, 2, 1, -4]', '2', '2', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[-1, 2, 1, -4]', '-3', '-1', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[-1, 2, 1, -4]', '0', '0', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[1, 1, 1, 0]', '-100', '2', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[1, 1, 1, 0]', '1', '2', '[]int', 'integer', 'integer'),
('3sum-closest', 'threeSumClosest', 'Submit', '[1, 1, 1, 0]', '100', '3', '[]int', 'integer', 'integer');

-- Letter Combinations of a Phone Number
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Run', '"23"', '["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Run', '"7"', '["p", "q", "r", "s"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"23"', '["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"7"', '["p", "q", "r", "s"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"2"', '["a", "b", "c"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"9"', '["w", "x", "y", "z"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"234"', '["adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"78"', '["pt","pu","pv","qt","qu","qv","rt","ru","rv","st","su","sv"]', 'string', '[]string'),
('letter-combinations-of-a-phone-number', 'letterCombinations', 'Submit', '"567"', '["jmp","jmq","jmr","jms","jnp","jnq","jnr","jns","jop","joq","jor","jos","kmp","kmq","kmr","kms","knp","knq","knr","kns","kop","koq","kor","kos","lmp","lmq","lmr","lms","lnp","lnq","lnr","lns","lop","loq","lor","los"]', 'string', '[]string')

-- 4Sum
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('4sum', 'fourSum', 'Run', '[1, 0, -1, 0, -2, 2]', '0', '[[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Run', '[2, 2, 2, 2, 2]', '8', '[[[2,2,2,2]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[1, 0, -1, 0, -2, 2]', '0', '[[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[0, 0, 0, 0]', '0', '[[[0,0,0,0]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[2, 2, 2, 2, 2]', '8', '[[[2,2,2,2]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[-1,0,1,2,-1,-4]', '-1', '[[[-4,0,1,2],[-1,-1,0,1]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[1,-2,-5,-4,-3,3,3,5]', '-11', '[[[-5,-4,-3,1]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[-3,-2,-1,0,0,1,2,3]', '0', '[[[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[-3,-1,0,2,4,5]', '0', '[]', '[]int', 'integer', '[][]int'),
('4sum', 'fourSum', 'Submit', '[-3,-1,0,2,4,5]', '3', '[[[-3,0,2,4]]]', '[]int', 'integer', '[][]int');

-- Remove Nth Node From End of List
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Run', '[1,2,3,4,5]', '2', '[1,2,3,5]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Run', '[1]', '1', '[]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2,3,4,5]', '2', '[1,2,3,5]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1]', '1', '[]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2]', '1', '[1]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2,3]', '3', '[2,3]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2,3,4,5]', '5', '[2,3,4,5]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2,3,4,5]', '1', '[1,2,3,4]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2,3,4,5]', '4', '[1,3,4,5]', '[]int', 'integer', '[]int'),
('remove-nth-node-from-end-of-list', 'removeNthFromEnd', 'Submit', '[1,2,3,4,5]', '3', '[1,2,4,5]', '[]int', 'integer', '[]int');

-- Valid Parentheses
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('valid-parentheses', 'isValid', 'Run', '"()"', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Run', '"()[]{}"', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Run', '"(]"', 'false', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Run', '"([)]"', 'false', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"()"', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"()[]{}"', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"(]"', 'false', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"([)]"', 'false', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '""', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"{{}}"', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"[[]]"', 'true', 'string', 'boolean'),
('valid-parentheses', 'isValid', 'Submit', '"[([]])"', 'false', 'string', 'boolean');

-- Merge Two Sorted Lists
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, arg2, answer, arg1_type, arg2_type, answer_type)
VALUES
('merge-two-sorted-lists', 'mergeTwoLists', 'Run', '[-10,-10,-9,-4,1,6,6]', '[-7]', '[-10,-10,-9,-7,-4,1,6,6]', 'array', 'array', 'array'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Run', '[1,2,4]', '[1,3,4]', '[1,1,2,3,4,4]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[1,2,4]', '[1,3,4]', '[1,1,2,3,4,4]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[]', '[]', '[]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[]', '[0]', '[0]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[-10,-10,-9,-4,1,6,6]', '[-7]', '[-10,-10,-9,-7,-4,1,6,6]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[1,3,4]', '[1,2,4]', '[1,1,2,3,4,4]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[5]', '[1]', '[1,5]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[1,2,4]', '[]', '[1,2,4]', '[]int', '[]int', '[]int'),
('merge-two-sorted-lists', 'mergeTwoLists', 'Submit', '[1,3,4]', '[]', '[1,3,4]', '[]int', '[]int', '[]int');


-- Generate Parentheses
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('generate-parentheses', 'generateParenthesis', 'Run', '3', '["((()))","(()())","(())()","()(())","()()()"]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Run', '1', '["()"]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Submit', '3', '["((()))","(()())","(())()","()(())","()()()"]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Submit', '2', '["(())","()()"]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Submit', '1', '["()"]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Submit', '4', '["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Submit', '0', '[""]', 'integer', '[]string'),
('generate-parentheses', 'generateParenthesis', 'Submit', '5', '["((((()))))","(((()())))","(((())()))","(((()))())","(((())))()","((()(())))","((()()()))","((()())())","((()()))()","((())(()))","((())()())","((())())()","((()))(())","((()))()()","(()((())))","(()(()()))","(()(())())","(()(()))()","(()()(()))","(()()()())","(()()())()","(()()()())","(()()()()","(())((()))","(())(()())","(())(())()","(())()(()","(())()()()","()(((())))","()((()()))","()((())())","()((()))()","()(()(()))","()(()()())","()(()())()","()(())(())","()(())()()","()()((()))","()()(()())","()()(())()","()()()(())","()()()()()"]', 'integer', '[]string');

-- Merge k Sorted Lists
INSERT INTO testcases_non_btree (problem_title, function_name, run_or_submit, arg1, answer, arg1_type, answer_type)
VALUES
('merge-k-sorted-lists', 'mergeKLists', 'Run', '[[1,4,5],[1,3,4],[2,6]]', '[1,1,2,3,4,4,5,6]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Run', '[]', '[]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[1,4,5],[1,3,4],[2,6]]', '[1,1,2,3,4,4,5,6]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[1,2,3],[4,5,6],[7,8,9]]', '[1,2,3,4,5,6,7,8,9]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[],[],[]]', '[]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[0,2,5],[1,3,4],[6,7,8]]', '[0,1,2,3,4,5,6,7,8]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[0],[2],[5]]', '[0,2,5]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[1,4,5],[1,3,4],[2,6],[0,0,0,0,0]]', '[0,0,0,0,0,1,1,2,3,4,4,5,6]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[1,4,5],[1,3,4],[2,6],[-1,-1,-1,-1]]', '[-1,-1,-1,-1,1,1,2,3,4,4,5,6]', '[][]int', '[]int'),
('merge-k-sorted-lists', 'mergeKLists', 'Submit', '[[10,20,30],[15,25,35],[5,15,25]]', '[5,10,15,15,20,25,25,30,35]', '[][]int', '[]int');
