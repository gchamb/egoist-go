CREATE TABLE user (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) DEFAULT NULL,
    current_weight FLOAT DEFAULT NULL,
    goal_weight FLOAT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE progress_entry (
    id VARCHAR(36) PRIMARY KEY,
    azure_blob_key VARCHAR(255) UNIQUE NOT NULL,
    current_weight FLOAT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE progress_report (
    id VARCHAR(36) PRIMARY KEY,
    current_weight FLOAT NOT NULL,
    last_week_weight FLOAT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE progress_video (
    id VARCHAR(36) PRIMARY KEY,
    azure_blob_key VARCHAR(255) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);