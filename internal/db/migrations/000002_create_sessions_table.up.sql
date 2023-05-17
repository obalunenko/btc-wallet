CREATE TABLE sessions (
    id BIGINT AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    token VARCHAR(255),
    created_at DATETIME NOT NULL,
    expires_at DATETIME DEFAULT NULL,

    PRIMARY KEY (id)
);
COMMIT ;