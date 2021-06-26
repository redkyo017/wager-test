DROP DATABASE IF EXISTS wager;
CREATE DATABASE wager;

Use wager;

DROP TABLE IF EXISTS `wagers`;
CREATE TABLE `wagers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `total_wager_value` int(11) DEFAULT NULL,
  `odds` int(11) DEFAULT NULL,
  `selling_price` float DEFAULT NULL,
  `current_selling_price` float DEFAULT NULL,
  `percentage_sold` int(11) DEFAULT NULL,
  `amount_sold` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `selling_percentage` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `buys`;
CREATE TABLE `buys` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `wager_id` bigint(20) unsigned DEFAULT NULL,
  `buying_price` float DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;