
-- +migrate Up
CREATE TABLE IF NOT EXISTS helps (
    `id`          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `lecture_id`  INT                NOT NULL,
    `student_id`  INT                NOT NULL,
    `contents`    TEXT               NOT NULL,
    `created_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX student_id_index(student_id),
    INDEX lecture_id_index(lecture_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE helps;
