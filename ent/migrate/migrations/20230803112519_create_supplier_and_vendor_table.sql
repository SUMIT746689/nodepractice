-- Create "suppliers" table
CREATE TABLE `suppliers` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `address` varchar(255) NOT NULL, `email` varchar(255) NULL, `representative` json NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "vendors" table
CREATE TABLE `vendors` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `address` varchar(255) NOT NULL, `email` varchar(255) NULL, `representative` json NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
