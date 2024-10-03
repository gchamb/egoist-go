CREATE TABLE user (
    id VARCHAR(36) PRIMARY KEY,
    apple_id VARCHAR(100) UNIQUE DEFAULT NULL,
    email VARCHAR(100) UNIQUE DEFAULT NULL,
    password VARCHAR(255) DEFAULT NULL,
    current_weight FLOAT DEFAULT NULL,
    goal_weight FLOAT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE progress_entry (
    id VARCHAR(36) PRIMARY KEY,
    blob_key VARCHAR(255) UNIQUE NOT NULL,
    current_weight FLOAT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);
ALTER TABLE progress_entry 
ADD UNIQUE INDEX (user_id, created_at);

CREATE TABLE progress_report (
    id VARCHAR(36) PRIMARY KEY,
    current_weight FLOAT NOT NULL,
    last_week_weight FLOAT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE progress_video (
    id VARCHAR(36) PRIMARY KEY,
    blob_key VARCHAR(255) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    frequency VARCHAR(10) NOT NULL,
    created_at DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE revenue_cat_subscriber (
    id VARCHAR(255) PRIMARY KEY,
    transaction_id VARCHAR(255) UNIQUE,
    product_id VARCHAR(30) NOT NULL,
    purchased_at_ms BIGINT NOT NULL,
    expiration_at_ms BIGINT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
)