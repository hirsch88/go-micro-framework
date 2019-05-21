-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users`
(
  `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp        NULL                   DEFAULT NULL,
  `updated_at` timestamp        NULL                   DEFAULT NULL,
  `deleted_at` timestamp        NULL                   DEFAULT NULL,
  `username`   varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email`      varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
