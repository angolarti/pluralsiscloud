package dockerengine

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker struct {
	ctx    context.Context
	cli    *client.Client
	config Config
}

func NewDocker(ctx context.Context, cli *client.Client, config Config) *Docker {
	return &Docker{
		ctx:    ctx,
		cli:    cli,
		config: config,
	}
}

func (dc *Docker) Run(image string, command []string) (container.CreateResponse, error) {
	dc.Pull(dc.config.Image)
	resp, err := dc.cli.ContainerCreate(dc.ctx, &container.Config{
		Image: image,
		Cmd:   command,
		Tty:   false,
	}, nil, nil, nil, "")

	if err != nil {
		return container.CreateResponse{}, err
	}

	if err := dc.cli.ContainerStart(dc.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return container.CreateResponse{}, err
	}

	return resp, nil
}

func (dc *Docker) Start(containerID string) error {
	if err := dc.cli.ContainerStart(dc.ctx, containerID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := dc.cli.ContainerWait(dc.ctx, containerID, container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	return nil
}

func (dc *Docker) Logs(containerID string) error {
	out, err := dc.cli.ContainerLogs(dc.ctx, containerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	if err != nil {
		return err
	}

	return nil
}

func (dc *Docker) List() ([]types.Container, error) {
	defer dc.cli.Close()

	containers, err := dc.cli.ContainerList(dc.ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (dc *Docker) ListAllImages() ([]types.ImageSummary, error) {
	defer dc.cli.Close()

	images, err := dc.cli.ImageList(dc.ctx, types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (dc *Docker) Pull(imageName string) error {
	out, err := dc.cli.ImagePull(dc.ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		return err
	}

	return nil
}

func (dc *Docker) PullImageWithAuthentication(imageName, username, password string) error {
	authConfig := registry.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	options := types.ImagePullOptions{
		RegistryAuth: authStr,
	}
	out, err := dc.cli.ImagePull(dc.ctx, imageName, options)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		return err
	}

	return nil
}

func (dc *Docker) CommitAcontainer(imageName, username, password string) error {
	createResp, err := dc.cli.ContainerCreate(dc.ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"touch", "/helloworld"},
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}

	if err := dc.cli.ContainerStart(dc.ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := dc.cli.ContainerWait(dc.ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	commitResp, err := dc.cli.ContainerCommit(dc.ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
	if err != nil {
		return err
	}

	fmt.Println(commitResp.ID)

	return nil
}
