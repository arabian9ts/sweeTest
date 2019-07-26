
-- +migrate Up
ALTER TABLE `teachers` MODIFY COLUMN `digest` VARCHAR(60) NOT NULL;

-- +migrate Down
ALTER TABLE `teachers` MODIFY COLUMN `digest` VARCHAR(50) NOT NULL;
