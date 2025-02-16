-- +goose Up
CREATE TABLE key_values (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		private_id INTEGER NOT NULL,
		value BYTEA DEFAULT NULL,
		type VARCHAR NOT NULL,
		CONSTRAINT "FK_key_values__private_id___privates__id" FOREIGN KEY ("private_id") REFERENCES "privates"("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- +goose Down
DROP TABLE key_values;
