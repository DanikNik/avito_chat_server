CREATE SCHEMA IF NOT EXISTS chat_service;

CREATE TABLE IF NOT EXISTS chat_service.users
(
    id         BIGSERIAL PRIMARY KEY NOT NULL,
    nickname   TEXT UNIQUE           NOT NULL,
    created_at timestamptz           NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_service.chats
(
    id         BIGSERIAL PRIMARY KEY NOT NULL,
    name       TEXT UNIQUE,
    created_at timestamptz           NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chat_service.chat_to_user_rel
(
    chat_id BIGINT NOT NULL REFERENCES chat_service.chats (id),
    user_id BIGINT NOT NULL REFERENCES chat_service.users (id),
    CONSTRAINT chat_user_pkey PRIMARY KEY (chat_id, user_id)
);

CREATE TABLE IF NOT EXISTS chat_service.messages
(
    id         BIGSERIAL PRIMARY KEY                     NOT NULL,
    chat       BIGINT REFERENCES chat_service.chats (id) NOT NULL,
    author     BIGINT REFERENCES chat_service.users (id) NOT NULL,
    text       TEXT                                      NOT NULL,
    created_at timestamptz                               NOT NULL DEFAULT now()
);