-- Modify "roles" table
ALTER TABLE `roles` ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL, ADD COLUMN `value` varchar(50) NOT NULL;
