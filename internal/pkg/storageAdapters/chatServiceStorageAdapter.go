package storageAdapters

import (
	"chat_server/dbProto"
	"chat_server/internal/pkg/grpcManager"
	"context"
)

type ChatStorageAdapter struct {
	Manager *grpcManager.Manager
}

func (c ChatStorageAdapter) CreateUser(request *dbProto.UserCreateRequest) (*dbProto.UserModel, error) {
	return c.Manager.DbClient.CreateUser(context.Background(), request)
}

func (c ChatStorageAdapter) CreateChat(request *dbProto.ChatCreateRequest) (*dbProto.ChatModel, error) {
	return c.Manager.DbClient.CreateChat(context.Background(), request)
}

func (c ChatStorageAdapter) PostMessage(request *dbProto.PostMessageRequest) (*dbProto.MessageModel, error) {
	return c.Manager.DbClient.PostMessage(context.Background(), request)
}

func (c ChatStorageAdapter) ListUserChats(request *dbProto.ListUserChatsRequest) (*dbProto.ListUserChatsResponse, error) {
	return c.Manager.DbClient.ListUserChats(context.Background(), request)
}

func (c ChatStorageAdapter) ListChatMessages(request *dbProto.ListChatMessagesRequest) (*dbProto.ListChatMessagesResponse, error) {
	return c.Manager.DbClient.ListChatMessages(context.Background(), request)
}

func NewChatStorageAdapter(manager *grpcManager.Manager) ChatStorageAdapter {
	return ChatStorageAdapter{Manager: manager}
}
