CREATE TABLE reactions (
    id VARCHAR(36) PRIMARY KEY AUTO_INCREMENT,
    user_id  VARCHAR(36) NOT NULL,
    target_user_id VARCHAR(36) NOT NULL,
    reaction_type INT(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL
);
