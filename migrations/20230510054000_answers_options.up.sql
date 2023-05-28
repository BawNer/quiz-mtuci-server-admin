CREATE TABLE IF NOT EXISTS answers_options(
    id serial primary key,
    question_id serial,
    label character varying(128),
    description character varying(128)
);