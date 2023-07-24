-- Modify "permissions" table
ALTER TABLE `permissions` DROP COLUMN `name`, ADD COLUMN `title` varchar(50) NOT NULL, ADD COLUMN `value` varchar(50) NOT NULL, ADD COLUMN `group` varchar(50) NULL;
-- Modify "roles" table
ALTER TABLE `roles` DROP COLUMN `name`, ADD COLUMN `title` varchar(50) NOT NULL;
-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `has_permission`, ADD INDEX `users_roles_users` (`role_id`), ADD CONSTRAINT `users_roles_users` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
