package main

import (
	"chat_server/dbProto"
	"chat_server/internal/pkg/storage/database"
	"log"
)

func main() {
	db, err := database.NewDbStorageAdapter()
	if err != nil {
		panic(err)
	}
	u := &dbProto.UserCreateRequest{
		Username: "slava123456",
	}
	a, e := db.CreateUser(u)
	if e != nil {
		panic(e)
	}
	log.Println(a)
}
