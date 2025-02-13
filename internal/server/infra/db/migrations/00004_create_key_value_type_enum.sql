-- +goose Up
CREATE TYPE "public"."key_value_type_enum" AS ENUM('meta', 'login', 'password', 'binary', 'number', 'date', 'cvv');

-- +goose Down
DROP TYPE "public"."key_value_type_enum";
