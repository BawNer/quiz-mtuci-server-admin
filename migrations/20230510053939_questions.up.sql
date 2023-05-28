CREATE TABLE IF NOT EXISTS questions(
    id serial primary key,
    quiz_id serial,
    label character varying(128),
    description character varying(128)
);