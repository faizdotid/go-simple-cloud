CREATE DATABASE IF NOT EXISTS `go_simple_cloud`;

USE `go_simple_cloud`;

CREATE TABLE IF NOT EXISTS `files` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255) NOT NULL,
    `filename` VARCHAR(255) NOT NULL,
    `filesize` INT(11) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `preview_file_id` INT(11) NULL DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expires_at` TIMESTAMP NULL DEFAULT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `url_unique_key` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `preview_files` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `url_unique_key` (`url`),
    UNIQUE KEY `name_unique_key` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `preview_files` (`id`, `name`, `url`) VALUES (0, 'default', 'assets/static/default.png');
INSERT INTO `preview_files` (`name`, `url`) VALUES ('php', 'assets/static/php.png');
INSERT INTO `preview_files` (`name`, `url`) VALUES ('html', 'assets/static/html.png');