DROP DATABASE IF EXISTS trains_service;
CREATE DATABASE IF NOT EXISTS trains_service;
USE trains_service;
DROP TABLE IF EXISTS trains;
CREATE TABLE IF NOT EXISTS trains
(
    id           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    train_number VARCHAR(50)     NOT NULL,
    type         VARCHAR(50)     NOT NULL,
    capacity     INT             NOT NULL,
    status       VARCHAR(50)     NOT NULL DEFAULT 'active',
    created_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE KEY train_number (train_number)
);
DROP TABLE IF EXISTS routes;
CREATE TABLE IF NOT EXISTS routes
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    train_id       BIGINT UNSIGNED NOT NULL,
    departure      VARCHAR(100)    NOT NULL,
    arrival        VARCHAR(100)    NOT NULL,
    departure_time TIMESTAMP       NOT NULL,
    arrival_time   TIMESTAMP       NOT NULL,
    distance       INT             NOT NULL,
    created_at     TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (train_id) REFERENCES trains (id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS stations;
CREATE TABLE IF NOT EXISTS stations
(
    id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name       VARCHAR(100)    NOT NULL,
    city       VARCHAR(100)    NOT NULL,
    region     VARCHAR(100)    NOT NULL,
    status     VARCHAR(50)     NOT NULL DEFAULT 'active',
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY name (name)
);

DROP TABLE IF EXISTS schedules;
CREATE TABLE IF NOT EXISTS schedules
(
    id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    train_id   BIGINT UNSIGNED NOT NULL,
    station_id BIGINT UNSIGNED NOT NULL,
    departure  VARCHAR(100)    NOT NULL,
    arrival    VARCHAR(100)    NOT NULL,
    status     VARCHAR(50)     NOT NULL DEFAULT 'active',
    price      INT             NOT NULL,
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (train_id) REFERENCES trains (id) ON DELETE CASCADE,
    FOREIGN KEY (station_id) REFERENCES stations (id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS schedules_stops;
CREATE TABLE IF NOT EXISTS schedules_stops
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    schedule_id    BIGINT UNSIGNED NOT NULL,
    station_id     BIGINT UNSIGNED NOT NULL,
    station_name   VARCHAR(100)    NOT NULL REFERENCES stations (name) ON DELETE CASCADE,
    stop_order     INT             NOT NULL,
    arrival_time   TIME            NOT NULL,
    departure_time TIME            NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (schedule_id) REFERENCES schedules (id) ON DELETE CASCADE,
    FOREIGN KEY (station_id) REFERENCES stations (id) ON DELETE CASCADE
)


