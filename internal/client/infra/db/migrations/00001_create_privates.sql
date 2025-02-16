-- +goose Up
CREATE TABLE privates (
		id integer PRIMARY KEY AUTOINCREMENT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		type VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE privates;
