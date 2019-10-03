package handlers

import "net/http"

type HandlerSet struct{}

func (h *HandlerSet) CreateUser(w http.ResponseWriter, req *http.Request)   {}
func (h *HandlerSet) CreateChat(w http.ResponseWriter, req *http.Request)   {}
func (h *HandlerSet) PostMessage(w http.ResponseWriter, req *http.Request)  {}
func (h *HandlerSet) ListChats(w http.ResponseWriter, req *http.Request)    {}
func (h *HandlerSet) ListMessages(w http.ResponseWriter, req *http.Request) {}
