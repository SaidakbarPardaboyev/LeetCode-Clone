create type gender as enum('Male', 'Female');

create table users (
	id 				uuid primary key gen_random_uuid()
	username 		varchar not null,
	full_name 		varchar not null,
    email 			varchar not null,
	password 		varchar not null,
	profile_image 	bytea,
	gender 			gender,
	location 		varchar,
	birthday 		date,
	summary 		text,
	website 		varchar,
	github 			varchar,
	linkedin 		varchar,
    created_at 		timestamp default now() not null,
    updated_at 		timestamp,
    deleted_at 		timestamp
);

-- Insert mock data into the users table
INSERT INTO users (username, full_name, email, password, profile_image, gender, location, birthday, summary, website, github, linkedin, created_at, updated_at, deleted_at)
VALUES
('jdoe', 'John Doe', 'jdoe@example.com', 'password123', NULL, 'Male', 'New York', '1985-07-15', 'Software developer with 10 years of experience.', 'https://johndoe.com', 'https://github.com/jdoe', 'https://linkedin.com/in/jdoe', NOW(), NULL, NULL),
('asmith', 'Alice Smith', 'asmith@example.com', 'password123', NULL, 'Female', 'Los Angeles', '1990-08-20', 'Project manager at a tech company.', 'https://alicesmith.com', 'https://github.com/asmith', 'https://linkedin.com/in/asmith', NOW(), NULL, NULL),
('bwayne', 'Bruce Wayne', 'bwayne@example.com', 'password123', NULL, 'Male', 'Gotham', '1972-02-19', 'Philanthropist and CEO of Wayne Enterprises.', 'https://brucewayne.com', 'https://github.com/bwayne', 'https://linkedin.com/in/bwayne', NOW(), NULL, NULL),
('clarkk', 'Clark Kent', 'clarkk@example.com', 'password123', NULL, 'Male', 'Metropolis', '1980-06-18', 'Journalist at the Daily Planet.', 'https://clarkkent.com', 'https://github.com/clarkk', 'https://linkedin.com/in/clarkk', NOW(), NULL, NULL),
('dprince', 'Diana Prince', 'dprince@example.com', 'password123', NULL, 'Female', 'Themyscira', '1985-03-22', 'Ambassador and warrior.', 'https://dianaprince.com', 'https://github.com/dprince', 'https://linkedin.com/in/dprince', NOW(), NULL, NULL),
('eallen', 'Barry Allen', 'eallen@example.com', 'password123', NULL, 'Male', 'Central City', '1992-09-25', 'Forensic scientist at CCPD.', 'https://barryallen.com', 'https://github.com/eallen', 'https://linkedin.com/in/eallen', NOW(), NULL, NULL),
('pparker', 'Peter Parker', 'pparker@example.com', 'password123', NULL, 'Male', 'New York', '1995-08-10', 'Photographer at the Daily Bugle.', 'https://peterparker.com', 'https://github.com/pparker', 'https://linkedin.com/in/pparker', NOW(), NULL, NULL),
('stark', 'Tony Stark', 'stark@example.com', 'password123', NULL, 'Male', 'Los Angeles', '1970-05-29', 'CEO of Stark Industries.', 'https://tonystark.com', 'https://github.com/stark', 'https://linkedin.com/in/stark', NOW(), NULL, NULL),
('rrogers', 'Steve Rogers', 'rrogers@example.com', 'password123', NULL, 'Male', 'Brooklyn', '1918-07-04', 'Soldier and strategist.', 'https://steverogers.com', 'https://github.com/rrogers', 'https://linkedin.com/in/rrogers', NOW(), NULL, NULL),
('nromanoff', 'Natasha Romanoff', 'nromanoff@example.com', 'password123', NULL, 'Female', 'Stalingrad', '1984-12-03', 'Spy and assassin.', 'https://natasharomanoff.com', 'https://github.com/nromanoff', 'https://linkedin.com/in/nromanoff', NOW(), NULL, NULL);
