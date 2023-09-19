-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema finalGogogo
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema finalGogogo
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `finalGogogo` DEFAULT CHARACTER SET utf8 ;
USE `finalGogogo` ;

-- -----------------------------------------------------
-- Table `finalGogogo`.`Dentists`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `finalGogogo`.`Dentists` (
  `idDentist` INT NOT NULL AUTO_INCREMENT,
  `last_name` VARCHAR(45) NULL,
  `first_name` VARCHAR(45) NULL,
  `registration_number` INT NULL,
  PRIMARY KEY (`idDentist`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `finalGogogo`.`Patients`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `finalGogogo`.`Patients` (
  `idPatient` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `address` VARCHAR(100) NULL,
  `credential_id` VARCHAR(45) NULL,
  `discharge_date` DATE NULL,
  PRIMARY KEY (`idPatient`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `finalGogogo`.`Appointments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `finalGogogo`.`Appointments` (
  `idAppointments` INT NOT NULL AUTO_INCREMENT,
  `date` DATE NULL,
  `notes` VARCHAR(100) NULL,
  `Dentists_idDentist` INT NOT NULL,
  `Patients_idPatient` INT NOT NULL,
  PRIMARY KEY (`idAppointments`, `Dentists_idDentist`, `Patients_idPatient`),
  INDEX `fk_Appointments_Dentists_idx` (`Dentists_idDentist` ASC),
  INDEX `fk_Appointments_Patients1_idx` (`Patients_idPatient` ASC),
  CONSTRAINT `fk_Appointments_Dentists`
    FOREIGN KEY (`Dentists_idDentist`)
    REFERENCES `finalGogogo`.`Dentists` (`idDentist`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Appointments_Patients1`
    FOREIGN KEY (`Patients_idPatient`)
    REFERENCES `finalGogogo`.`Patients` (`idPatient`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
