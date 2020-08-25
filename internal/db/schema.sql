CREATE TABLE users (
       id BIGINT NOT NULL AUTO_INCREMENT,
       created_at DATETIME,

       PRIMARY KEY (id)
       );

CREATE TABLE sessions (
    id BIGINT AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    token VARCHAR(255),
    created_at DATETIME NOT NULL,
    expires_at DATETIME DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE wallets (
    id BIGINT AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    address VARCHAR(256) UNIQUE NOT NULL,

    PRIMARY KEY (id)
);