
-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    `id`          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `lecture_id`  INT                NOT NULL,
    `title`       VARCHAR(100)       NOT NULL,
    `desc`        TEXT,
    `deadline`    TIMESTAMP,
    `created_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX id_and_lecture_id_index (`id`, `lecture_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS tasks;
