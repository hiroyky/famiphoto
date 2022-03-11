-- MySQL Script generated by MySQL Workbench
-- Fri Mar 11 09:04:07 2022
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
  `group_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(256) NOT NULL,
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
  `group_id` INT NOT NULL,
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
  `group_id` INT NOT NULL,
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
-- Table `famiphoto_db`.`photo_exif`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `famiphoto_db`.`photo_exif` (
  `photo_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`photo_id`),
  CONSTRAINT `fk_photo_exif_photos`
    FOREIGN KEY (`photo_id`)
    REFERENCES `famiphoto_db`.`photos` (`photo_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
