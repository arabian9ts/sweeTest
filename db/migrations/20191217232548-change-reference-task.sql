
-- +migrate Up
ALTER TABLE tasks CHANGE `lecture_id` `class_id` int;
ALTER TABLE tasks DROP INDEX id_and_lecture_id_index;
ALTER TABLE tasks ADD INDEX class_id_index (`class_id`);

-- +migrate Down
ALTER TABLE tasks CHANGE `class_id` `lecture_id` int;
ALTER TABLE tasks DROP INDEX class_id_index;
ALTER TABLE tasks ADD INDEX id_and_lecture_id_index (`id`, `lecture_id`);

