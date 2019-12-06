
-- +migrate Up
ALTER TABLE helps ADD `title` TEXT NOT NULL AFTER `student_id`;

-- +migrate Down
ALTER TABLE helps DROP `title`;

