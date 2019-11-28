
-- +migrate Up
CREATE TABLE IF NOT EXISTS teachers_lectures (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    teacher_id  INT                NOT NULL,
    lecture_id  INT                NOT NULL,
    crated_at   TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    unique student_lecture_index(teacher_id, lecture_id),
    FOREIGN KEY (teacher_id) REFERENCES teachers(id),
    FOREIGN KEY (lecture_id) REFERENCES lectures(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS teachers_lectures;
