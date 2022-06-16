DROP TABLE IF EXISTS users;

CREATE SEQUENCE users_id_seq;

CREATE TABLE users(
    id INTEGER PRIMARY KEY NOT NULL DEFAULT nextval('users_id_seq'),
    first_name VARCHAR(75) NOT NULL,
    last_name VARCHAR(75) NOT NULL,
    birthday VARCHAR(75) NOT NULL,
    username VARCHAR(75) NOT NULL,
    password VARCHAR(75) NOT NULL,
    email VARCHAR(75) NOT NULL,
    city VARCHAR(75) NOT NULL,
    code_zip VARCHAR(75) NOT NULL,
    state VARCHAR(75) NOT NULL

);

ALTER SEQUENCE users_id_seq
OWNED BY users.id;