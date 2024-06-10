select 
		p.id, p.question_number, p.title, p.difficulty_level, p.description, p.examples, p.hints, p.constraints
	from 
		topics_problems as tp
	join
		topics as t
	on 
		tp.topic_id = t.id and t.deleted_at is null
	join
		problems as p
	on 
		p.id = tp.problem_id and p.deleted_at is null
	where
		tp.topic_id = 'a41f6449-9eae-4b40-93be-1d4673bfe8d8' and tp.deleted_at is null


select 
		t.id, t.name
	from 
		topics_problems as tp
	join
		topics as t
	on 
		tp.topic_id = t.id and t.deleted_at is null
	join
		problems as p
	on 
		p.id = tp.problem_id and p.deleted_at is null
	where
		tp.problem_id = 'c81c3b88-6937-47cc-9a8f-32f195911209' and tp.deleted_at is null