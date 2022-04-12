START TRANSACTION;

CREATE TABLE `tags`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(255)    NOT NULL DEFAULT '',
    `description` varchar(255)    NOT NULL DEFAULT '',
    `status`      tinyint(1)      NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 disable',
    `created_at`  timestamp       NULL     DEFAULT NULL,
    `updated_at`  timestamp       NULL     DEFAULT NULL,
    `deleted_at`  timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `tags_name_index` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

COMMIT;