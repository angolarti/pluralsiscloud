package container

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker struct {
	context   context.Context
	client    client.Client
	ImageName string
	Hostname  string
	Command   strslice.StrSlice
}

func NewDocker(ctx context.Context, cli *client.Client, imageName string, command strslice.StrSlice) *Docker {
	return &Docker{
		context:   ctx,
		client:    *cli,
		ImageName: imageName,
		Command:   command,
	}
}

func (d *Docker) Run() (container.CreateResponse, error) {

	d.Validate()

	d.Pull(d.context, d.client, d.ImageName)
	resp, err := d.client.ContainerCreate(d.context, &container.Config{
		Image: d.ImageName,
		Cmd:   d.Command,
		Tty:   false,
	}, nil, nil, nil, "")

	if err != nil {
		panic(err)
	}

	if err := d.client.ContainerStart(d.context, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return resp, nil
}

func (d *Docker) Start(containerID string) error {

	if err := d.client.ContainerStart(d.context, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := d.client.ContainerWait(d.context, containerID, container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	return nil
}

// TODO:

// Pull an image with authentication

// Commit a containe

func (d *Docker) Logs(containerID string) {
	out, err := d.client.ContainerLogs(d.context, containerID, types.ContainerLogsOptions{ShowStdout: true})

	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

func (d *Docker) List() ([]types.Container, error) {
	defer d.client.Close()

	containers, err := d.client.ContainerList(d.context, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	return containers, nil
}

func (d *Docker) ListAllImages() ([]types.ImageSummary, error) {
	defer d.client.Close()

	images, err := d.client.ImageList(d.context, types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	return images, nil

}

func (d *Docker) Pull(ctx context.Context, cli client.Client, imageName string) {

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
}

func (d *Docker) Validate() error {

	if d.ImageName == "" {
		return errors.New("image name is required")
	}

	return nil
}

func (d *Docker) PullImageWithAuthentication(ctx context.Context, imageName, username, password string) error {
	authConfig := types.AuthConfig{
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
	out, err := d.client.ImagePull(ctx, imageName, options)
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	return nil
}

func (d *Docker) CommitAcontainer(ctx context.Context, imageName, username, password string) error {
	createResp, err := d.client.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"touch", "/helloworld"},
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}

	if err := d.client.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := d.client.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	commitResp, err := d.client.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})

	if err != nil {
		return err
	}

	fmt.Println(commitResp.ID)

	return nil
}
