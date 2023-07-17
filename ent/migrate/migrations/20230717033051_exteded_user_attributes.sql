-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `username` varchar(50) NOT NULL, ADD COLUMN `first_name` varchar(50) NOT NULL, ADD COLUMN `last_name` varchar(50) NOT NULL, ADD COLUMN `phone_number` varchar(255) NULL, ADD COLUMN `email` varchar(255) NULL, ADD COLUMN `role` enum('SUPERADMIN','ADMIN','CASHIER','CUSTOMER') NOT NULL;
