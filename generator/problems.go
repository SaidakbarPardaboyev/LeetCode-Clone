package generator

import (
	"database/sql"
	"leetcode/model"
	"leetcode/storage/postgres"
)

var problems = []model.Problem{
	{
		QuestionNumber:  1,
		Title:           "Two Sum",
		DifficultyLevel: "Easy",
		Description:     "Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.",
		Examples: []string{
			"Input: nums = [2,7,11,15], target = 9\nOutput: [0,1]\nExplanation: Because nums[0] + nums[1] == 9, we return [0, 1].",
			"Input: nums = [3,2,4], target = 6\nOutput: [1,2]",
			"Input: nums = [3,3], target = 6\nOutput: [0,1]",
		},
		Hints: []string{
			"Use a hashmap to store each element's index and check if the complement exists.",
			"Consider edge cases such as no solution or multiple solutions.",
		},
		Constraints: []string{
			"You may assume that each input would have exactly one solution.",
			"You may not use the same element twice.",
			"The order of the result does not matter.",
		},
	},
	{
		QuestionNumber: 2,
		Title:          "Add Two Numbers",
		DifficultyLevel: "Medium",
		Description:    "You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.",
		Examples: []string{
			"Input: l1 = [2,4,3], l2 = [5,6,4]\nOutput: [7,0,8]\nExplanation: 342 + 465 = 807.",
			"Input: l1 = [0], l2 = [0]\nOutput: [0]\nExplanation: 0 + 0 = 0.",
		},
		Hints: []string{
			"Initialize a dummy head for the result list.",
			"Use a carry variable to keep track of the carry from the addition of each digit.",
			"Traverse both lists and add corresponding digits along with the carry.",
		},
		Constraints: []string{
			"The number of nodes in each linked list is in the range [1, 100].",
			"0 <= Node.val <= 9",
			"It is guaranteed that the list represents a number that does not have leading zeros.",
		},
	},
	{
		QuestionNumber: 3,
		Title:          "Longest Substring Without Repeating Characters",
		DifficultyLevel: "Medium",
		Description:    "Given a string s, find the length of the longest substring without repeating characters.",
		Examples: []string{
			"Input: s = \"abcabcbb\"\nOutput: 3\nExplanation: The answer is \"abc\", with the length of 3.",
			"Input: s = \"bbbbb\"\nOutput: 1\nExplanation: The answer is \"b\", with the length of 1.",
			"Input: s = \"pwwkew\"\nOutput: 3\nExplanation: The answer is \"wke\", with the length of 3. Notice that the answer must be a substring, \"pwke\" is a subsequence and not a substring.",
		},
		Hints: []string{
			"Use a sliding window to keep track of the current substring without repeating characters.",
			"Use a hash map to store the last index of each character.",
			"Update the start of the window when a repeating character is found.",
		},
		Constraints: []string{
			"0 <= s.length <= 5 * 10^4",
			"s consists of English letters, digits, symbols and spaces.",
		},
	},
	{
		QuestionNumber: 4,
		Title:          "Median of Two Sorted Arrays",
		DifficultyLevel: "Hard",
		Description:    "Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.",
		Examples: []string{
			"Input: nums1 = [1,3], nums2 = [2]\nOutput: 2.00000\nExplanation: merged array = [1,2,3] and median is 2.",
			"Input: nums1 = [1,2], nums2 = [3,4]\nOutput: 2.50000\nExplanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.",
			"Input: nums1 = [0,0], nums2 = [0,0]\nOutput: 0.00000",
			"Input: nums1 = [], nums2 = [1]\nOutput: 1.00000",
			"Input: nums1 = [2], nums2 = []\nOutput: 2.00000",
		},
		Hints: []string{
			"Try to come up with a solution of logarithmic time complexity.",
			"Use binary search to partition the arrays such that the left part contains the smallest half of the elements.",
			"Ensure the partition is balanced to find the median.",
		},
		Constraints: []string{
			"nums1.length == m",
			"nums2.length == n",
			"0 <= m <= 1000",
			"0 <= n <= 1000",
			"1 <= m + n <= 2000",
			"-10^6 <= nums1[i], nums2[i] <= 10^6",
		},
	},
	{
        QuestionNumber: 5,
        Title:          "Longest Palindromic Substring",
        DifficultyLevel: "Medium",
        Description:    "Given a string s, return the longest palindromic substring in s.",
        Examples: []string{
            "Input: s = \"babad\"\nOutput: \"bab\"\nNote: \"aba\" is also a valid answer.",
            "Input: s = \"cbbd\"\nOutput: \"bb\"",
            "Input: s = \"a\"\nOutput: \"a\"",
            "Input: s = \"ac\"\nOutput: \"a\"",
        },
        Hints: []string{
            "Expand Around Center: You could use an approach that considers each character (and between characters) as the center of a potential palindrome and expand outwards.",
            "Dynamic Programming: Use a 2D table to keep track of palindromic substrings.",
            "Optimize with Manacher's Algorithm for linear time complexity.",
        },
        Constraints: []string{
            "1 <= s.length <= 1000",
            "s consist of only digits and English letters.",
        },
    },
	{
        QuestionNumber: 6,
        Title:          "Zigzag Conversion",
        DifficultyLevel: "Medium",
        Description:    "The string \"PAYPALISHIRING\" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)\nP   A   H   N\nA P L S I I G\nY   I   R\nAnd then read line by line: \"PAHNAPLSIIGYIR\"\nWrite the code that will take a string and make this conversion given a number of rows:",
        Examples: []string{
            "Input: s = \"PAYPALISHIRING\", numRows = 3\nOutput: \"PAHNAPLSIIGYIR\"",
            "Input: s = \"PAYPALISHIRING\", numRows = 4\nOutput: \"PINALSIGYAHRPI\"\nExplanation:\nP     I    N\nA   L S  I G\nY A   H R\nP     I",
            "Input: s = \"A\", numRows = 1\nOutput: \"A\"",
        },
        Hints: []string{
            "You can create an array of strings, one for each row.",
            "Iterate through the characters of the string, adding each character to the appropriate row.",
            "The pattern alternates between moving down and moving up.",
        },
        Constraints: []string{
            "1 <= s.length <= 1000",
            "s consists of English letters (lower-case and upper-case), ',' and '.'.",
            "1 <= numRows <= 1000",
        },
    },
	{
        QuestionNumber: 7,
        Title:          "Reverse Integer",
        DifficultyLevel: "Medium",
        Description:    "Given a 32-bit signed integer, reverse digits of an integer.",
        Examples: []string{
            "Input: 123\nOutput: 321",
            "Input: -123\nOutput: -321",
            "Input: 120\nOutput: 21",
        },
        Hints: []string{
            "Think about how you would do it on paper.",
            "Ensure that your integer does not overflow when reversed.",
        },
        Constraints: []string{
            "The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.",
        },
    },
	{
        QuestionNumber: 8,
        Title:          "String to Integer (atoi)",
        DifficultyLevel: "Medium",
        Description:    "Implement the `atoi` function, which converts a string to an integer.",
        Examples: []string{
            `Input: "42"\nOutput: 42`,
            `Input: "   -42"\nOutput: -42`,
            `Input: "4193 with words"\nOutput: 4193`,
            `Input: "words and 987"\nOutput: 0`,
            `Input: "-91283472332"\nOutput: -2147483648 (clamped to the minimum 32-bit signed integer value)`,
        },
        Hints: []string{
            "Consider all possible edge cases: whitespaces, signs, overflow, invalid characters, etc.",
            "You might want to use the `long` data type to handle overflow cases before clamping the result.",
        },
        Constraints: []string{
            "The function discards all whitespace characters at the beginning of the string.",
            "The function processes an optional initial plus or minus sign followed by as many numerical digits as possible and interprets them as a numerical value.",
            "The function stops processing as soon as it encounters a non-numerical character.",
            "If the string is empty or does not contain valid digits, return 0.",
            "If the numerical value is out of the range of a 32-bit signed integer, clamp the value to the range [−2^31, 2^31 − 1].",
        },
    },
	{
        QuestionNumber: 9,
        Title:          "Palindrome Number",
        DifficultyLevel: "Easy",
        Description:    "Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.",
        Examples: []string{
            `Input: 121\nOutput: true`,
            `Input: -121\nOutput: false (from left to right, it reads "-121". From right to left, it becomes "121-". Therefore, it is not a palindrome.)`,
            `Input: 10\nOutput: false (reads 01 from right to left. Therefore, it is not a palindrome.)`,
        },
        Hints: []string{
            "Negative numbers are not palindromes.",
            "Reverse the number and compare it with the original number.",
        },
        Constraints: []string{
            "Follow up: Could you solve it without converting the integer to a string?",
        },
    },
	{
        QuestionNumber: 10,
        Title:          "Regular Expression Matching",
        DifficultyLevel: "Hard",
        Description:    "Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.",
        Examples: []string{
            `Input: s = "aa", p = "a"\nOutput: false\nExplanation: "a" does not match the entire string "aa".`,
            `Input: s = "aa", p = "a*"\nOutput: true\nExplanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".`,
            `Input: s = "ab", p = ".*"\nOutput: true\nExplanation: ".*" means "zero or more (*) of any character (.)".`,
            `Input: s = "aab", p = "c*a*b"\nOutput: true\nExplanation: 'c' can be repeated 0 times, 'a' can be repeated 1 time. Therefore, it matches "aab".`,
            `Input: s = "mississippi", p = "mis*is*p*."\nOutput: false`,
        },
        Hints: []string{
            "Consider the cases for '.' and '*'.",
            "Dynamic Programming (DP) can be useful here.",
        },
        Constraints: []string{
            "The input strings consist of only lowercase English letters.",
            "The length of both strings s and p is <= 20.",
            "It is guaranteed that s and p are non-empty.",
        },
    },
	{
        QuestionNumber: 11,
        Title:          "Container With Most Water",
        DifficultyLevel: "Medium",
        Description:    "Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.",
        Examples: []string{
            "Input: height = [1,8,6,2,5,4,8,3,7]\nOutput: 49\nExplanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.",
            "Input: height = [1,1]\nOutput: 1",
            "Input: height = [4,3,2,1,4]\nOutput: 16",
            "Input: height = [1,2,1]\nOutput: 2",
        },
        Hints: []string{
            "The aim is to maximize the area formed between the vertical lines.",
            "Start with two pointers approach.",
        },
        Constraints: []string{
            "n == height.length",
            "2 <= n <= 10^5",
            "0 <= height[i] <= 10^4",
        },
    },
	{
        QuestionNumber: 12,
        Title:          "Integer to Roman",
        DifficultyLevel: "Medium",
        Description:    "Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.",
        Examples: []string{
            "Input: num = 3\nOutput: \"III\"",
            "Input: num = 4\nOutput: \"IV\"",
            "Input: num = 9\nOutput: \"IX\"",
            "Input: num = 58\nOutput: \"LVIII\"\nExplanation: L = 50, V = 5, III = 3.",
            "Input: num = 1994\nOutput: \"MCMXCIV\"\nExplanation: M = 1000, CM = 900, XC = 90, IV = 4.",
        },
        Hints: []string{
            "Start with the largest value and reduce the problem to smaller subproblems.",
            "Use predefined mappings of roman numeral symbols to integers.",
        },
        Constraints: []string{
            "1 <= num <= 3999",
        },
    },
	{
        QuestionNumber: 13,
        Title:          "Roman to Integer",
        DifficultyLevel: "Easy",
        Description:    "Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.",
        Examples: []string{
            "Input: s = \"III\"\nOutput: 3",
            "Input: s = \"IV\"\nOutput: 4",
            "Input: s = \"IX\"\nOutput: 9",
            "Input: s = \"LVIII\"\nOutput: 58\nExplanation: L = 50, V= 5, III = 3.",
            "Input: s = \"MCMXCIV\"\nOutput: 1994\nExplanation: M = 1000, CM = 900, XC = 90, IV = 4.",
        },
        Hints: []string{
            "Use a hashmap to store the mapping of roman numerals to integers.",
            "Iterate through the string and compare current and next characters to determine the integer value.",
        },
        Constraints: []string{
            "1 <= s.length <= 15",
            "s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').",
            "It is guaranteed that s is a valid roman numeral in the range [1, 3999].",
        },
    },
	{
        QuestionNumber: 14,
        Title:          "Longest Common Prefix",
        DifficultyLevel: "Easy",
        Description:    "Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string \"\".",
        Examples: []string{
            "Input: strs = [\"flower\",\"flow\",\"flight\"]\nOutput: \"fl\"",
            "Input: strs = [\"dog\",\"racecar\",\"car\"]\nOutput: \"\"",
        },
        Hints: []string{
            "Start with the first string as the initial prefix candidate.",
            "Compare this prefix candidate with each subsequent string, updating the prefix candidate as you find common prefixes.",
        },
        Constraints: []string{
            "1 <= strs.length <= 200",
            "0 <= strs[i].length <= 200",
            "strs[i] consists of only lower-case English letters.",
        },
    },
	{
        QuestionNumber: 15,
        Title:          "3Sum",
        DifficultyLevel: "Medium",
        Description:    "Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0. Notice that the solution set must not contain duplicate triplets.",
        Examples: []string{
            "Input: nums = [-1,0,1,2,-1,-4]\nOutput: [[-1,-1,2],[-1,0,1]]",
            "Input: nums = []\nOutput: []",
            "Input: nums = [0]\nOutput: []",
        },
        Hints: []string{
            "Sort the array to use two pointers technique.",
            "Iterate through the array, using the current element as the first element of the triplet.",
            "Use two pointers to find the other two elements that sum up to the negative of the current element.",
        },
        Constraints: []string{
            "0 <= nums.length <= 3000",
            "-10^5 <= nums[i] <= 10^5",
        },
    },
	{
        QuestionNumber: 16,
        Title:          "3Sum Closest",
        DifficultyLevel: "Medium",
        Description:    "Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.",
        Examples: []string{
            "Input: nums = [-1,2,1,-4], target = 1\nOutput: 2\nExplanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).",
            "Input: nums = [0,0,0], target = 1\nOutput: 0",
        },
        Hints: []string{
            "Sort the array to use two pointers technique.",
            "Iterate through the array, using the current element as the first element of the triplet.",
            "Use two pointers to find the other two elements that sum up to the target minus the current element.",
        },
        Constraints: []string{
            "3 <= nums.length <= 10^3",
            "-10^3 <= nums[i] <= 10^3",
            "-10^4 <= target <= 10^4",
        },
    },
	{
        QuestionNumber: 17,
        Title:          "Letter Combinations of a Phone Number",
        DifficultyLevel: "Medium",
        Description:    "Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.",
        Examples: []string{
            "Input: digits = \"23\"\nOutput: [\"ad\",\"ae\",\"af\",\"bd\",\"be\",\"bf\",\"cd\",\"ce\",\"cf\"]",
            "Input: digits = \"\"\nOutput: []",
            "Input: digits = \"2\"\nOutput: [\"a\",\"b\",\"c\"]",
        },
        Hints: []string{
            "Use backtracking to generate all possible combinations.",
            "Create a mapping from each digit to its corresponding letters on a phone keypad.",
            "Iterate through each digit in the input string and recursively generate combinations by appending each possible letter.",
        },
        Constraints: []string{
            "0 <= digits.length <= 4",
            "digits[i] is a digit in the range ['2', '9'].",
        },
    },
	{
        QuestionNumber: 18,
        Title:          "4Sum",
        DifficultyLevel: "Medium",
        Description:    "Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:\n1. 0 <= a, b, c, d < n\n2. a, b, c, and d are distinct.\n3. nums[a] + nums[b] + nums[c] + nums[d] == target\nYou may return the answer in any order.",
        Examples: []string{
            "Input: nums = [1,0,-1,0,-2,2], target = 0\nOutput: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]",
            "Input: nums = [2,2,2,2,2], target = 8\nOutput: [[2,2,2,2]]",
        },
        Hints: []string{
            "Sort the array, and use recursion with backtracking to find all unique quadruplets.",
            "Avoid duplicates by skipping identical numbers when iterating through the array.",
        },
        Constraints: []string{
            "1 <= nums.length <= 200",
            "-10^9 <= nums[i] <= 10^9",
            "-10^9 <= target <= 10^9",
        },
    },
	{
        QuestionNumber: 19,
        Title:          "Remove Nth Node From End of List",
        DifficultyLevel: "Medium",
        Description:    "Given the head of a linked list, remove the nth node from the end of the list and return its head.",
        Examples: []string{
            "Input: head = [1,2,3,4,5], n = 2\nOutput: [1,2,3,5]",
            "Input: head = [1], n = 1\nOutput: []",
            "Input: head = [1,2], n = 1\nOutput: [1]",
        },
        Hints: []string{
            "Use two pointers to find the nth node from the end in one pass.",
            "Handle edge cases where the list contains only one node or removing the first node.",
        },
        Constraints: []string{
            "The number of nodes in the list is sz.",
            "1 <= sz <= 30",
            "0 <= Node.val <= 100",
            "1 <= n <= sz",
        },
    },
	{
        QuestionNumber: 20,
        Title:          "Valid Parentheses",
        DifficultyLevel: "Easy",
        Description:    "Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.",
        Examples: []string{
            "Input: s = \"()\"\nOutput: true",
            "Input: s = \"()[]{}\"\nOutput: true",
            "Input: s = \"(]\"\nOutput: false",
            "Input: s = \"([)]\"\nOutput: false",
            "Input: s = \"{[]}\"\nOutput: true",
        },
        Hints: []string{
            "Use a stack to keep track of opening brackets.",
            "Iterate through the string, pushing opening brackets onto the stack and popping from the stack when encountering a closing bracket.",
            "Check if the stack is empty at the end to determine validity.",
        },
        Constraints: []string{
            "1 <= s.length <= 104",
            "s consists of parentheses only '()[]{}'.",
        },
    },
	{
        QuestionNumber: 21,
        Title:          "Merge Two Sorted Lists",
        DifficultyLevel: "Easy",
        Description:    "Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.",
        Examples: []string{
            "Input: l1 = [1,2,4], l2 = [1,3,4]\nOutput: [1,1,2,3,4,4]",
            "Input: l1 = [], l2 = []\nOutput: []",
            "Input: l1 = [], l2 = [0]\nOutput: [0]",
        },
        Hints: []string{
            "Use a dummy node to simplify the code.",
            "Iterate through both lists and compare the values, appending the smaller value to the result list.",
            "Handle cases where one list is exhausted before the other.",
        },
        Constraints: []string{
            "The number of nodes in both lists is in the range [0, 50].",
            "-100 <= Node.val <= 100",
            "Both l1 and l2 are sorted in non-decreasing order.",
        },
    },
	{
        QuestionNumber: 22,
        Title:          "Generate Parentheses",
        DifficultyLevel: "Medium",
        Description:    "Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.",
        Examples: []string{
            "Input: n = 3\nOutput: [\"((()))\",\"(()())\",\"(())()\",\"()(())\",\"()()()\"]",
            "Input: n = 1\nOutput: [\"()\"]",
        },
        Hints: []string{
            "Think recursively: to generate parentheses for n pairs, you can add '(' or ')' as long as it doesn't exceed the count of each type.",
            "Use backtracking to explore all possible combinations, keeping track of the balance of '(' and ')'.",
        },
        Constraints: []string{
            "1 <= n <= 8",
        },
    },
	{
        QuestionNumber: 23,
        Title:          "Merge k Sorted Lists",
        DifficultyLevel: "Hard",
        Description:    "Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.",
        Examples: []string{
            "Input: lists = [[1,4,5],[1,3,4],[2,6]]  Output: [1,1,2,3,4,4,5,6]",
            "Input: lists = []  Output: []",
            "Input: lists = [[]]  Output: []",
        },
        Hints: []string{
            "Use a min-heap to efficiently get the smallest element among all lists.",
            "Merge lists in a divide-and-conquer approach, reducing the number of lists to merge at each step until only one list remains.",
        },
        Constraints: []string{
            "0 <= k <= 10^4",
            "0 <= lists[i].length <= 500",
            "-10^4 <= lists[i][j] <= 10^4",
            "lists[i] is sorted in ascending order.",
            "The sum of lists[i].length won't exceed 10^4.",
        },
    },
	{
        QuestionNumber: 24,
        Title:          "Swap Nodes in Pairs",
        DifficultyLevel: "Medium",
        Description:    "Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed).",
        Examples: []string{
            "Input: head = [1,2,3,4]  Output: [2,1,4,3]",
            "Input: head = []  Output: []",
            "Input: head = [1]  Output: [1]",
        },
        Hints: []string{
            "If you swap nodes A and B, you need to make sure their previous node points to B instead of A, and B points to A instead of its next node.",
            "Handle edge cases such as an empty list or a list with only one node.",
        },
        Constraints: []string{
            "The number of nodes in the list is in the range [0, 100].",
            "0 <= Node.val <= 100",
        },
    },
	{
        QuestionNumber: 25,
        Title:          "Reverse Nodes in k-Group",
        DifficultyLevel: "Hard",
        Description:    "Given a linked list, reverse the nodes of a linked list k at a time and return its modified list. k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k, then left-out nodes, in the end, should remain as it is.",
        Examples: []string{
            "Input: head = [1,2,3,4,5], k = 2  Output: [2,1,4,3,5]",
            "Input: head = [1,2,3,4,5], k = 3  Output: [3,2,1,4,5]",
            "Input: head = [1,2,3,4,5], k = 1  Output: [1,2,3,4,5]",
            "Input: head = [1], k = 1  Output: [1]",
        },
        Hints: []string{
            "If you have less than k nodes left in the linked list, you should reverse them as well.",
            "Use recursion to solve the problem, breaking down the problem into smaller parts (reversing each k-group of nodes).",
        },
        Constraints: []string{
            "The number of nodes in the list is in the range sz.",
            "1 <= sz <= 5000",
            "0 <= Node.val <= 1000",
            "1 <= k <= sz",
        },
    },
	{
		QuestionNumber:  26,
		Title:           "Remove Duplicates from Sorted Array",
		DifficultyLevel: "Easy",
		Description: `
		Given an integer array nums sorted in non-decreasing order, 
		remove the duplicates in-place such that each unique element appears only once.
		The relative order of the elements should be kept the same. 
		Then return the number of unique elements in nums.
		Consider the number of unique elements of nums to be k, to get accepted, you need 
		to do the following things:

		Change the array nums susersuch that the first k elements of nums contain 
		the unique elements in the order they were present in nums initially. 
		The remaining elements of nums are not important as well as the size of nums.
		Return k.

		Custom Judge:

		The judge will test your solution with the following code:

		int[] nums = [...]; // Input array
		int[] expectedNums = [...]; // The expected answer with correct length

		int k = removeDuplicates(nums); // Calls your implementation

		assert k == expectedNums.length;
		for (int i = 0; i < k; i++) {
			assert nums[i] == expectedNums[i];
		}

		If all assertions pass, then your solution will be accepted.`,

		Examples: []string{
			`Input: nums = [1,1,2]
			Output: 2, nums = [1,2,_]
			Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
			It does not matter what you leave beyond the returned k (hence they are underscores).`,
			`Input: nums = [0,0,1,1,1,2,2,3,3,4]
			Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
			Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
			It does not matter what you leave beyond the returned k (hence they are underscores).`,
		},
		Hints: []string{
			`In this problem, the key point to focus on is
			the input array being sorted. As far as duplicate
			elements are concerned, what is their positioning 
			in the array when the given array is sorted? Look at 
			the image below for the answer. If we know the position 
			of one of the elements, do we also know the positioning of all the duplicate elements? `,

			`We need to modify the array in-place and 
			the size of the final array would potentially 
			be smaller than the size of the input array. So, we 
			ought to use a two-pointer approach here. One, that would 
			keep track of the current element in the original array and another 
			one for just the unique elements.`,

			`Essentially, once an element is encountered, you simply
			need to bypass its duplicates and move on to the next unique element.`,
		},
		Constraints: []string{
			`1 <= nums.length <= 3 * 104`,
			`-100 <= nums[i] <= 100`,
			`nums is sorted in non-decreasing order.`,
		},
	},
	{
		QuestionNumber:  27,
		Title:           "Remove Element",
		DifficultyLevel: "Easy",
		Description:     "Given an array nums and a value val, remove all instances of that value in-place and return the new length.",
		Examples: []string{
			"Input: nums = [3,2,2,3], val = 3\nOutput: 2 with nums = [2,2]",
			"Input: nums = [0,1,2,2,3,0,4,2], val = 2\nOutput: 5 with nums = [0,1,3,0,4]",
		},
		Hints: []string{
			"The order of elements can be changed. It doesn't matter what you leave beyond the new length.",
			"Two pointers approach can be useful here.",
		},
		Constraints: []string{
			"0 <= nums.length <= 100",
			"0 <= nums[i] <= 50",
			"0 <= val <= 100",
		},
	},
	{
		QuestionNumber:  28,
		Title:           "Implement strStr()",
		DifficultyLevel: "Easy",
		Description:     "Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.",
		Examples: []string{
			"Input: haystack = 'hello', needle = 'll'\nOutput: 2",
			"Input: haystack = 'aaaaa', needle = 'bba'\nOutput: -1",
		},
		Hints: []string{
			"The naive approach is to check each substring of the haystack with the needle.",
			"Use the built-in string library functions in Go to implement this efficiently.",
		},
		Constraints: []string{
			"0 <= haystack.length, needle.length <= 5 * 10^4",
			"haystack and needle consist of only lower-case English characters.",
		},
	},
	{
		QuestionNumber:  29,
		Title:           "Divide Two Integers",
		DifficultyLevel: "Medium",
		Description:     "Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.",
		Examples: []string{
			"Input: dividend = 10, divisor = 3\nOutput: 3",
			"Input: dividend = 7, divisor = -3\nOutput: -2",
		},
		Hints: []string{
			"Think about bit manipulation (shift operations).",
			"Consider edge cases like overflow and handling of negative numbers.",
		},
		Constraints: []string{
			"-2^31 <= dividend, divisor <= 2^31 - 1",
			"divisor != 0",
		},
	},
	{
		QuestionNumber:  30,
		Title:           "Substring with Concatenation of All Words",
		DifficultyLevel: "Hard",
		Description:     "You are given a string s and an array of words words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.",
		Examples: []string{
			"Input: s = 'barfoothefoobarman', words = ['foo','bar']\nOutput: [0,9]",
			"Input: s = 'wordgoodgoodgoodbestword', words = ['word','good','best','word']\nOutput: []",
		},
		Hints: []string{
			"Use a sliding window approach combined with a hash map to efficiently find the substrings.",
			"Consider the length of words and how they can concatenate in any order.",
		},
		Constraints: []string{
			"1 <= s.length <= 10^4",
			"words.length <= 5000",
			"1 <= words[i].length <= 30",
		},
	},
	{
		QuestionNumber:  31,
		Title:           "Next Permutation",
		DifficultyLevel: "Medium",
		Description:     "Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers. If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).",
		Examples: []string{
			"Input: nums = [1,2,3]\nOutput: [1,3,2]",
			"Input: nums = [3,2,1]\nOutput: [1,2,3]",
			"Input: nums = [1,1,5]\nOutput: [1,5,1]",
		},
		Hints: []string{
			"To generate the next permutation, you need to find the first pair of two successive numbers a[i] and a[i-1], from the right, which satisfy a[i] > a[i-1].",
			"Once the successor is found, the next step is to find the smallest number on right side of partion index which is greater than value found in previous step.",
		},
		Constraints: []string{
			"1 <= nums.length <= 100",
			"0 <= nums[i] <= 100",
		},
	},
	{
		QuestionNumber:  32,
		Title:           "Longest Valid Parentheses",
		DifficultyLevel: "Hard",
		Description:     "Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.",
		Examples: []string{
			"Input: s = '(()'\nOutput: 2",
			"Input: s = ')()())'\nOutput: 4",
		},
		Hints: []string{
			"Use a stack to keep track of the indices of '(' characters.",
			"Scan the string from left to right and from right to left to find the longest valid substring.",
		},
		Constraints: []string{
			"0 <= s.length <= 3 * 10^4",
			"s[i] is '(' or ')'.",
		},
	},
	{
		QuestionNumber:  33,
		Title:           "Search in Rotated Sorted Array",
		DifficultyLevel: "Medium",
		Description:     "There is an integer array nums sorted in ascending order (with distinct values). Prior to being rotated at some unknown pivot, nums was originally a sorted ascending array. You are given a target value to search. If found in the array return its index, otherwise return -1.",
		Examples: []string{
			"Input: nums = [4,5,6,7,0,1,2], target = 0\nOutput: 4",
			"Input: nums = [4,5,6,7,0,1,2], target = 3\nOutput: -1",
		},
		Hints: []string{
			"Consider using binary search for an efficient solution.",
			"Think about the conditions for determining whether to search left or right of the mid-point in the array.",
		},
		Constraints: []string{
			"1 <= nums.length <= 5000",
			"-10^4 <= nums[i], target <= 10^4",
			"All values of nums are unique.",
			"nums is guaranteed to be rotated at some pivot.",
		},
	},
	{
		QuestionNumber:  34,
		Title:           "Find First and Last Position of Element in Sorted Array",
		DifficultyLevel: "Medium",
		Description:     "Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value. If the target is not found in the array, return [-1, -1].",
		Examples: []string{
			"Input: nums = [5,7,7,8,8,10], target = 8\nOutput: [3,4]",
			"Input: nums = [5,7,7,8,8,10], target = 6\nOutput: [-1,-1]",
		},
		Hints: []string{
			"Consider using binary search to find the leftmost and rightmost positions of the target.",
			"Implement two separate binary searches for finding the starting and ending positions.",
		},
		Constraints: []string{
			"0 <= nums.length <= 10^5",
			"-10^9 <= nums[i] <= 10^9",
			"nums is a non-decreasing array.",
		},
	},
	{
		QuestionNumber:  35,
		Title:           "Search Insert Position",
		DifficultyLevel: "Easy",
		Description:     "Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.",
		Examples: []string{
			"Input: nums = [1,3,5,6], target = 5\nOutput: 2",
			"Input: nums = [1,3,5,6], target = 2\nOutput: 1",
			"Input: nums = [1,3,5,6], target = 7\nOutput: 4",
			"Input: nums = [1,3,5,6], target = 0\nOutput: 0",
		},
		Hints: []string{
			"Consider using binary search to find the insertion position efficiently.",
			"Handle edge cases such as the target being smaller or larger than any element in the array.",
		},
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"-10^4 <= nums[i] <= 10^4",
			"nums is sorted in ascending order.",
		},
	},
	{
		QuestionNumber:  36,
		Title:           "Valid Sudoku",
		DifficultyLevel: "Medium",
		Description:     "Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:\n\n1. Each row must contain the digits 1-9 without repetition.\n2. Each column must contain the digits 1-9 without repetition.\n3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.",
		Examples: []string{
			"Input: board = [['5','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]\nOutput: true",
			"Input: board = [['8','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]\nOutput: false",
		},
		Hints: []string{
			"Use sets or arrays to track the presence of numbers in rows, columns, and sub-boxes.",
			"Validate each row, column, and sub-box separately using nested loops.",
		},
		Constraints: []string{
			"board.length == 9",
			"board[i].length == 9",
			"board[i][j] is a digit or '.'.",
			"It is guaranteed that the Sudoku board will be valid.",
		},
	},
	{
		QuestionNumber:  37,
		Title:           "Sudoku Solver",
		DifficultyLevel: "Hard",
		Description:     "Write a program to solve a Sudoku puzzle by filling the empty cells.\n\nA sudoku solution must satisfy all of the following rules:\n\n1. Each of the digits 1-9 must occur exactly once in each row.\n2. Each of the digits 1-9 must occur exactly once in each column.\n3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.\n\nThe '.' character indicates empty cells.",
		Examples: []string{
			"Input: board = [['5','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]\nOutput: [['5','3','4','6','7','8','9','1','2'],['6','7','2','1','9','5','3','4','8'],['1','9','8','3','4','2','5','6','7'],['8','5','9','7','6','1','4','2','3'],['4','2','6','8','5','3','7','9','1'],['7','1','3','9','2','4','8','5','6'],['9','6','1','5','3','7','2','8','4'],['2','8','7','4','1','9','6','3','5'],['3','4','5','2','8','6','1','7','9']]",
		},
		Hints: []string{
			"Use backtracking to explore all possibilities.",
			"Implement functions to check the validity of placing a number in a specific cell based on row, column, and 3x3 sub-box constraints.",
		},
		Constraints: []string{
			"board.length == 9",
			"board[i].length == 9",
			"board[i][j] is a digit or '.'.",
			"It is guaranteed that the Sudoku puzzle will have a single unique solution.",
		},
	},
	{
		QuestionNumber:  38,
		Title:           "Count and Say",
		DifficultyLevel: "Easy",
		Description:     "The count-and-say sequence is a sequence of digit strings defined by the recursive formula:\n\n1. countAndSay(1) = '1'\n2. countAndSay(n) is the way you would 'say' the digit string from countAndSay(n-1), which is then converted into a different representation.\n\nTo determine how you 'say' a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.\n\nFor example, the saying and conversion for digit string '3322251' is '2 3 3 2 2 1' ('two 3s, two 2s, then one 1').",
		Examples: []string{
			"Input: n = 1\nOutput: '1'\nExplanation: This is the base case.",
			"Input: n = 4\nOutput: '1211'\nExplanation: countAndSay(1) = '1', countAndSay(2) = '11', countAndSay(3) = '21', countAndSay(4) = '1211'.",
		},
		Hints: []string{
			"Use iterative approach to generate the next sequence based on the previous one.",
			"Consider using two pointers to traverse and count consecutive characters in the string.",
		},
		Constraints: []string{
			"1 <= n <= 30",
		},
	},
	{
		QuestionNumber:  39,
		Title:           "Combination Sum",
		DifficultyLevel: "Medium",
		Description:     "Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.\n\nThe same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.",
		Examples: []string{
			"Input: candidates = [2,3,6,7], target = 7\nOutput: [[2,2,3],[7]]\nExplanation:\n2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times, similarly with 3.\n7 is a candidate, and 7 = 7.\nThese are the only two combinations.",
			"Input: candidates = [2,3,5], target = 8\nOutput: [[2,2,2,2],[2,3,3],[3,5]]\nExplanation:\n2, 3, and 5 are candidates, and 2 + 2 + 2 + 2 = 8, 2 + 3 + 3 = 8, and 3 + 5 = 8 are the only three combinations.",
			"Input: candidates = [2], target = 1\nOutput: []\nExplanation:\n2 is the only candidate, and 2 < 1. There are no combinations.",
			"Input: candidates = [1], target = 1\nOutput: [[1]]",
			"Input: candidates = [1], target = 2\nOutput: [[1,1]]",
		},
		Hints: []string{
			"Use backtracking to explore all possible combinations.",
			"Sort the candidates array to handle duplicates easily and improve efficiency.",
		},
		Constraints: []string{
			"1 <= candidates.length <= 30",
			"1 <= candidates[i] <= 200",
			"All elements of candidates are distinct.",
			"1 <= target <= 500",
		},
	},
	{
		QuestionNumber:  40,
		Title:           "Combination Sum II",
		DifficultyLevel: "Medium",
		Description:     "Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.\n\nEach number in candidates may only be used once in the combination.\n\nNote: The solution set must not contain duplicate combinations.",
		Examples: []string{
			"Input: candidates = [10,1,2,7,6,1,5], target = 8\nOutput: [[1,1,6],[1,2,5],[1,7],[2,6]]\nExplanation:\n1 + 1 + 6 = 8\n1 + 2 + 5 = 8\n1 + 7 = 8\n2 + 6 = 8\nThese are the only unique combinations.",
			"Input: candidates = [2,5,2,1,2], target = 5\nOutput: [[1,2,2],[5]]\nExplanation:\n1 + 2 + 2 = 5\n5 = 5\nThese are the only unique combinations.",
		},
		Hints: []string{
			"Use backtracking to explore all possible combinations.",
			"Sort the candidates array to handle duplicates easily and improve efficiency.",
		},
		Constraints: []string{
			"1 <= candidates.length <= 100",
			"1 <= candidates[i] <= 50",
			"1 <= target <= 30",
		},
	},
	{
		QuestionNumber:  41,
		Title:           "First Missing Positive",
		DifficultyLevel: "Hard",
		Description:     "Given an unsorted integer array nums, return the smallest missing positive integer.\n\nYou must implement an algorithm that runs in O(n) time and uses constant extra space.",
		Examples: []string{
			"Input: nums = [1,2,0]\nOutput: 3",
			"Input: nums = [3,4,-1,1]\nOutput: 2",
			"Input: nums = [7,8,9,11,12]\nOutput: 1",
		},
		Hints: []string{
			"Think about how to make use of the array itself to store information.",
			"Consider using a cyclic sort approach.",
		},
		Constraints: []string{
			"1 <= nums.length <= 5 * 10^5",
			"-2 * 10^9 <= nums[i] <= 2 * 10^9",
		},
	},
	{
		QuestionNumber:  42,
		Title:           "Trapping Rain Water",
		DifficultyLevel: "Hard",
		Description:     "Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.\n\nThe height at each index i represents the elevation of the terrain at that point. The width of each bar is 1.\n\nConstraints:\n\nn == height.length\n0 <= n <= 3 * 10^4\n0 <= height[i] <= 10^5",
		Examples: []string{
			"Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]\nOutput: 6\nExplanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.",
			"Input: height = [4,2,0,3,2,5]\nOutput: 9",
		},
		Hints: []string{
			"Try to solve it with two pointers.",
			"Simulate the process of water flowing from the highest point to both ends.",
		},
		Constraints: []string{
			"1 <= n <= 3 * 10^4",
			"0 <= height[i] <= 10^5",
		},
	},
	{
		QuestionNumber:  43,
		Title:           "Multiply Strings",
		DifficultyLevel: "Medium",
		Description:     "Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.\n\nNote: You must not use any built-in BigInteger library or convert the inputs to integer directly.",
		Examples: []string{
			"Input: num1 = '2', num2 = '3'\nOutput: '6'",
			"Input: num1 = '123', num2 = '456'\nOutput: '56088'",
		},
		Hints: []string{
			"Use an array to store intermediate results of multiplication.",
			"Consider how traditional multiplication works with pen and paper.",
		},
		Constraints: []string{
			"1 <= num1.length, num2.length <= 200",
			"num1 and num2 consist of digits only.",
			"Both num1 and num2 do not contain any leading zero, except the number 0 itself.",
		},
	},
	{
		QuestionNumber:  44,
		Title:           "Wildcard Matching",
		DifficultyLevel: "Hard",
		Description:     "Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.\n\n'?' Matches any single character.\n'*' Matches any sequence of characters (including the empty sequence).\nThe matching should cover the entire input string (not partial).\n\nConstraints:\n\n0 <= s.length, p.length <= 2000\ns contains only lowercase English letters.\np contains only lowercase English letters, '?' or '*'.",
		Examples: []string{
			"Input: s = 'aa', p = 'a'\nOutput: false\nExplanation: 'a' does not match the entire string 'aa'.",
			"Input: s = 'adceb', p = '*a*b'\nOutput: true\nExplanation: The first '*' matches the empty sequence, while the second '*' matches the substring 'dce'.",
			"Input: s = 'acdcb', p = 'a*c?b'\nOutput: false\nExplanation: The matching substring is 'acdcb', which clearly does not match 'a*c?b'.",
		},
		Hints: []string{
			"Dynamic programming approach can be used to solve this problem efficiently.",
			"Consider how '?' and '*' can be handled in the matching process.",
		},
		Constraints: []string{
			"0 <= s.length, p.length <= 2000",
			"s contains only lowercase English letters.",
			"p contains only lowercase English letters, '?' or '*'.",
		},
	},
	{
		QuestionNumber:  45,
		Title:           "Jump Game II",
		DifficultyLevel: "Hard",
		Description:     "Given an array of non-negative integers nums, you are initially positioned at the first index of the array.\n\nEach element in the array represents your maximum jump length at that position.\n\nYour goal is to reach the last index in the minimum number of jumps.\n\nYou can assume that you can always reach the last index.",
		Examples: []string{
			"Input: nums = [2,3,1,1,4]\nOutput: 2\nExplanation: The minimum jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.",
			"Input: nums = [2,3,0,1,4]\nOutput: 2\nExplanation: The minimum jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 4 steps to the last index.",
		},
		Hints: []string{
			"Use a greedy approach to track the furthest point you can reach with the current number of jumps.",
			"Consider how to optimize the jumps needed using the maximum reach strategy.",
		},
		Constraints: []string{
			"1 <= nums.length <= 1000",
			"0 <= nums[i] <= 10^5",
		},
	},
	{
		QuestionNumber:  46,
		Title:           "Permutations",
		DifficultyLevel: "Medium",
		Description:     "Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.",
		Examples: []string{
			"Input: nums = [1,2,3]\nOutput: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]",
			"Input: nums = [0,1]\nOutput: [[0,1],[1,0]]",
			"Input: nums = [1]\nOutput: [[1]]",
		},
		Hints: []string{
			"Backtracking is an efficient way to solve this problem.",
			"Think about how to swap elements to generate different permutations.",
		},
		Constraints: []string{
			"1 <= nums.length <= 6",
			"-10 <= nums[i] <= 10",
			"All the integers of nums are unique.",
		},
	},
	{
		QuestionNumber:  47,
		Title:           "Permutations II",
		DifficultyLevel: "Medium",
		Description:     "Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.",
		Examples: []string{
			"Input: nums = [1,1,2]\nOutput: [[1,1,2],[1,2,1],[2,1,1]]",
			"Input: nums = [1,2,3]\nOutput: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]",
		},
		Hints: []string{
			"Use a similar approach to generate permutations as in the previous problem, but handle duplicates carefully.",
			"Consider how to skip duplicates to generate unique permutations.",
		},
		Constraints: []string{
			"1 <= nums.length <= 8",
			"-10 <= nums[i] <= 10",
		},
	},
	{
		QuestionNumber:  48,
		Title:           "Rotate Image",
		DifficultyLevel: "Medium",
		Description:     "You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).\n\nYou have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.",
		Examples: []string{
			"Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]\nOutput: [[7,4,1],[8,5,2],[9,6,3]]",
			"Input: matrix = [[1]]\nOutput: [[1]]",
			"Input: matrix = [[1,2],[3,4]]\nOutput: [[3,1],[4,2]]",
		},
		Hints: []string{
			"To rotate the matrix in-place, consider how elements move during rotation.",
			"Think about how to transpose the matrix and then reverse each row to achieve the rotation.",
		},
		Constraints: []string{
			"matrix.length == n",
			"matrix[i].length == n",
			"1 <= n <= 20",
			"-1000 <= matrix[i][j] <= 1000",
		},
	},
	{
		QuestionNumber:  49,
		Title:           "Group Anagrams",
		DifficultyLevel: "Medium",
		Description:     "Given an array of strings strs, group the anagrams together. You can return the answer in any order.\n\nAn Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.",
		Examples: []string{
			"Input: strs = ['eat','tea','tan','ate','nat','bat']\nOutput: [['bat'],['nat','tan'],['ate','eat','tea']]",
			"Input: strs = ['']\nOutput: [['']]",
			"Input: strs = ['a']\nOutput: [['a']]",
		},
		Hints: []string{
			"Use a hash map to store groups of anagrams.",
			"Consider how to represent and compare anagrams efficiently.",
		},
		Constraints: []string{
			"1 <= strs.length <= 10^4",
			"0 <= strs[i].length <= 100",
			"strs[i] consists of lower-case English letters.",
		},
	},
	{
		QuestionNumber:  50,
		Title:           "Pow(x, n)",
		DifficultyLevel: "Medium",
		Description:     "Implement `pow(x, n)`, which calculates `x` raised to the power `n` (i.e., `x^n`).",
		Examples: []string{
			"Input: x = 2.00000, n = 10\nOutput: 1024.00000",
			"Input: x = 2.10000, n = 3\nOutput: 9.26100",
			"Input: x = 2.00000, n = -2\nOutput: 0.25000\nExplanation: 2^-2 = 1/2^2 = 1/4 = 0.25",
		},
		Hints: []string{
			"Consider how to optimize the calculation using recursion and divide-and-conquer techniques.",
			"Handle both positive and negative values of `n`.",
		},
		Constraints: []string{
			"-100.0 < x < 100.0",
			"-2^31 <= n <= 2^31-1",
			"-10^4 <= x^n <= 10^4",
		},
	},
	{
		QuestionNumber:  51,
		Title:           "N-Queens",
		DifficultyLevel: "Hard",
		Description:     "The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.\n\nGiven an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.",
		Examples: []string{
			"Input: n = 4\nOutput: [['.Q..','...Q','Q...','..Q.'],['..Q.','Q...','...Q','.Q..']]\nExplanation: There exist two distinct solutions to the 4-queens puzzle as shown above.",
			"Input: n = 1\nOutput: [['Q']]\nExplanation: The 1-queens puzzle has only one solution, which is shown above.",
		},
		Hints: []string{
			"Use backtracking to try placing queens row by row and backtrack when a conflict is found.",
			"Consider how to efficiently check conflicts for each placement of a queen.",
		},
		Constraints: []string{
			"1 <= n <= 9",
		},
	},
	{
		QuestionNumber:  52,
		Title:           "N-Queens II",
		DifficultyLevel: "Hard",
		Description:     "The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.\n\nGiven an integer n, return the number of distinct solutions to the n-queens puzzle.",
		Examples: []string{
			"Input: n = 4\nOutput: 2\nExplanation: There are two distinct solutions to the 4-queens puzzle as shown above.",
			"Input: n = 1\nOutput: 1\nExplanation: The 1-queens puzzle has only one solution.",
		},
		Hints: []string{
			"Use backtracking to try placing queens row by row and count valid solutions.",
			"Consider how to efficiently check conflicts for each placement of a queen.",
		},
		Constraints: []string{
			"1 <= n <= 9",
		},
	},
	{
		QuestionNumber:  53,
		Title:           "Maximum Subarray",
		DifficultyLevel: "Easy",
		Description:     "Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.",
		Examples: []string{
			"Input: nums = [-2,1,-3,4,-1,2,1,-5,4]\nOutput: 6\nExplanation: [4,-1,2,1] has the largest sum = 6.",
			"Input: nums = [1]\nOutput: 1",
			"Input: nums = [5,4,-1,7,8]\nOutput: 23",
		},
		Hints: []string{
			"Consider using Kadane's algorithm to solve this problem efficiently.",
			"Think about how to keep track of the current maximum subarray sum as you iterate through the array.",
		},
		Constraints: []string{
			"1 <= nums.length <= 3 * 10^4",
			"-10^5 <= nums[i] <= 10^5",
		},
	},
	{
		QuestionNumber:  54,
		Title:           "Spiral Matrix",
		DifficultyLevel: "Medium",
		Description:     "Given an m x n matrix, return all elements of the matrix in spiral order.",
		Examples: []string{
			"Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]\nOutput: [1,2,3,6,9,8,7,4,5]",
			"Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\nOutput: [1,2,3,4,8,12,11,10,9,5,6,7]",
		},
		Hints: []string{
			"Consider simulating the process of moving in a spiral order through the matrix.",
			"Think about how to handle boundary conditions and direction changes efficiently.",
		},
		Constraints: []string{
			"m == matrix.length",
			"n == matrix[i].length",
			"1 <= m, n <= 10",
			"-100 <= matrix[i][j] <= 100",
		},
	},
	{
		QuestionNumber:  55,
		Title:           "Jump Game",
		DifficultyLevel: "Medium",
		Description:     "Given an array of non-negative integers nums, you are initially positioned at the first index of the array.\nEach element in the array represents your maximum jump length at that position.\nDetermine if you can reach the last index.",
		Examples: []string{
			"Input: nums = [2,3,1,1,4]\nOutput: true\nExplanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.",
			"Input: nums = [3,2,1,0,4]\nOutput: false\nExplanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.",
		},
		Hints: []string{
			"Use a greedy approach to keep track of the furthest reachable index.",
			"Consider how to iterate through the array and update the maximum reachable index at each step.",
		},
		Constraints: []string{
			"1 <= nums.length <= 10^4",
			"0 <= nums[i] <= 10^5",
		},
	},
	{
		QuestionNumber:  56,
		Title:           "Merge Intervals",
		DifficultyLevel: "Medium",
		Description:     "Given an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.",
		Examples: []string{
			"Input: intervals = [[1,3],[2,6],[8,10],[15,18]]\nOutput: [[1,6],[8,10],[15,18]]\nExplanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].",
			"Input: intervals = [[1,4],[4,5]]\nOutput: [[1,5]]\nExplanation: Intervals [1,4] and [4,5] are considered overlapping.",
		},
		Hints: []string{
			"Sort the intervals based on their start times to simplify the merging process.",
			"Iterate through the sorted intervals and merge overlapping intervals.",
		},
		Constraints: []string{
			"1 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= start_i <= end_i <= 10^4",
		},
	},
	{
		QuestionNumber:  57,
		Title:           "Insert Interval",
		DifficultyLevel: "Hard",
		Description:     "Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).\n\nYou may assume that the intervals were initially sorted according to their start times.",
		Examples: []string{
			"Input: intervals = [[1,3],[6,9]], newInterval = [2,5]\nOutput: [[1,5],[6,9]]\nExplanation: Insert [2,5] into the intervals [1,3] and [6,9] and merge them to [1,5],[6,9].",
			"Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]\nOutput: [[1,2],[3,10],[12,16]]\nExplanation: Insert [4,8] into the intervals [3,5],[6,7],[8,10] and merge them to [3,10].",
		},
		Hints: []string{
			"Sort the intervals based on their start times to simplify the insertion process.",
			"Consider how to handle the merging of intervals after inserting the new interval.",
		},
		Constraints: []string{
			"0 <= intervals.length <= 10^4",
			"intervals[i].length == 2",
			"0 <= intervals[i][0] <= intervals[i][1] <= 10^5",
			"intervals is sorted by intervals[i][0] in ascending order.",
		},
	},
	{
		QuestionNumber:  58,
		Title:           "Length of Last Word",
		DifficultyLevel: "Easy",
		Description:     "Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.\nA word is a maximal substring consisting of non-space characters only.",
		Examples: []string{
			"Input: s = \"Hello World\"\nOutput: 5\nExplanation: The last word is \"World\" with length 5.",
			"Input: s = \"   fly me   to   the moon  \"\nOutput: 4\nExplanation: The last word is \"moon\" with length 4.",
			"Input: s = \"\"\nOutput: 0",
		},
		Hints: []string{
			"Trim the trailing and leading spaces of the string before processing.",
			"Scan the string from the end to find the last word and calculate its length.",
		},
		Constraints: []string{
			"1 <= s.length <= 10^4",
			"s consists of only English letters and spaces ' '.",
			"There will be at least one word in s.",
		},
	},
	{
		QuestionNumber:  59,
		Title:           "Spiral Matrix II",
		DifficultyLevel: "Medium",
		Description:     "Given a positive integer n, generate an n x n matrix filled with elements from 1 to n^2 in spiral order.",
		Examples: []string{
			"Input: n = 3\nOutput: [[1,2,3],[8,9,4],[7,6,5]]",
			"Input: n = 1\nOutput: [[1]]",
		},
		Hints: []string{
			"Simulate the process of filling the matrix in spiral order.",
			"Keep track of the boundaries and direction changes while filling the matrix.",
		},
		Constraints: []string{
			"1 <= n <= 20",
		},
	},
	{
		QuestionNumber:  60,
		Title:           "Permutation Sequence",
		DifficultyLevel: "Medium",
		Description:     "The set [1, 2, 3, ..., n] contains a total of n! unique permutations.\nBy listing and labeling all of the permutations in order, return the kth permutation sequence.",
		Examples: []string{
			"Input: n = 3, k = 3\nOutput: \"213\"\nExplanation: The permutations of [1,2,3] are \"123\", \"132\", \"213\", \"231\", \"312\", and \"321\". The 3rd permutation is \"213\".",
			"Input: n = 4, k = 9\nOutput: \"2314\"",
		},
		Hints: []string{
			"Use a factorial-based approach to determine the digit at each position.",
			"Consider how to efficiently find and construct the kth permutation sequence.",
		},
		Constraints: []string{
			"1 <= n <= 9",
			"1 <= k <= n!",
		},
	},
}

func InsertProblems(db *sql.DB) {
	l := postgres.NewProblemRepo(db)
	for _, p := range problems {
		l.CreateProblem(p)
	}
}
