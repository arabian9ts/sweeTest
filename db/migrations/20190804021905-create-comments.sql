
-- +migrate Up
CREATE TABLE IF NOT EXISTS comments (
    `id`          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `help_id`     INT                NOT NULL,
    `user_type`   VARCHAR(10)        NOT NULL,
    `user_id`     INT                NOT NULL,
    `contents`    TEXT               NOT NULL,
    `created_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX user_idef_index(user_type, user_id),
    INDEX help_id_index(help_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE comments;
