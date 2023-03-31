DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
    `id` varchar(36) NOT NULL,
    `parent_id` varchar(36) DEFAULT NULL,
    `title` varchar(75) NOT NULL,
    `meta_title` varchar(100) DEFAULT NULL,
    `slug` varchar(100) NOT NULL,
    `content` text,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `parent_id` (`parent_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `images`;
CREATE TABLE `images` (
    `id` varchar(36) NOT NULL,
    `file_name` varchar(100) NOT NULL,
    `url` varchar(150) NOT NULL,
    `width` int NOT NULL,
    `height` int NOT NULL,
    `extension` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `cloud_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `file_name` (`file_name`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `post_categories`;
CREATE TABLE `post_categories` (
    `category_id` varchar(36) NOT NULL,
    `post_id` varchar(36) NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`category_id`, `post_id`),
    KEY `category_id` (`category_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `post_comments`;
CREATE TABLE `post_comments` (
    `id` varchar(36) NOT NULL,
    `post_id` varchar(36) NOT NULL,
    `parent_id` varchar(36) DEFAULT NULL,
    `title` varchar(100) NOT NULL,
    `content` text,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `published_at` timestamp NULL DEFAULT NULL,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `post_id` (`post_id`) USING BTREE,
    KEY `parent_id` (`parent_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `post_metas`;
CREATE TABLE `post_metas` (
    `id` varchar(36) NOT NULL,
    `post_id` varchar(36) NOT NULL,
    `key` varchar(50) NOT NULL,
    `content` text,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `post_id` (`post_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `post_tags`;
CREATE TABLE `post_tags` (
    `tag_id` varchar(36) NOT NULL,
    `post_id` varchar(36) NOT NULL,
    PRIMARY KEY (`tag_id`, `post_id`),
    KEY `tag_id` (`tag_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
    `id` varchar(36) NOT NULL,
    `author_id` varchar(36) NOT NULL,
    `parent_id` varchar(36) DEFAULT NULL,
    `title` varchar(75) NOT NULL,
    `meta_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `slug` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `sumary` tinytext,
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
    `thumbnail` json DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `published_at` timestamp NULL DEFAULT NULL,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `slug` (`slug`) USING BTREE,
    KEY `author_id` (`author_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
    `id` varchar(36) NOT NULL,
    `title` varchar(75) NOT NULL,
    `slug` varchar(100) NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` varchar(36) NOT NULL,
    `first_name` varchar(50) DEFAULT NULL,
    `middle_name` varchar(50) DEFAULT NULL,
    `last_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `mobile` varchar(15) NOT NULL,
    `email` varchar(50) NOT NULL,
    `password` varchar(50) NOT NULL,
    `salt` varchar(50) DEFAULT NULL,
    `role` enum('user', 'admin', 'writer') NOT NULL DEFAULT 'user',
    `last_login` timestamp NULL DEFAULT NULL,
    `intro` tinytext,
    `profile` text,
    `avatar` json DEFAULT NULL,
    `cover` json DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` int NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`) USING BTREE,
    UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;