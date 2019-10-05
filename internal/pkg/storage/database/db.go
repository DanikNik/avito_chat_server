package database

import (
	"chat_server/dbProto"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"time"
)

type DbStorageAdapter struct {
	DbObj *pgx.ConnPool
}

var (
	NotInit     = fmt.Errorf("db wasn't initialized")
	AlreadyInit = fmt.Errorf("db already initialized")
	ErrNotFound = fmt.Errorf("not found")
	ErrConflict = fmt.Errorf("conflict")
	ErrNotNull  = fmt.Errorf("not null violation")
)

const (
	uniqueIntegrityError = "23505"
	//foreignKeyError      = "23503"
	notNullError = "23502"
)

func MakeProdConnPool() (*pgx.ConnPool, error) {

	port := 5432

	connConfig := pgx.ConnConfig{
		User:              "postgres",
		Password:          "postgres",
		Host:              "localhost",
		Port:              uint16(port),
		Database:          "chat",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 50,
		AcquireTimeout: 10 * time.Second,
		AfterConnect:   nil,
	}

	dbObj, err := pgx.NewConnPool(poolConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to establish connection: %v\n", err)
	}
	log.Println("Connection established...")
	return dbObj, nil
}

func NewDbStorageAdapter(cp *pgx.ConnPool) (DbStorageAdapter, error) {
	a := DbStorageAdapter{}
	a.DbObj = cp
	return a, nil
}

func (d DbStorageAdapter) CreateUser(request *dbProto.UserCreateRequest) (*dbProto.UserModel, error) {

	if request.Username == "" {
		return nil, ErrNotNull
	}
	row := d.DbObj.QueryRow(CreateUserQuery, request.GetUsername(), request.GetCreatedAt())

	u := &dbProto.UserModel{}
	err := row.Scan(
		&u.Id,
		&u.Username,
		&u.CreatedAt,
	)

	if err != nil {
		if pqError, ok := err.(pgx.PgError); ok {
			switch pqError.Code {
			case uniqueIntegrityError:
				return nil, ErrConflict
			case notNullError:
				return nil, ErrNotNull
			default:
				return nil, err
			}
		}
	}
	return u, nil
}

func (d DbStorageAdapter) CreateChat(request *dbProto.ChatCreateRequest) (*dbProto.ChatModel, error) {
	panic("implement me")
}

func (d DbStorageAdapter) PostMessage(request *dbProto.PostMessageRequest) (*dbProto.MessageModel, error) {
	panic("implement me")
}

func (d DbStorageAdapter) ListUserChats(request *dbProto.ListUserChatsRequest) (*dbProto.ListUserChatsResponse, error) {
	panic("implement me")
}

func (d DbStorageAdapter) ListChatMessages(request *dbProto.ListChatMessagesRequest) (*dbProto.ListChatMessagesResponse, error) {
	panic("implement me")
}
