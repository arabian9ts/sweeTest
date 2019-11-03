
-- +migrate Up
CREATE TABLE IF NOT EXISTS assistants_lectures (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    student_id  INT                NOT NULL,
    lecture_id  INT                NOT NULL,
    crated_at   TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    unique student_lecture_index(student_id, lecture_id),
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (lecture_id) REFERENCES lectures(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS assistants_lectures;
