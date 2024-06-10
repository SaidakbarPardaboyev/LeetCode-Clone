package generator

import (
	"database/sql"
	"leetcode/model"
	"leetcode/storage/postgres"
)

var topics = []string{
	"Array",
	"String",
	"Linked List",
	"Stack",
	"Queue",
	"Tree",
	"Binary Search",
	"Heap",
	"Hash Table",
	"Two Pointers",
	"Depth-First Search (DFS)",
	"Breadth-First Search (BFS)",
	"Backtracking",
	"Dynamic Programming",
	"Greedy",
	"Design",
	"Math",
	"Bit Manipulation",
	"Sort",
	"Recursion",
	"Memoization",
	"Segment Tree",
	"Union Find",
	"Trie",
	"Sliding Window",
	"Binary Indexed Tree (BIT)",
	"Topological Sort",
	"Minimum Spanning Tree (MST)",
	"Suffix Array",
	"Geometry",
	"Simulation",
	"Probability",
	"Concurrency",
	"Database",
}

func InsertTopics(db *sql.DB) {
	l := postgres.NewTopicRepo(db)
	for _, tp := range topics {
		lm := model.Topic{Name: tp}
		l.CreateTopic(&lm)
	}
}
