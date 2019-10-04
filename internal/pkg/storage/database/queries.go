package database

const CreateUserQuery = `
INSERT INTO chat_service.users (nickname)
VALUES ($1)
RETURNING id, nickname, extract(epoch from created_at)::int;`

const CreateChatQuery = ``
