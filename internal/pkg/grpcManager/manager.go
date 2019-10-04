package grpcManager

import (
	"chat_server/dbProto"
	"google.golang.org/grpc"
)

type Manager struct {
	conn     *grpc.ClientConn
	DbClient dbProto.DatabaseServiceClient
}

func NewManager(conn *grpc.ClientConn) *Manager {
	return &Manager{
		conn:     conn,
		DbClient: dbProto.NewDatabaseServiceClient(conn),
	}
}
