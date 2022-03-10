START TRANSACTION;

CREATE TABLE IF NOT EXISTS `articles`
(
    `id`               bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`          bigint unsigned NOT NULL DEFAULT 0,
    `title`            varchar(255)    NOT NULL DEFAULT '',
    `meta_title`       varchar(255)    NOT NULL DEFAULT '',
    `meta_description` varchar(255)    NOT NULL DEFAULT '',
    `published_time`   timestamp       NULL     DEFAULT NULL,
    `updated_time`     timestamp       NULL     DEFAULT NULL,
    `from_text`        varchar(255)    NOT NULL DEFAULT '',
    `from_url`         varchar(255)    NOT NULL DEFAULT '',
    `summary`          varchar(255)    NOT NULL DEFAULT '',
    `content`          text                     DEFAULT NULL,
    `status`           tinyint(1)      NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 disable',
    `created_at`       timestamp       NULL     DEFAULT NULL,
    `updated_at`       timestamp       NULL     DEFAULT NULL,
    `deleted_at`       timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `articles_user_id_index` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `books`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`           bigint unsigned NOT NULL DEFAULT 0,
    `name`              varchar(255)    NOT NULL DEFAULT '' comment '书名',
    `categories`        varchar(255)    NOT NULL DEFAULT '' comment '分类',
    `publisher`         varchar(255)    NOT NULL DEFAULT '' comment '出版社',
    `year`              varchar(255)    NOT NULL DEFAULT '' comment '出版年',
    `pages`             varchar(255)    NOT NULL DEFAULT '' comment '页数',
    `price`             varchar(255)    NOT NULL DEFAULT '' comment '定价',
    `binding`           varchar(255)    NOT NULL DEFAULT '' comment '装帧',
    `series`            varchar(255)    NOT NULL DEFAULT '' comment '丛书',
    `isbn`              varchar(255)    NOT NULL DEFAULT '' comment 'ISBN',
    `book_description`  varchar(255)    NOT NULL DEFAULT '' comment '图书简介',
    `about the author`  varchar(255)    NOT NULL DEFAULT '' comment '作者简介',
    `table_of_contents` varchar(255)    NOT NULL DEFAULT '' comment '目录',
    `content`           varchar(255)    NOT NULL DEFAULT '' comment '内容',
    `created_at`        timestamp       NULL     DEFAULT NULL,
    `updated_at`        timestamp       NULL     DEFAULT NULL,
    `deleted_at`        timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `books_user_id_index` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

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

CREATE TABLE IF NOT EXISTS `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `username`   varchar(255)    NOT NULL DEFAULT '',
    `password`   varchar(255)    NOT NULL DEFAULT '',
    `email`      varchar(255)    NOT NULL DEFAULT '',
    `phone`      varchar(255)    NOT NULL DEFAULT '',
    `status`     tinyint(1)      NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 disable',
    `created_at` timestamp       NULL     DEFAULT NULL,
    `updated_at` timestamp       NULL     DEFAULT NULL,
    `deleted_at` timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_username_unique` (`username`),
    UNIQUE KEY `users_email_unique` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

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
