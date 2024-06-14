create type enum as difficulty_lavels("Easy", "Medium", "Hard");

create table problems {
	title varchar primary key not null,
	problem_number int increments unique  not null,
	difficulty enum references difficulty_lavels not null,
	description text not null,
	constraints text[] not null,
	hints text[]
}