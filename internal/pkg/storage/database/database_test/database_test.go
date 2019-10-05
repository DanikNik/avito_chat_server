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
	"time"
)

type TestFixture struct {
	testConnPool     *pgx.ConnPool
	testContainerId  string
	testDockerClient *client.Client
}

var MainFixture = &TestFixture{}

func mockDbConnPool() (*pgx.ConnPool, error) {
	log.Println("Started mocking dbConn...")
	port := 12345

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

	var dbObj *pgx.ConnPool
	err := fmt.Errorf("not init")
	for range [5]interface{}{} {
		dbObj, err = pgx.NewConnPool(poolConfig)
		if err != nil {
			log.Printf("Retrying connection")
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("Unable to open connection: %v\n", err)
	}

	log.Println("Connection established.")
	return dbObj, nil
}

func setup() error {
	err := os.Chdir(os.Getenv("CHAT_PROJECT_ROOT"))
	if err != nil {
		return fmt.Errorf("Error changing workdir: %v", err)
	}
	log.Println("Started setup...")
	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("Error while creating Docker API client: %v", err.Error())
	}
	MainFixture.testDockerClient = cli
	id, err := dockerloader.CreateTestDbEnv(
		cli,
		".",
		"chat-test-db",
	)
	if err != nil {
		return err
	}
	MainFixture.testContainerId = id
	log.Println("Finished container initialization.")

	d, err := mockDbConnPool()
	if err != nil {
		return fmt.Errorf("Error while creating test database connection: %v", err.Error())
	}
	MainFixture.testConnPool = d

	log.Println("Finished setup.")
	return nil
}

func teardown(conn bool) {
	log.Println("Started teardown...")
	if conn {
		MainFixture.testConnPool.Close()
	}

	err := dockerloader.RemoveContainer(MainFixture.testDockerClient, MainFixture.testContainerId)
	if err != nil {
		log.Fatal(fmt.Errorf("Error while removing container: %v", err.Error()))
	}
	log.Println("Finished teardown")
}

func midTeardown() error {
	log.Println("Started midTeardown...")
	_, err := MainFixture.testConnPool.Exec(`
	TRUNCATE chat_service.users, chat_service.chats, chat_service.messages CASCADE;
`)
	if err != nil {
		return fmt.Errorf("Error while database intermidiate truncating: %v", err.Error())
	}
	log.Println("Finished.")
	return nil
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
	err := midTeardown()
	if err != nil {
		t.Fatalf("MidTeardown errored, test failed data structure: %v", err.Error())
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
	err := midTeardown()
	if err != nil {
		t.Fatalf("MidTeardown errored, test failed data structure: %v", err.Error())
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
	err := midTeardown()
	if err != nil {
		t.Fatalf("MidTeardown errored, test failed data structure: %v", err.Error())
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
	err := midTeardown()
	if err != nil {
		t.Fatalf("MidTeardown errored, test failed data structure: %v", err.Error())
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
	err := midTeardown()
	if err != nil {
		t.Fatalf("MidTeardown errored, test failed data structure: %v", err.Error())
	}
}

func TestMain(m *testing.M) {
	code := 1
	if err := setup(); err == nil {
		code = m.Run()
		teardown(true)
	} else {
		log.Printf("Setup error: %v", err.Error())
		teardown(false)
	}
	os.Exit(code)
}
