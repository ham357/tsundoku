create table IF not exists `books`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(100) NOT NULL UNIQUE,
 `detail`           VARCHAR(10000) NOT NULL,
 `price`              INT NOT NULL,
 `img`	            MEDIUMBLOB NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `users`
(
 `id`               INT(20) AUTO_INCREMENT,
 `uid`              VARCHAR(255) NOT NULL,
 `name`             VARCHAR(255) NOT NULL,
 `email`            VARCHAR(255) NOT NULL,
 `password`         VARCHAR(255) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
