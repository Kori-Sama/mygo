CREATE DATABASE IF NOT EXISTS mygo;

USE mygo;

CREATE TABLE IF NOT EXISTS user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    role ENUM('Old', 'Volunteer', 'Admin') DEFAULT 'Old',
    age INT,
    wallet VARCHAR(200),
    passphrase VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction (
    transaction_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    status ENUM('Draft', 'Censoring', 'Passed','Rejected') DEFAULT 'Draft',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS history (
    history_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    transaction_id INT,
    new_value TEXT,
    action ENUM('Create', 'Edit', 'Delete', 'Respond') NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (transaction_id) REFERENCES transaction(transaction_id)
);