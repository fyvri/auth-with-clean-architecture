BEGIN;
DROP DATABASE IF EXISTS `server_development_dev`;
CREATE DATABASE `server_development_dev`;
USE `server_development_dev`;

CREATE TABLE `roles` (
  `id` int unsigned PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(20) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp
);

CREATE TABLE `permissions` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(20) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp
);

CREATE TABLE `role_permission` (
  `role_id` int unsigned,
  `permission_id` bigint unsigned
);

CREATE TABLE `users` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255) NOT NULL,
  `username` varchar(100) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp
);

CREATE TABLE `customers` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `avatar` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp
);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
ALTER TABLE `role_permission` ADD FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`);
ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);


LOCK TABLES `roles` WRITE;
INSERT INTO `roles` (`id`, `name`) VALUES (1,'super-admin'), (2,'admin');
UNLOCK TABLES;

COMMIT;