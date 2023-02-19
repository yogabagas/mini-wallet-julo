CREATE TABLE `employees` (
    `employee_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `title` varchar(50) NOT NULL,
    `work_phone` varchar(30) NOT NULL,
    `is_deleted` tinyint(1) DEFAULT 0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` bigint NOT NULL DEFAULT 0,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_by` bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (`employee_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;