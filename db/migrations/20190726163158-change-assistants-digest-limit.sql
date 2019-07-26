
-- +migrate Up
ALTER TABLE `assistants` MODIFY COLUMN `digest` VARCHAR(60) NOT NULL;

-- +migrate Down
ALTER TABLE `assistants` MODIFY COLUMN `digest` VARCHAR(50) NOT NULL;
