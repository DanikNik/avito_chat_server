package clientmanager

import (
	"chat_server/dbProto"
	"google.golang.org/grpc"
)

type ClientManager struct {
	conn     *grpc.ClientConn
	DbClient dbProto.DatabaseServiceClient
}

func NewClientManager(conn *grpc.ClientConn) *ClientManager {
	return &ClientManager{
		conn:     conn,
		DbClient: dbProto.NewDatabaseServiceClient(conn),
	}
}
