
-- +migrate Up
CREATE TABLE IF NOT EXISTS submission_texts (
    `id`                INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `submission_id`     INT                NOT NULL,
    `file_name`         VARCHAR(50)        NOT NULL,
    `contents`          TEXT               NOT NULL,
    `created_at`        TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX submission_id_index(`submission_id`),
    FOREIGN KEY (`submission_id`) REFERENCES submissions(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE submission_texts
