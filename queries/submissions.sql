create type status as enum(
  'Accepted',
  'Run Time Error',
  'Compile Error',
  'Wrong Answer',
  'Time Limit Exceeded',
  'Memory Limit Exceeded',
  'Output Limit Exceeded'
);

create table submissions (
	id serial primary key unique not null,
	user_id uuid references users(id) not null,
	problem_id uuid references problems(id) not null,
	language_id uuid references languages(id) not null,
	code text not null,
	submission_status status not null,
	runtime numeric,
	submission_date timestamp default now() not null,
  created_at timestamp default now() not null,
  updated_at timestamp,
  deleted_at timestamp
);
