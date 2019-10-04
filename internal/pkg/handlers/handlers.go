package handlers

import (
	"chat_server/dbProto"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HandlerSet struct{}

func (h *HandlerSet) CreateUser(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	userCreateReqData := dbProto.UserCreateRequest{}
	err := json.NewDecoder(req.Body).Decode(&userCreateReqData)
	if err != nil {
		log.Println(err)
		http.Error(w, "decoding error", http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", userCreateReqData)
	data, _ := json.Marshal(&userCreateReqData)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (h *HandlerSet) CreateChat(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	chatCreateReqData := dbProto.ChatCreateRequest{}
	err := json.NewDecoder(req.Body).Decode(&chatCreateReqData)
	if err != nil {
		log.Println(err)
		http.Error(w, "decoding error", http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", chatCreateReqData)
	data, _ := json.Marshal(&chatCreateReqData)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (h *HandlerSet) PostMessage(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	postMessageReqData := dbProto.PostMessageRequest{}
	err := json.NewDecoder(req.Body).Decode(&postMessageReqData)
	if err != nil {
		log.Println(err)
		http.Error(w, "decoding error", http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", postMessageReqData)
	data, _ := json.Marshal(&postMessageReqData)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (h *HandlerSet) ListChats(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	chatListReqData := dbProto.ListUserChatsRequest{}
	err := json.NewDecoder(req.Body).Decode(&chatListReqData)
	if err != nil {
		log.Println(err)
		http.Error(w, "decoding error", http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", chatListReqData)
	data, _ := json.Marshal(&chatListReqData)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func (h *HandlerSet) ListMessages(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	messageListReqData := dbProto.ListChatMessagesRequest{}
	err := json.NewDecoder(req.Body).Decode(&messageListReqData)
	if err != nil {
		log.Println(err)
		http.Error(w, "decoding error", http.StatusBadRequest)
	}

	fmt.Printf("%+v\n", messageListReqData)
	data, _ := json.Marshal(&messageListReqData)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
