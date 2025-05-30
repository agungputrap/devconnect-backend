CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       TEXT        NOT NULL,
    email      TEXT UNIQUE NOT NULL,
    password   TEXT        NOT NULL,
    bio        TEXT,
    skills     TEXT,
    github_url TEXT,
    role       TEXT      DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);