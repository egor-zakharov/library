CREATE USER 'dev'@'%' IDENTIFIED BY 'dev';

GRANT ALL PRIVILEGES ON books_db.* TO 'dev'@'%';
GRANT ALL PRIVILEGES ON users_db.* TO 'dev'@'%';
GRANT ALL PRIVILEGES ON borrow_db.* TO 'dev'@'%';