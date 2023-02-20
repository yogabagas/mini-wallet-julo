ALTER TABLE wallet_balance_histories ADD INDEX idx_wallet_id (wallet_id);

ALTER TABLE wallet_balance_histories ADD INDEX idx_reference_id (reference_id);

SET GLOBAL innodb_lock_wait_timeout = 120;