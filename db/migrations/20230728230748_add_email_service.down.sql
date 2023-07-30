DROP TABLE email_service;
ALTER TABLE requests MODIFY COLUMN service_family ENUM('');
ALTER TABLE changes MODIFY COLUMN service_family ENUM('');