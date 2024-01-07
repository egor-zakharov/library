CREATE DATABASE borrow_db DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;

USE borrow_db;
CREATE TABLE borrows (
  book_id int NOT NULL, 
  user_id int NOT NULL,
 CONSTRAINT un UNIQUE (book_id, user_id)
);