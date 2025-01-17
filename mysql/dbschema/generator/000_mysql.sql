-- MySQL Script generated by MySQL Workbench
-- Sat Jan 11 16:04:35 2025
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema famiphoto_db
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `famiphoto_db` ;

-- -----------------------------------------------------
-- Schema famiphoto_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `famiphoto_db` DEFAULT CHARACTER SET utf8 ;
USE `famiphoto_db` ;

-- -----------------------------------------------------
-- Table `famiphoto_db`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `famiphoto_db`.`users` ;

CREATE TABLE IF NOT EXISTS `famiphoto_db`.`users` (
  `user_id` VARCHAR(128) NOT NULL,
  `name` VARCHAR(256) NOT NULL,
  `status` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photos`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `famiphoto_db`.`photos` ;

CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photos` (
  `photo_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(1024) NOT NULL,
  `imported_at` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `file_name_hash` VARCHAR(128) NOT NULL,
  `description_en` TEXT NOT NULL,
  `description_ja` TEXT NOT NULL,
  PRIMARY KEY (`photo_id`),
  UNIQUE INDEX `file_name_UNIQUE` (`file_name_hash` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`user_passwords`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `famiphoto_db`.`user_passwords` ;

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
-- Table `famiphoto_db`.`oauth_clients`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `famiphoto_db`.`oauth_clients` ;

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
DROP TABLE IF EXISTS `famiphoto_db`.`user_auth` ;

CREATE TABLE IF NOT EXISTS `famiphoto_db`.`user_auth` (
  `user_id` VARCHAR(128) NOT NULL,
  `oauth_client_id` VARCHAR(128) NOT NULL,
  `refresh_token` VARCHAR(1024) CHARACTER SET 'ascii' NOT NULL,
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
DROP TABLE IF EXISTS `famiphoto_db`.`oauth_client_redirect_urls` ;

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
DROP TABLE IF EXISTS `famiphoto_db`.`exif` ;

CREATE TABLE IF NOT EXISTS `famiphoto_db`.`exif` (
  `exif_id` INT NOT NULL AUTO_INCREMENT,
  `photo_id` INT NOT NULL,
  `tag_id` INT NOT NULL,
  `tag_name` VARCHAR(512) NOT NULL,
  `tag_type` VARCHAR(128) NOT NULL,
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


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photo_files`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `famiphoto_db`.`photo_files` ;

CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photo_files` (
  `photo_file_id` INT NOT NULL AUTO_INCREMENT,
  `photo_id` INT NOT NULL,
  `file_type` VARCHAR(45) NOT NULL,
  `file_path` TEXT NOT NULL,
  `imported_at` DATETIME NOT NULL,
  `file_hash` VARCHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`photo_file_id`),
  INDEX `fk_photo_files_photos_idx` (`photo_id` ASC) VISIBLE,
  CONSTRAINT `fk_photo_files_photos`
    FOREIGN KEY (`photo_id`)
    REFERENCES `famiphoto_db`.`photos` (`photo_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `famiphoto_db`.`photo_thumbnails`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `famiphoto_db`.`photo_thumbnails` ;

CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photo_thumbnails` (
  `photo_id` INT NOT NULL,
  `thumbnail_name` VARCHAR(128) NOT NULL,
  `file_path` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`photo_id`, `thumbnail_name`),
  CONSTRAINT `fk_photo_thumbanils_photos`
    FOREIGN KEY (`photo_id`)
    REFERENCES `famiphoto_db`.`photos` (`photo_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
