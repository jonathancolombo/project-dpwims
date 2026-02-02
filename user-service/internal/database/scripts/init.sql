-- ============================================================
--  USERS TABLE INITIALIZATION SCRIPT
--  This script creates the "users" table with all constraints
--  needed for a production‑ready user service.
-- ============================================================
DROP DATABASE IF EXISTS identity_users;
CREATE DATABASE IF NOT EXISTS identity_users;
USE identity_users;
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users
(
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    username    VARCHAR(255)    NOT NULL,
    password    VARCHAR(255)    NOT NULL,
    password_salt VARCHAR(255),
    email       VARCHAR(255)    NOT NULL,
    fiscal_code VARCHAR(16)     NOT NULL,
    telephone   VARCHAR(20),

    role VARCHAR(50) NOT NULL DEFAULT ( 'user'),

    created_at  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    -- Email must be unique across the system
    UNIQUE KEY idx_users_email (email),

    -- Username often needs to be unique as well
    UNIQUE KEY idx_users_username (username),

    -- Fiscal code must be unique for each user
    UNIQUE idx_users_fiscal_code (fiscal_code),

    -- Telephone number must be unique for each user
    UNIQUE idx_users_telephone (telephone)
);
