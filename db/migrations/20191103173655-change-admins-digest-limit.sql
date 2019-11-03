
-- +migrate Up
ALTER TABLE `admins` MODIFY COLUMN `digest` VARCHAR(60) NOT NULL;

-- +migrate Down
ALTER TABLE `admins` MODIFY COLUMN `digest` VARCHAR(50) NOT NULL;
