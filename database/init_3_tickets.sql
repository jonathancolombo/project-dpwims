DROP DATABASE IF EXISTS tickets_db;
CREATE DATABASE IF NOT EXISTS tickets_db;
USE tickets_db;
DROP TABLE IF EXISTS tickets;
CREATE TABLE IF NOT EXISTS tickets
(
    uuid         VARCHAR(36),
    user_id      BIGINT UNSIGNED NOT NULL,
    train_uuid      VARCHAR(36) NOT NULL,
    schedule_id   BIGINT UNSIGNED NOT NULL,
    seat_number   VARCHAR(10) NOT NULL,
    price         DOUBLE NOT NULL,
    status        VARCHAR(50) NOT NULL DEFAULT 'booked',
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (uuid),
    FOREIGN KEY (user_id) REFERENCES identity_users.users (id) ON DELETE CASCADE,
    FOREIGN KEY (train_uuid) REFERENCES trains_db.trains (uuid) ON DELETE CASCADE,
    FOREIGN KEY (schedule_id) REFERENCES trains_db.schedules (id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS payments;
CREATE TABLE IF NOT EXISTS payments
(
    uuid           VARCHAR(36),
    ticket_id      VARCHAR(36) NOT NULL,
    amount         DOUBLE NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    provider_reference VARCHAR(255) NOT NULL DEFAULT 'nothing',
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (uuid),
    FOREIGN KEY (ticket_id) REFERENCES tickets (uuid) ON DELETE CASCADE
);