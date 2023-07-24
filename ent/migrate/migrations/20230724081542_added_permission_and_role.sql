-- Create "permissions" table
CREATE TABLE `permissions` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(50) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "roles" table
CREATE TABLE `roles` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(50) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "role_permissions" table
CREATE TABLE `role_permissions` (`role_id` bigint NOT NULL, `permission_id` bigint NOT NULL, PRIMARY KEY (`role_id`, `permission_id`), INDEX `role_permissions_permission_id` (`permission_id`), CONSTRAINT `role_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `role_permissions_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "users" table
ALTER TABLE `users` DROP COLUMN `role`, ADD COLUMN `role_id` bigint NOT NULL, ADD COLUMN `has_permission` enum('NULL','ROLE','USER') NOT NULL DEFAULT "NULL";
-- Create "user_permissions" table
CREATE TABLE `user_permissions` (`user_id` bigint NOT NULL, `permission_id` bigint NOT NULL, PRIMARY KEY (`user_id`, `permission_id`), INDEX `user_permissions_permission_id` (`permission_id`), CONSTRAINT `user_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `user_permissions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
