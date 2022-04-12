START TRANSACTION;

CREATE TABLE IF NOT EXISTS `users`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `username`          varchar(255)    NOT NULL DEFAULT '',
    `password`          varchar(255)    NOT NULL DEFAULT '',
    `email`             varchar(255)    NOT NULL DEFAULT '',
    `email_verified_at` datetime        NULL     DEFAULT NULL,
    `phone`             varchar(255)    NOT NULL DEFAULT '',
    `status`            tinyint(1)      NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 disable',
    `created_at`        timestamp       NULL     DEFAULT NULL,
    `updated_at`        timestamp       NULL     DEFAULT NULL,
    `deleted_at`        timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_username_unique` (`username`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;


COMMIT;