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
    INDEX `user_tags_user_id_index` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `attachments`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`           bigint unsigned NOT NULL DEFAULT 0,
    `name`              varchar(255)    NOT NULL DEFAULT '' comment '书名',
    `sha1`        varchar(255)    NOT NULL DEFAULT '' comment 'hash',
    `size`         varchar(255)    NOT NULL DEFAULT '' comment 'size',
    `addr`              varchar(255)    NOT NULL DEFAULT '' comment 'addr',
    `status`             varchar(255)    NOT NULL DEFAULT '' comment '页数',
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
    INDEX `user_tags_user_id_index` (`user_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;