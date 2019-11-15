-- +migrate Up

CREATE TABLE role (
    role_id VARCHAR(26) NOT NULL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE
);

-- +migrate Down

DROP TABLE IF EXISTS role;
