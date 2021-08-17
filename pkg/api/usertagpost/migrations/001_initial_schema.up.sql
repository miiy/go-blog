CREATE TABLE IF NOT EXISTS `user_tags_posts` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `user_id` bigint unsigned NOT NULL DEFAULT 0,
                        `tag_id` bigint unsigned NOT NULL DEFAULT 0,
                        `post_id` bigint unsigned NOT NULL DEFAULT 0,
                        `created_at` timestamp NULL DEFAULT NULL,
                        `updated_at` timestamp NULL DEFAULT NULL,
                        `deleted_at` timestamp NULL DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        INDEX `user_tags_posts_user_id_index` (`user_id`),
                        INDEX `user_tags_posts_post_id_index` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;