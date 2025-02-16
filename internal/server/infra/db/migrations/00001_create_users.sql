-- +goose Up
CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		email VARCHAR NOT NULL,
		pass_hash VARCHAR NOT NULL,
		CONSTRAINT "uniq_users__email" UNIQUE (email)
);

-- +goose Down
DROP TABLE users;
