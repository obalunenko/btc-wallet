CREATE TABLE ledgers (
  id BIGINT AUTO_INCREMENT ,
  wallet_id BIGINT NOT NULL ,
  balance BIGINT UNSIGNED,
  available_balance BIGINT UNSIGNED ,
  created_at DATETIME ,
  updated_at DATETIME ,


  PRIMARY KEY (id)
);
COMMIT ;