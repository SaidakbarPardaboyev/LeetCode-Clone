
problems {
	id uuid pk unique
	title varchar unique
	problem_number integer increments unique
	difficulty varchar
	description text
	constraints text
	hints text
}