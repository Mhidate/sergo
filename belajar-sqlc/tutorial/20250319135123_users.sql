-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
