package app

import (
	"chat_server/internal/pkg/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Config struct {
	Port string
}

type ChatService struct {
	Router   *mux.Router
	Handlers handlers.HandlerSet
	Config   Config
}

func NewChatService(router *mux.Router) *ChatService {
	return &ChatService{Router: router}
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

	return http.ListenAndServe(c.Config.Port, c.Router)
}
