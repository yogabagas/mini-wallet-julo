CREATE TABLE `order_details` (
    `order_detail_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id` bigint unsigned NOT NULL,
    `product_id` bigint unsigned NOT NULL,
    `qty` smallint(2) NOT NULL DEFAULT 1,
    `unit_price` decimal(15, 2) NOT NULL,
    `discount` decimal(3,0) DEFAULT 0.0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` bigint NOT NULL DEFAULT 0,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (`order_detail_id`),
    CONSTRAINT `order_details_product_id_FK` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`),
    CONSTRAINT `order_details_order_id_FK` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;