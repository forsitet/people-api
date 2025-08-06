-- +goose Up
CREATE INDEX idx_people_surname ON people(surname);

-- +goose Down
DROP INDEX IF EXISTS idx_people_surname; 
