-- +migrate Up

CREATE TABLE user_role (
    user_id VARCHAR(26) NOT NULL,
    role_id VARCHAR(26) NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

-- +migrate Down

DROP TABLE IF EXISTS user_role;
