CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id    UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    username   VARCHAR(50)                 NOT NULL CHECK (username <> ''),
    email      VARCHAR(50) UNIQUE          NOT NULL CHECK (email <> ''),
    password   VARCHAR(100)                NOT NULL CHECK (password <> ''),
    created_at TIMESTAMP(0) WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    last_login TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

---- create above / drop below ----

DROP TABLE users;
