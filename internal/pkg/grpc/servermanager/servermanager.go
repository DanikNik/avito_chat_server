package servermanager

import (
	"chat_server/dbProto"
	"chat_server/internal/pkg/storage/adapters"
	"context"
)

type ServerManager struct {
	Storage adapters.StorageAdapter
}

func (s ServerManager) CreateUser(ctx context.Context, r *dbProto.UserCreateRequest) (*dbProto.UserModel, error) {
	return s.Storage.CreateUser(r)
}

func (s ServerManager) CreateChat(ctx context.Context, r *dbProto.ChatCreateRequest) (*dbProto.ChatModel, error) {
	return s.Storage.CreateChat(r)
}

func (s ServerManager) PostMessage(ctx context.Context, r *dbProto.PostMessageRequest) (*dbProto.MessageModel, error) {
	return s.Storage.PostMessage(r)
}

func (s ServerManager) ListUserChats(ctx context.Context, r *dbProto.ListUserChatsRequest) (*dbProto.ListUserChatsResponse, error) {
	return s.Storage.ListUserChats(r)
}

func (s ServerManager) ListChatMessages(ctx context.Context, r *dbProto.ListChatMessagesRequest) (*dbProto.ListChatMessagesResponse, error) {
	return s.Storage.ListChatMessages(r)
}
