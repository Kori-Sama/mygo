CREATE DATABASE IF NOT EXISTS mygo;

USE mygo;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    role ENUM('Old', 'Volunteer', 'Admin') DEFAULT 'Old',
    age INT,
    wallet VARCHAR(200),
    passphrase VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    status ENUM('Draft', 'Published', 'Closed') DEFAULT 'Draft',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS history (
    history_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    transaction_id INT,
    action ENUM('Create', 'Edit', 'Delete', 'Respond') NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)
);