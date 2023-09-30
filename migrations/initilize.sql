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
)