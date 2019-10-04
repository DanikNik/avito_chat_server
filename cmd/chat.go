package main

import (
	"chat_server/internal/app/chat_service"
	"chat_server/internal/pkg/grpcManager"
	"chat_server/internal/pkg/handlers"
	"chat_server/internal/pkg/storage/adapters"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
)

func main() {
	manager := grpcManager.NewManager(&grpc.ClientConn{})
	service := chat_service.NewChatService(
		mux.NewRouter(),
		&handlers.HandlerSet{},
		adapters.NewChatStorageAdapter(manager),
		chat_service.Config{Port: ":9000"},
	)
	log.Fatalln(service.Start())
}
