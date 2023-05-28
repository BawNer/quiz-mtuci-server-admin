CREATE TABLE IF NOT EXISTS reviewers(
    id serial primary key,
    user_id serial,
    quiz_id serial,
    answers jsonb,
    closed_at timestamp without time zone default now()
);