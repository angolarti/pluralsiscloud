package container

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
		panic(err)
	}

	if err := dc.cli.ContainerStart(dc.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return resp, nil
}

func (dc *Docker) Start(containerID string) error {

	if err := dc.cli.ContainerStart(dc.ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := dc.cli.ContainerWait(dc.ctx, containerID, container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	return nil
}

func (dc *Docker) Logs(containerID string) {
	_, err := dc.cli.ContainerLogs(dc.ctx, containerID, types.ContainerLogsOptions{ShowStdout: true})

	if err != nil {
		panic(err)
	}

	// std.copy.StdC.configopy(os.Stdout, os.Stderr, out)
}

func (dc *Docker) List() ([]types.Container, error) {
	defer dc.cli.Close()

	containers, err := dc.cli.ContainerList(dc.ctx, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	return containers, nil
}

func (dc *Docker) ListAllImages() ([]types.ImageSummary, error) {
	defer dc.cli.Close()

	images, err := dc.cli.ImageList(dc.ctx, types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	return images, nil
}

func (dc *Docker) Pull(imageName string) {

	out, err := dc.cli.ImagePull(dc.ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
}

// TODO:

// Pull an image with authentication

// Commit a containe
