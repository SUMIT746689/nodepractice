-- Create "companies" table
CREATE TABLE `companies` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(50) NOT NULL, `domain` varchar(50) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `company_id` bigint NOT NULL, ADD INDEX `users_companies_users` (`company_id`), ADD CONSTRAINT `users_companies_users` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
