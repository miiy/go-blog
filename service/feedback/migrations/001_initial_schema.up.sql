START TRANSACTION;

CREATE TABLE IF NOT EXISTS `feedbacks`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    bigint unsigned NOT NULL DEFAULT 0,
    `content`    varchar(255)    NOT NULL DEFAULT '',
    `created_at` timestamp       NULL     DEFAULT NULL,
    `updated_at` timestamp       NULL     DEFAULT NULL,
    `deleted_at` timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

COMMIT;