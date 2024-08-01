-- +goose Up
CREATE TABLE characters (
    id VARCHAR(255) PRIMARY KEY,
    name TEXT,
    description TEXT,
    currentHP INTEGER,
    maxHP INTEGER,
    currentMP INTEGER,
    maxMP INTEGER
);

-- +goose Down
DROP TABLE characters;