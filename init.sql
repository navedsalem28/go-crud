#Use the database
CREATE DATABASE IF NOT EXISTS crud;

USE crud;


CREATE TABLE IF NOT EXISTS `users`
(`id` INT primary key NOT NULL AUTO_INCREMENT,
 `name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(255)NOT NULL,
    `phone` VARCHAR(20)NOT NULL,
    `created_datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (`email`));
