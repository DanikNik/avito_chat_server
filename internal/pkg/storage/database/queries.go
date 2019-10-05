package database

const CreateUserQuery = `
INSERT INTO chat_service.users (nickname, created_at)
VALUES ($1, $2)
RETURNING id, nickname, created_at`

const CreateChatQuery = ``
