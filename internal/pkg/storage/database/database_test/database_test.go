package database_test

import (
	"chat_server/dbProto"
	"chat_server/internal/pkg/storage/database"
	"chat_server/internal/pkg/testloaders/dockerloader"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/jackc/pgx"
	"log"
	"os"
	"reflect"
	"testing"
)

type TestFixture struct {
	testConnPool     *pgx.ConnPool
	testContainerId  string
	testDockerClient *client.Client
}

var MainFixture = &TestFixture{}

func setup() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(fmt.Errorf("Error while creating Docker API client: %v", err.Error()))
	}
	MainFixture.testDockerClient = cli
	id, _ := dockerloader.CreateTestDbEnv(cli, ".", "chat-test-db")
	MainFixture.testContainerId = id

	//TODO: add database connection init

	log.Println("Finished setup.")
}

func teardown() {
	log.Println("Started teardown...")
	err := dockerloader.RemoveContainer(MainFixture.testDockerClient, MainFixture.testContainerId)
	if err != nil {
		panic(err)
	}
}

func TestDbStorageAdapter_CreateChat(t *testing.T) {
	type fields struct {
		DbObj *pgx.ConnPool
	}
	type args struct {
		request *dbProto.ChatCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dbProto.ChatModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.DbStorageAdapter{
				DbObj: tt.fields.DbObj,
			}
			got, err := d.CreateChat(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateChat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbStorageAdapter_CreateUser(t *testing.T) {
	type fields struct {
		DbObj *pgx.ConnPool
	}
	type args struct {
		request *dbProto.UserCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantU   *dbProto.UserModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.DbStorageAdapter{
				DbObj: tt.fields.DbObj,
			}
			gotU, err := d.CreateUser(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("CreateUser() gotU = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func TestDbStorageAdapter_ListChatMessages(t *testing.T) {
	type fields struct {
		DbObj *pgx.ConnPool
	}
	type args struct {
		request *dbProto.ListChatMessagesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dbProto.ListChatMessagesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.DbStorageAdapter{
				DbObj: tt.fields.DbObj,
			}
			got, err := d.ListChatMessages(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListChatMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListChatMessages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbStorageAdapter_ListUserChats(t *testing.T) {
	type fields struct {
		DbObj *pgx.ConnPool
	}
	type args struct {
		request *dbProto.ListUserChatsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dbProto.ListUserChatsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.DbStorageAdapter{
				DbObj: tt.fields.DbObj,
			}
			got, err := d.ListUserChats(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUserChats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUserChats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbStorageAdapter_PostMessage(t *testing.T) {
	type fields struct {
		DbObj *pgx.ConnPool
	}
	type args struct {
		request *dbProto.PostMessageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dbProto.MessageModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := database.DbStorageAdapter{
				DbObj: tt.fields.DbObj,
			}
			got, err := d.PostMessage(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	log.Println("middleware")
	code := m.Run()
	os.Exit(code)
}
