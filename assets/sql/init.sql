CREATE DATABASE IF NOT EXISTS `go_simple_cloud`;

USE `go_simple_cloud`;

CREATE TABLE IF NOT EXISTS `files` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255) NOT NULL,
    `filename` VARCHAR(255) NOT NULL,
    `filesize` INT(11) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expires_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `url_unique_key` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


