CREATE TABLE email_service (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    u_id INT NOT NULL,
    data JSON DEFAULT NULL,
    FOREIGN KEY(u_id) REFERENCES uids(id)
);
ALTER TABLE requests MODIFY COLUMN service_family ENUM('email');
ALTER TABLE changes MODIFY COLUMN service_family ENUM('email');