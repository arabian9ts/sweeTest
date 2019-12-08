
-- +migrate Up
CREATE TABLE IF NOT EXISTS classes (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    lecture_id  INT                NOT NULL,
    name        VARCHAR(50)        NOT NULL,
    created_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (lecture_id) REFERENCES lectures(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS classes;

