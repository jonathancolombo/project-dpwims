DROP DATABASE IF EXISTS identity_subscriptions;
CREATE DATABASE IF NOT EXISTS identity_subscriptions;
USE identity_subscriptions;
DROP TABLE IF EXISTS subscriptions;
CREATE TABLE IF NOT EXISTS subscriptions
(
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id     BIGINT UNSIGNED NOT NULL,
    train_uuid  VARCHAR(36)    NOT NULL,
    schedule_id BIGINT UNSIGNED NOT NULL,
    created_at  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES identity_users.users (id) ON DELETE CASCADE,
    FOREIGN KEY (train_uuid) REFERENCES trains_db.trains (uuid) ON DELETE CASCADE,
    FOREIGN KEY (schedule_id) REFERENCES trains_db.schedules(id) ON DELETE CASCADE
)