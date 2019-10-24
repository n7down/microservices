
/* users */

CREATE DATABASE `users`;
USE `users`;

CREATE TABLE `users`.`userinfo` (`id` VARCHAR(36) NOT NULL, `username` VARCHAR(100) NOT NULL, `password` VARCHAR(100) NOT NULL, `firstname` VARCHAR(100) NOT NULL, `lastname` VARCHAR(100) NOT NULL, `is_active` BOOLEAN NOT NULL DEFAULT 1, PRIMARY KEY(`id`));

CREATE UNIQUE INDEX `username_unique_index` ON `users`.`userinfo` (`username`);

/* products */

CREATE DATABASE `products`;
USE `products`;

CREATE TABLE `products`.`product` (`id` VARCHAR(36) NOT NULL, `name` VARCHAR(100) NOT NULL, `sold_by` VARCHAR(36) NOT NULL, `is_active` BOOLEAN NOT NULL DEFAULT 1, PRIMARY KEY(`id`));

/* FIXME: combine CREATE_USER and INSERT_NAME as one stored
   procedure that checks if the username exists and insert it otherwise
   return 1 if the username already exists and 0 if it created it 
 */

-- CREATE PROCEDURE CREATE_USER(IN GIVEN_USERNAME VARCHAR(32))
-- BEGIN
	-- SET @User_exists = 0;
	-- SELECT 1 INTO @User_exists
	-- FROM users
	-- WHERE username = GIVEN_USERNAME;       
	-- SELECT @User_exists
-- END

-- CREATE PROCEDURE InsertName
-- (
	-- @username varchar(25),
	-- @userpassword varchar(25)
-- )
-- AS
-- IF EXISTS(SELECT 'True' FROM MyTable WHERE username = @username)
-- BEGIN
	-- --This means it exists, return it to ASP and tell us
	-- SELECT 'This record already exists!'
-- END
-- ELSE
-- BEGIN
	-- --This means the record isn't in there already, let's go ahead and add it
	-- SELECT 'Record Added'
	-- INSERT into MyTable(username, userpassword) VALUES(@username, @userpassword)
-- END
