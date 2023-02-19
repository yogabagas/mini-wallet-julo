CREATE TABLE `shipping_methods` (
    `shipping_method_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `shipping_method` varchar(20) NOT NULL,
    `is_deleted` tinyint(1) DEFAULT 0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` bigint NOT NULL DEFAULT 0,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (`shipping_method_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;