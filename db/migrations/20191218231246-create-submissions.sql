
-- +migrate Up
CREATE TABLE IF NOT EXISTS submissions (
    `id`          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `task_id`     INT                NOT NULL,
    `student_id`  INT                NOT NULL,
    `created_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX task_id_index(`task_id`),
    INDEX student_id_index(`student_id`),
    FOREIGN KEY (`task_id`) REFERENCES tasks(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE submissions
