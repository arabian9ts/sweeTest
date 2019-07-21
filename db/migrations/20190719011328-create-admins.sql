
-- +migrate Up
CREATE TABLE IF NOT EXISTS admins (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    first_name  VARCHAR(15)        NOT NULL,
    last_name   VARCHAR(15)        NOT NULL,
    email       varchar(100)       NOT NULL,
    digest      VARCHAR(50)        NOT NULL,
    created_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX email_index(email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE admins;
