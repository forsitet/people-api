-- +goose Up
INSERT INTO people (name, surname, patronymic, gender, nationality, age) VALUES
('Иван', 'Иванов', 'Иванович', 'мужской', 'русский', 25),
('Мария', 'Петрова', 'Сергеевна', 'женский', 'русская', 30),
('Алексей', 'Сидоров', 'Александрович', 'мужской', 'русский', 28),
('Елена', 'Козлова', 'Дмитриевна', 'женский', 'русская', 35),
('Дмитрий', 'Новиков', 'Владимирович', 'мужской', 'русский', 22),
('Анна', 'Морозова', 'Андреевна', 'женский', 'русская', 27),
('Сергей', 'Волков', 'Игоревич', 'мужской', 'русский', 33),
('Ольга', 'Соколова', 'Николаевна', 'женский', 'русская', 29),
('Андрей', 'Лебедев', 'Петрович', 'мужской', 'русский', 31),
('Наталья', 'Крылова', 'Александровна', 'женский', 'русская', 26);

INSERT INTO emails (person_id, email) VALUES
(1, 'ivan.ivanov@example.com'),
(1, 'ivanov.i@company.ru'),
(2, 'maria.petrova@example.com'),
(2, 'petrova.m@company.ru'),
(3, 'alexey.sidorov@example.com'),
(3, 'sidorov.a@company.ru'),
(4, 'elena.kozlova@example.com'),
(4, 'kozlova.e@company.ru'),
(5, 'dmitry.novikov@example.com'),
(5, 'novikov.d@company.ru'),
(6, 'anna.morozova@example.com'),
(6, 'morozova.a@company.ru'),
(7, 'sergey.volkov@example.com'),
(7, 'volkov.s@company.ru'),
(8, 'olga.sokolova@example.com'),
(8, 'sokolova.o@company.ru'),
(9, 'andrey.lebedev@example.com'),
(9, 'lebedev.a@company.ru'),
(10, 'natalya.krylova@example.com'),
(10, 'krylova.n@company.ru');

INSERT INTO friends (person_id, friend_id) VALUES
(1, 2),
(1, 3),
(2, 1),
(2, 4),
(3, 1),
(3, 5),
(4, 2),
(4, 6),
(5, 3),
(5, 7),
(6, 4),
(6, 8),
(7, 5),
(7, 9),
(8, 6),
(8, 10),
(9, 7),
(9, 10),
(10, 8),
(10, 9);

-- +goose Down
DELETE FROM friends;
DELETE FROM emails;
DELETE FROM people; 
