-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
  id       int NOT NULL,
  username varchar(45),
  email    varchar(45),
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
