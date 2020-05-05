CREATE SCHEMA IF NOT EXISTS `documents_api` DEFAULT CHARACTER SET utf8;
USE `documents_api`;

CREATE TABLE IF NOT EXISTS `documents_api`.`users` (
  `id`         VARCHAR(255) NOT NULL,
  `username`   VARCHAR(255) NOT NULL UNIQUE,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `documents_api`.`user_auths` (
  `user_id`         VARCHAR(255) NOT NULL,
  `email`           VARCHAR(255) NOT NULL UNIQUE,
  `hash`            VARCHAR(255) NOT NULL,
  `updated_at`      DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_auths_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `documents_api`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `documents_api`.`auth_tokens` (
  `id`              VARCHAR(255) NOT NULL,
  `user_id`         VARCHAR(255) NOT NULL,
  `token`           VARCHAR(255) NOT NULL,
  `expiry`          DATETIME,
  `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_auth_tokens_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `documents_api`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
ENGINE = InnoDB;
