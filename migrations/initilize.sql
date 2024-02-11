BEGIN;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name TEXT,
    surname TEXT,
    created_time TIME,
    edited_time TIME,
    profile_picture TEXT,
    email TEXT,
    nickname TEXT,
    posts INTEGER,
    status TEXT,
    bio TEXT,
    likes INTEGER
);

CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    headline TEXT,
    likes INTEGER,
    author_id INTEGER
);

CREATE TABLE IF NOT EXISTS post_blocks(
    id SERIAL PRIMARY KEY,
    type INTEGER,
    content TEXT,
    style TEXT,
    post_id INTEGER
);


COMMIT;