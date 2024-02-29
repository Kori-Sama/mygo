CREATE TABLE history (
    history_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    transaction_id INT,
    action ENUM('Create', 'Edit', 'Delete', 'Respond') NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)
);
