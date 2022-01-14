-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema libDb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema libDb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `libDb` DEFAULT CHARACTER SET utf8 ;
USE `libDb` ;

-- -----------------------------------------------------
-- Table `libDb`.`persona`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `libDb`.`persona` (
  `idpersona` INT NOT NULL,
  `nombre` VARCHAR(45) NOT NULL,
  `apellido` VARCHAR(45) NOT NULL,
  `telefono` VARCHAR(45) NOT NULL,
  `correo` VARCHAR(45) NOT NULL,
  `direccion` VARCHAR(45) NOT NULL,
  `prestamos` INT NULL,
  PRIMARY KEY (`idpersona`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `libDb`.`libro`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `libDb`.`libro` (
  `idlibro` INT NOT NULL,
  `titulo` VARCHAR(45) NOT NULL,
  `autor` VARCHAR(45) NULL,
  `fecha` DATETIME NULL,
  `categoria` VARCHAR(45) NULL,
  PRIMARY KEY (`idlibro`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `libDb`.`prestamos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `libDb`.`prestamos` (
  `idprestamos` INT NOT NULL,
  `persona_idpersona` INT NOT NULL,
  `libro_idlibro` INT NOT NULL,
  `fechaPrestamo` DATETIME NOT NULL,
  `fechaDevolucion` DATETIME NOT NULL,
  PRIMARY KEY (`idprestamos`, `persona_idpersona`, `libro_idlibro`),
  INDEX `fk_prestamos_persona1_idx` (`persona_idpersona` ASC) VISIBLE,
  INDEX `fk_prestamos_libro1_idx` (`libro_idlibro` ASC) VISIBLE,
  CONSTRAINT `fk_prestamos_persona1`
    FOREIGN KEY (`persona_idpersona`)
    REFERENCES `libDb`.`persona` (`idpersona`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_prestamos_libro1`
    FOREIGN KEY (`libro_idlibro`)
    REFERENCES `libDb`.`libro` (`idlibro`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
