CREATE TABLE `products` (
    `product_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `product_name` varchar(50) NOT NULL,
    `unit_price` decimal(15,2) NOT NULL,
    `in_stock` char(1) NOT NULL DEFAULT 0,
    `is_deleted` tinyint(1) DEFAULT 0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` bigint NOT NULL DEFAULT 0,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;