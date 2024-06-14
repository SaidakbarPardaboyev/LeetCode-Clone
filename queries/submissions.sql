create type status as enum(
  'Passed',
  'Run time Error',
  'Compile Error',
  'Wrong Answer',
  'Time Limit Exceeded',
  'Memory Limit Exceeded',
  'Output Limit Exceeded'
);

create table submissions (
	id uuid primary key unique,
	user_username varchar references users(username),
	problem_title varchar references problems(title),
	language_id uuid references language(id),
	code text,
	submission_status status,
	runtime numeric,
	submission_date timestamp
)