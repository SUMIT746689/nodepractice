-- Modify "companies" table
ALTER TABLE `companies` MODIFY COLUMN `domain` varchar(50) NOT NULL, ADD COLUMN `create_time` timestamp NOT NULL, ADD COLUMN `update_time` timestamp NOT NULL;
