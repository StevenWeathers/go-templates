-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA example;
CREATE TABLE example.authors (
    id   BIGSERIAL PRIMARY KEY,
    name text      NOT NULL,
    bio  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE example.authors;
DROP SCHEMA example;
-- +goose StatementEnd
