CREATE DATABASE users_db DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;

USE users_db;
CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  first_name text NOT NULL,
  last_name text NOT NULL,
  PRIMARY KEY (id)
);

insert into users (first_name, last_name) values ("It's", "Me");