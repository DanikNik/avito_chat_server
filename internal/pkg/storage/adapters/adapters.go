package adapters

import "chat_server/dbProto"

type StorageAdapter interface {
	CreateUser(request *dbProto.UserCreateRequest) (*dbProto.UserModel, error)
	CreateChat(request *dbProto.ChatCreateRequest) (*dbProto.ChatModel, error)
	PostMessage(request *dbProto.PostMessageRequest) (*dbProto.MessageModel, error)
	ListUserChats(request *dbProto.ListUserChatsRequest) (*dbProto.ListUserChatsResponse, error)
	ListChatMessages(request *dbProto.ListChatMessagesRequest) (*dbProto.ListChatMessagesResponse, error)
}
