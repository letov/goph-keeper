-- +goose Up
CREATE TABLE privates (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		owner INTEGER NOT NULL,
		type private_type_enum NOT NULL,
		CONSTRAINT "FK_login_passwords__owner___users__id" FOREIGN KEY ("owner") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE NO ACTION
);

-- +goose Down
DROP TABLE privates;
