
-- +migrate Up
ALTER TABLE `tas` RENAME TO `assistants`;

-- +migrate Down
ALTER TABLE `assistants` RENAME TO `tas`;
