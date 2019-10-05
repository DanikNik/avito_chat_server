package dockerloader

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/docker/go-connections/nat"
	"github.com/jhoonb/archivex"
	"os"
	"time"
)

func tarBuildContext(rootPath string) {
	tar := new(archivex.TarFile)
	tar.Create("/tmp/chat-test-db.tar")

	tar.AddAll(rootPath+"/build", true)
	tar.AddAll(rootPath+"/init", true)
	tar.Close()
}

func BuildImage(dockerBuildCtxDir string, cli *client.Client, ctx context.Context, tagName string) error {
	tarBuildContext(dockerBuildCtxDir)
	dockerBuildContext, err := os.Open("/tmp/chat-test-db.tar")
	if err != nil {
		return fmt.Errorf("Build Context reading error: %v", err.Error())
	}
	defer dockerBuildContext.Close()

	buildResp, err := cli.ImageBuild(ctx, dockerBuildContext, types.ImageBuildOptions{
		Tags:           []string{tagName},
		Dockerfile:     dockerBuildCtxDir + "/build/db.Dockerfile",
		SuppressOutput: false,
	})
	if err != nil {
		return fmt.Errorf("Error building image: %s", err.Error())
	}
	defer buildResp.Body.Close()
	termFd, isTerm := term.GetFdInfo(os.Stdout)
	_ = jsonmessage.DisplayJSONMessagesStream(buildResp.Body, os.Stdout, termFd, isTerm, nil)
	return nil
}

func CreateContainer(cli *client.Client, ctx context.Context, tagName string) (container.ContainerCreateCreatedBody, error) {
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "12345",
	}
	containerPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return container.ContainerCreateCreatedBody{}, fmt.Errorf("Unable to get the port: %v", err.Error())
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	containerResp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: tagName,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		},
		nil,
		fmt.Sprintf("%s-build", tagName),
	)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, fmt.Errorf("Error creating container: %v", err.Error())
	}
	return containerResp, nil
}

func CreateTestDbEnv(cli *client.Client, dockerBuildCtxDir, tagName string) (string, error) {
	ctx := context.Background()
	err := BuildImage(dockerBuildCtxDir, cli, ctx, tagName)
	if err != nil {
		return "", err
	}

	containerResp, err := CreateContainer(cli, ctx, tagName)
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, containerResp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("Error starting container: %v", err.Error())
	}

	fmt.Println(containerResp.ID)

	data, err := cli.ContainerInspect(ctx, containerResp.ID)
	if err != nil {
		return containerResp.ID, fmt.Errorf("Error inspecting container: %v", err.Error())
	}
	timer := time.NewTimer(60 * time.Second)
	for !data.State.Running {
		select {
		case <-timer.C:
			return containerResp.ID, fmt.Errorf("Container startup timed out")
		}
	}
	timer.Stop()
	return containerResp.ID, nil
}

func RemoveContainer(cli *client.Client, containerID string) error {
	err := cli.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		return fmt.Errorf("Error stopping container: %v", err.Error())
	}
	err = cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
	if err != nil {
		return fmt.Errorf("Error deleting container: %v", err.Error())
	}
	return nil
}
