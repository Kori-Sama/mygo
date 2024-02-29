CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    role ENUM('Old', 'Volunteer', 'Admin') DEFAULT 'Old',
    age INT,
    wallet VARCHAR(200),
    passphrase VARCHAR(200)
);