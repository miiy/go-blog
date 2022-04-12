START TRANSACTION;

CREATE TABLE IF NOT EXISTS `attachments`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    bigint unsigned NOT NULL DEFAULT 0,
    `name`       varchar(255)    NOT NULL DEFAULT '',
    `sha1`       varchar(255)    NOT NULL DEFAULT '',
    `size`       varchar(255)    NOT NULL DEFAULT '',
    `addr`       varchar(255)    NOT NULL DEFAULT '',
    `status`     varchar(255)    NOT NULL DEFAULT '',
    `created_at` timestamp       NULL     DEFAULT NULL,
    `updated_at` timestamp       NULL     DEFAULT NULL,
    `deleted_at` timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `attachments_user_id_index` (`user_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;


COMMIT;