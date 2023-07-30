CREATE TABLE uids (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT UNIQUE NOT NULL,
    access ENUM('all', 'no_hidden', 'self') NOT NULL,
    hidden TINYINT(1) NOT NULL DEFAULT 0
);

-- INSERT INTO uids SET user_id = 120, access = 'all', hidden = 1;

CREATE TABLE requests (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    u_id INT NOT NULL,
    hash VARCHAR(256) UNIQUE NOT NULL,
    service_family ENUM('') NOT NULL,
    data JSON DEFAULT NULL,
    FOREIGN KEY(u_id) REFERENCES uids(id)
);
-- INSERT INTO requests SET u_id = 16, hash = 'afwghrheffihiuewhif', service_family = 'email', data = '{"d":"dd","f":{"g":"gg","h":"hh"}}';

CREATE TABLE changes (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    u_id INT NOT NULL,
    field_name VARCHAR(64) NOT NULL,
    service_family ENUM('') NOT NULL,
    old_value TEXT,
    new_value TEXT,
    FOREIGN KEY(u_id) REFERENCES uids(id)
);