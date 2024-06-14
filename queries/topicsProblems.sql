topics {
	id uuid pk unique
	name varchar
}

topics_problems {
	id uuid pk unique
	problem_id uuid *>* problems.id
	topic_id uuid *>* topics.id
}