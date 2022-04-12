START TRANSACTION;


CREATE TABLE IF NOT EXISTS `user_posts`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`        bigint unsigned NOT NULL DEFAULT 0,
    `title`          varchar(255)    NOT NULL DEFAULT '',
    `content`        text                     DEFAULT NULL,
    `status`         tinyint(1)      NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 draft, 3 disable',
    `published_time` DATETIME        NULL     DEFAULT NULL,
    `updated_time`   DATETIME        NULL     DEFAULT NULL,
    `sort`           int             NOT NULL DEFAULT 0,
    `created_at`     timestamp       NULL     DEFAULT NULL,
    `updated_at`     timestamp       NULL     DEFAULT NULL,
    `deleted_at`     timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `user_posts_user_id_index` (`user_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `user_tags`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    bigint unsigned NOT NULL DEFAULT 0,
    `name`       varchar(255)    NOT NULL DEFAULT '',
    `created_at` timestamp       NULL     DEFAULT NULL,
    `updated_at` timestamp       NULL     DEFAULT NULL,
    `deleted_at` timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `user_tags_user_id_index` (`user_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `user_tags_posts`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    bigint unsigned NOT NULL DEFAULT 0,
    `tag_id`     bigint unsigned NOT NULL DEFAULT 0,
    `post_id`    bigint unsigned NOT NULL DEFAULT 0,
    `created_at` timestamp       NULL     DEFAULT NULL,
    `updated_at` timestamp       NULL     DEFAULT NULL,
    `deleted_at` timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `user_tags_posts_user_id_index` (`user_id`),
    INDEX `user_tags_posts_post_id_index` (`post_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;


COMMIT;