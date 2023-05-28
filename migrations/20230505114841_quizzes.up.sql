CREATE TABLE IF NOT EXISTS quizzes(
  id serial PRIMARY KEY,
  author_id serial,
  quiz_hash character varying(64),
  title character varying(128),
  type character varying(64),
  active bool default false,
  created_at timestamp without time zone default now(),
  updated_at timestamp without time zone default now()
);