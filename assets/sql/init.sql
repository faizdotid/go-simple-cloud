CREATE DATABASE IF NOT EXISTS `go_simple_cloud`;

USE `go_simple_cloud`;

-- Table to store file extensions
CREATE TABLE IF NOT EXISTS `file_extensions` (
    `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `ext` CHAR(4) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `ext_unique_key` (`ext`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table to store preview files
CREATE TABLE IF NOT EXISTS `preview_files` (
    `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `ext_id` INT(11) UNSIGNED NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `url_unique_key` (`url`),
    UNIQUE KEY `name_unique_key` (`name`),
    FOREIGN KEY (`ext_id`) REFERENCES `file_extensions`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Main files table
CREATE TABLE IF NOT EXISTS `files` (
    `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255) NOT NULL,
    `filename` VARCHAR(255) NOT NULL,
    `filesize` INT(11) UNSIGNED NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `preview_file_id` INT(11) UNSIGNED DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `expires_at` TIMESTAMP NULL DEFAULT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `url_unique_key` (`url`),
    FOREIGN KEY (`preview_file_id`) REFERENCES `preview_files`(`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Reference icons: https://www.flaticon.com/packs/file-extension-4
-- Insert file extensions
INSERT INTO `file_extensions` (`id`, `ext`) VALUES
(0, 'unknown'),
(1, 'gif'),
(2, 'jpeg'),
(3, 'jpg'),
(4, 'png'),
(5, 'php');

-- Insert preview files
INSERT INTO `preview_files` (`name`, `ext_id`, `url`) VALUES
('Unknown FILES', 0, 'assets/static/unknown.png'),
('PHP', 5, 'assets/static/php.png'),
('GIF', 1, 'assets/static/gif.png'),
('JPEG', 2, 'assets/static/jpeg.png'),
('JPG', 3, 'assets/static/jpg.png'),
('PNG', 4, 'assets/static/png.png');
