-- +migrate Up

CREATE TABLE role (
    role_id BIGINT UNSIGNED NOT NULL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);

-- +migrate Down

DROP TABLE IF EXISTS role;
