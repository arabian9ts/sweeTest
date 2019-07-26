
-- +migrate Up
ALTER TABLE `students` MODIFY COLUMN `digest` VARCHAR(60) NOT NULL;

-- +migrate Down
ALTER TABLE `students` MODIFY COLUMN `digest` VARCHAR(50) NOT NULL;
