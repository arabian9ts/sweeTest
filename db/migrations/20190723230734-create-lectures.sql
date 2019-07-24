
-- +migrate Up
CREATE TABLE IF NOT EXISTS lectures (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name        VARCHAR(50)        NOT NULL,
    created_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS  lectures;
