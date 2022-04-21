-- MySQL Script generated by MySQL Workbench
-- Fri Apr 22 03:19:55 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema famiphoto_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema famiphoto_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `famiphoto_db` DEFAULT CHARACTER SET utf8 ;
USE `famiphoto_db` ;

-- -----------------------------------------------------
-- Table `famiphoto_db`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`users` (
  `user_id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(256) NOT NULL,
  `status` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`groups`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`groups` (
  `group_id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(256) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`group_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photos` (
  `photo_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(1024) NOT NULL,
  `file_path` TEXT NOT NULL,
  `imported_at` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `group_id` VARCHAR(128) NOT NULL,
  `owner_id` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`photo_id`),
  INDEX `fk_photos_groups_idx` (`group_id` ASC) VISIBLE,
  INDEX `fk_photos_users_idx` (`owner_id` ASC) VISIBLE,
  CONSTRAINT `fk_photos_groups`
    FOREIGN KEY (`group_id`)
    REFERENCES `famiphoto_db`.`groups` (`group_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_photos_users`
    FOREIGN KEY (`owner_id`)
    REFERENCES `famiphoto_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`user_passwords`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`user_passwords` (
  `user_id` VARCHAR(128) NOT NULL,
  `password` VARCHAR(512) NOT NULL,
  `last_modified_at` DATETIME NOT NULL,
  `is_initialized` TINYINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_passwords_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `famiphoto_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`group_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`group_users` (
  `group_id` VARCHAR(128) NOT NULL,
  `user_id` VARCHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`group_id`, `user_id`),
  INDEX `fk_group_users_users_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_group_users_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `famiphoto_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_group_users_groups`
    FOREIGN KEY (`group_id`)
    REFERENCES `famiphoto_db`.`groups` (`group_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`oauth_clients`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`oauth_clients` (
  `oauth_client_id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(256) NOT NULL,
  `client_secret` TEXT NOT NULL,
  `scope` VARCHAR(1024) NOT NULL,
  `client_type` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`oauth_client_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`user_auth`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`user_auth` (
  `user_id` VARCHAR(128) NOT NULL,
  `oauth_client_id` VARCHAR(128) NOT NULL,
  `refresh_token` VARCHAR(1024) NOT NULL,
  `refresh_token_published_at` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `oauth_client_id`),
  INDEX `fk_user_auth_oauth_clients_idx` (`oauth_client_id` ASC) VISIBLE,
  UNIQUE INDEX `refresh_token_UNIQUE` (`refresh_token` ASC) VISIBLE,
  CONSTRAINT `fk_user_auth_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `famiphoto_db`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_auth_oauth_clients`
    FOREIGN KEY (`oauth_client_id`)
    REFERENCES `famiphoto_db`.`oauth_clients` (`oauth_client_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`oauth_client_redirect_urls`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`oauth_client_redirect_urls` (
  `oauth_client_id` VARCHAR(128) NOT NULL,
  `redirect_url` TEXT NOT NULL,
  `oauth_client_redirect_url_id` INT NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`oauth_client_redirect_url_id`),
  INDEX `fk_oauth_client_redirect_url_oauth_client_idx` (`oauth_client_id` ASC) VISIBLE,
  CONSTRAINT `fk_oauth_client_redirect_url_oauth_client`
    FOREIGN KEY (`oauth_client_id`)
    REFERENCES `famiphoto_db`.`oauth_clients` (`oauth_client_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`exif`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`exif` (
  `exif_id` INT NOT NULL AUTO_INCREMENT,
  `photo_id` INT NOT NULL,
  `tag_id` INT NOT NULL,
  `tag_name` VARCHAR(512) NOT NULL,
  `tag_type` VARCHAR(128) NOT NULL,
  `value` BLOB NOT NULL,
  `value_string` TEXT NOT NULL,
  `sort_order` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`exif_id`),
  INDEX `fk_exif_photos_idx` (`photo_id` ASC) VISIBLE,
  CONSTRAINT `fk_exif_photos`
    FOREIGN KEY (`photo_id`)
    REFERENCES `famiphoto_db`.`photos` (`photo_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
