CREATE TABLE classes (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,   -- "11-A"
    grade INT NOT NULL    -- 11
);
CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    class_id BIGINT NOT NULL REFERENCES classes(id) ON DELETE RESTRICT,
    points TEXT,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    stage TEXT,
    created_at TIMESTAMP DEFAULT now()
);