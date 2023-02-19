CREATE TABLE wallet_balance_histories (
  `wallet_id` CHAR(36) NOT NULL,
  `reference_id` CHAR(36) NOT NULL,
  `amount` DECIMAL(15, 2) NOT NULL,
  `type` TINYINT(1) NOT NULL DEFAULT 1,
  `description` TEXT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` CHAR(36) NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` CHAR(36) NOT NULL,
  PRIMARY KEY (wallet_id, reference_id),
  CONSTRAINT fk_wallet_id FOREIGN KEY (wallet_id) REFERENCES wallets (id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
