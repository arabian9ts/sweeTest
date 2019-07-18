
-- +migrate Up
CREATE TABLE IF NOT EXISTS tas (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    student_no  VARCHAR(20)        NOT NULL,
    first_name  VARCHAR(15)        NOT NULL,
    last_name   VARCHAR(15)        NOT NULL,
    email       varchar(100)       NOT NULL,
    digest      VARCHAR(50)        NOT NULL,
    created_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX student_no_index(student_no)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE tas;
