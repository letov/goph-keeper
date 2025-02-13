-- +goose Up
CREATE TYPE "public"."private_type_enum" AS ENUM('login_password', 'binary', 'bank_card');

-- +goose Down
DROP TYPE "public"."private_type_enum";
