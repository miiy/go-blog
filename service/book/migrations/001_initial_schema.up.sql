START TRANSACTION;

CREATE TABLE IF NOT EXISTS `books`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`           bigint unsigned NOT NULL DEFAULT 0,
    `category_id`       bigint unsigned NOT NULL DEFAULT 0 comment '分类',
    `name`              varchar(255)    NOT NULL DEFAULT '' comment '书名',
    `publisher`         varchar(255)    NOT NULL DEFAULT '' comment '出版社',
    `year`              int(10)         NOT NULL DEFAULT 0  comment '出版年',
    `pages`             int(10)         NOT NULL DEFAULT 0  comment '页数',
    `price`             float           NOT NULL DEFAULT 0  comment '定价',
    `binding`           varchar(255)    NOT NULL DEFAULT '' comment '装帧',
    `series`            varchar(255)    NOT NULL DEFAULT '' comment '丛书',
    `isbn`              varchar(255)    NOT NULL DEFAULT '' comment 'ISBN',
    `book_description`  varchar(255)    NOT NULL DEFAULT '' comment '图书简介',
    `about_the_author`  varchar(255)    NOT NULL DEFAULT '' comment '作者简介',
    `table_of_contents` varchar(255)    NOT NULL DEFAULT '' comment '目录',
    `content`           varchar(255)    NOT NULL DEFAULT '' comment '内容',
    `status`            tinyint(1)      NOT NULL DEFAULT 0 COMMENT 'status:0 default, 1 active, 2 disable',
    `created_at`        timestamp       NULL     DEFAULT NULL,
    `updated_at`        timestamp       NULL     DEFAULT NULL,
    `deleted_at`        timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `books_user_id_index` (`user_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `book_seo`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `book_id`           bigint unsigned NOT NULL DEFAULT 0,
    `meta_title`        varchar(255)    NOT NULL DEFAULT '',
    `meta_description`  varchar(255)    NOT NULL DEFAULT '',
    `created_at`        timestamp       NULL     DEFAULT NULL,
    `updated_at`        timestamp       NULL     DEFAULT NULL,
    `deleted_at`        timestamp       NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `books_book_id_index` (`book_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

COMMIT;