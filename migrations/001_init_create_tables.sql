CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id    UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    user_name  VARCHAR(50)                 NOT NULL CHECK (user_name <> ''),
    email      VARCHAR(50) UNIQUE          NOT NULL CHECK (email <> ''),
    PASSWORD   VARCHAR(32)                 NOT NULL CHECK (PASSWORD <> ''),
    created_at TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_date TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

---- create above / drop below ----

DROP TABLE users;
