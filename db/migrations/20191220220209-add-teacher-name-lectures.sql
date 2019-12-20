
-- +migrate Up
ALTER TABLE lectures ADD `teacher_name` VARCHAR(30) NOT NULL AFTER `name`;

-- +migrate Down
ALTER TABLE lectures DROP `teacher_name`;

