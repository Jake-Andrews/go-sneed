CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id    UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    username   VARCHAR(50)                 NOT NULL CHECK (username <> ''),
    email      VARCHAR(50) UNIQUE          NOT NULL CHECK (email <> ''),
    password   VARCHAR(100)                NOT NULL CHECK (password <> ''),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_login TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE videos (
    video_id       UUID PRIMARY KEY                        DEFAULT uuid_generate_v4(),
    user_id        UUID REFERENCES users(user_id),
    title          VARCHAR(100)                   NOT NULL CHECK (title <> ''),
    description    TEXT,
    duration       INTERVAL                       NOT NULL,
    file_path      TEXT                           NOT NULL,
    thumbnail_path TEXT                           NOT NULL,
    quality        JSONB                          NOT NULL,
    views          INTEGER DEFAULT 0,
    likes          INTEGER DEFAULT 0,
    dislikes       INTEGER DEFAULT 0,
    created_at     TIMESTAMP(0) WITH TIME ZONE    NOT NULL DEFAULT NOW(),
);

---- create above / drop below ----

DROP TABLE users;
DROP TABLE videos;
