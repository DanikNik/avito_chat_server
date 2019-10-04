package chat_service

import (
	"chat_server/internal/pkg/handlers"
	"chat_server/internal/pkg/storageAdapters"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Config struct {
	Port string
}

type ChatService struct {
	Router   *mux.Router
	Handlers *handlers.HandlerSet
	Storage  storageAdapters.StorageAdapter
	Config   Config
}

func NewChatService(
	router *mux.Router,
	handlers *handlers.HandlerSet,
	storage storageAdapters.StorageAdapter,
	config Config,
) *ChatService {
	return &ChatService{
		Router:   router,
		Handlers: handlers,
		Storage:  storage,
		Config:   config,
	}
}

func (c *ChatService) Start() error {
	userRouter := c.Router.PathPrefix("/users").Subrouter()
	chatRouter := c.Router.PathPrefix("/chats").Subrouter()
	messageRouter := c.Router.PathPrefix("/messages").Subrouter()

	userRouter.HandleFunc("/add", c.Handlers.CreateUser).Methods("POST")
	chatRouter.HandleFunc("/add", c.Handlers.CreateChat).Methods("POST")
	chatRouter.HandleFunc("/get", c.Handlers.ListChats).Methods("GET")
	messageRouter.HandleFunc("/add", c.Handlers.PostMessage).Methods("POST")
	messageRouter.HandleFunc("/get", c.Handlers.ListMessages).Methods("GET")
	log.Println("Server started ad", c.Config.Port)

	return http.ListenAndServe(c.Config.Port, c.Router)
}
