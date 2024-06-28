CREATE TABLE languages (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name varchar NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TYPE difficulty_levels AS ENUM ('Easy', 'Medium', 'Hard');

CREATE TABLE problems (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    title varchar UNIQUE NOT NULL,
    problem_number serial UNIQUE NOT NULL,
    difficulty difficulty_levels NOT NULL,
    description text NOT NULL,
    constraints text[] NOT NULL,
    hints text[],
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    search_vector tsvector
);

UPDATE problems 
SET search_vector = to_tsvector('english', coalesce(title, '') || ' ' || coalesce(problem_number::text, ''));

CREATE INDEX problems_search_vector_idx 
ON problems USING gin(search_vector);

CREATE OR REPLACE FUNCTION update_search_vector() RETURNS TRIGGER AS $$
BEGIN
  NEW.search_vector := to_tsvector('english', coalesce(NEW.title, '') || ' ' || coalesce(NEW.problem_number::text, ''));
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
ON problems FOR EACH ROW EXECUTE FUNCTION update_search_vector();


CREATE TYPE gender AS ENUM ('Male', 'Female');

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    username varchar NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL,
    full_name varchar,
    profile_image bytea,
    gender gender,
    location varchar,
    birthday date,
    summary text,
    website varchar,
    github varchar,
    linkedin varchar,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE examples (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    problem_id uuid REFERENCES problems(id) NOT NULL,
    input text NOT NULL,
    output text NOT NULL,
    explanation text,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE topics (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name varchar UNIQUE NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE topics_problems (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    problem_id uuid REFERENCES problems(id) NOT NULL,
    topic_id uuid REFERENCES topics(id) NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TYPE status AS ENUM (
    'Accepted',
    'Run Time Error',
    'Compile Error',
    'Wrong Answer',
    'Time Limit Exceeded',
    'Memory Limit Exceeded',
    'Output Limit Exceeded'
);

CREATE TABLE submissions (
    id serial PRIMARY KEY NOT NULL,
    user_id uuid REFERENCES users(id) NOT NULL,
    problem_id uuid REFERENCES problems(id) NOT NULL,
    language_id uuid REFERENCES languages(id) NOT NULL,
    code text NOT NULL,
    submission_status status NOT NULL,
    runtime numeric,
    submission_date timestamp DEFAULT now() NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TYPE run_submit AS ENUM ('Run', 'Submit');

CREATE TABLE testcases_non_btree (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    problem_id uuid REFERENCES problems(id) NOT NULL,
    function_name varchar NOT NULL,
    run_or_submit run_submit NOT NULL,
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
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);
