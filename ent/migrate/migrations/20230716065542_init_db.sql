-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, `username` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
