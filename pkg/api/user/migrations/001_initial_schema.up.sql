CREATE TABLE IF NOT EXISTS `users` (
                                       `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                       `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                       `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                       `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                       `phone` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                       `status`  tinyint(1) NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 disable',
                                       `created_at` timestamp NULL DEFAULT NULL,
                                       `updated_at` timestamp NULL DEFAULT NULL,
                                       `deleted_at` timestamp NULL DEFAULT NULL,
                                       PRIMARY KEY (`id`),
                                       UNIQUE KEY `users_username_unique` (`username`),
                                       UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
