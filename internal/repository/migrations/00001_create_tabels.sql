-- +goose Up
CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100) NOT NULL,
    gender VARCHAR(20) NOT NULL,
    nationality VARCHAR(50) NOT NULL,
    age INT NOT NULL
);

CREATE TABLE IF NOT EXISTS emails (
    id SERIAL PRIMARY KEY,
    person_id INT NOT NULL REFERENCES people(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS friends (
    person_id INT NOT NULL REFERENCES people(id) ON DELETE CASCADE,
    friend_id INT NOT NULL REFERENCES people(id) ON DELETE CASCADE,
    PRIMARY KEY (person_id, friend_id)
);

-- +goose Down
DROP TABLE IF EXISTS friends;
DROP TABLE IF EXISTS emails;
DROP TABLE IF EXISTS people;
