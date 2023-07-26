package infra

import (
	"context"
	"github/angolarti/pluralcloud/pkg/container/dockerengine"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

func TestDocker_Run(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	config := dockerengine.NewConfig("alpine", strslice.StrSlice{"echo", "Hello, World!"})
	d := dockerengine.NewDocker(ctx, cli, *config)
	resp, err := d.Run(config.Image, config.Cmd)

	assert.NoError(t, err, "Run() returned an unexpected error")
	assert.NotNil(t, resp, "Run() response should not be nil")
}

func TestDocker_Start(t *testing.T) {
	// Setup
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	config := dockerengine.NewConfig("alpine", strslice.StrSlice{"echo", "Hello, World!"})
	d := dockerengine.NewDocker(ctx, cli, *config)
	resp, err := d.Run(config.Image, config.Cmd)

	if err != nil {
		t.Fatalf("Error running container: %v", err)
	}

	err = d.Start(resp.ID)

	assert.NoError(t, err, "Start() returned an unexpected error")
}

func TestDocker_Logs(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	config := dockerengine.NewConfig("alpine", strslice.StrSlice{"echo", "Hello, World!"})
	d := dockerengine.NewDocker(ctx, cli, *config)
	resp, err := d.Run(config.Image, config.Cmd)

	if err != nil {
		t.Fatalf("Error running container: %v", err)
	}
	defer func() {
		_ = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})
	}()

	d.Logs(resp.ID)

}
func TestDocker_List(t *testing.T) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	config := dockerengine.NewConfig("alpine", strslice.StrSlice{"echo", "Hello, World!"})
	d := dockerengine.NewDocker(ctx, cli, *config)
	resp, err := d.Run(config.Image, config.Cmd)

	if err != nil {
		t.Fatalf("Error running container: %v", err)
	}
	defer func() {
		_ = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})
	}()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	assert.NoError(t, err, "ContainerList() returned an unexpected error")
	assert.NotEmpty(t, containers, "ContainerList() should return a non-empty list of containers")
}

func TestDocker_ListAllImages(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	config := dockerengine.NewConfig("alpine", strslice.StrSlice{"echo", "Hello, World!"})
	d := dockerengine.NewDocker(ctx, cli, *config)
	images, err := d.ListAllImages()

	assert.NoError(t, err, "ListAllImages() returned an unexpected error")
	assert.NotEmpty(t, images, "ListAllImages() should return a non-empty list of images")
}

func TestDocker_PullImageWithAuthentication(t *testing.T) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	config := dockerengine.NewConfig("alpine", strslice.StrSlice{"echo", "Hello, World!"})
	d := dockerengine.NewDocker(ctx, cli, *config)

	username := "any_username"
	password := "any_password"
	imageName := "alpine:latest"

	err = d.PullImageWithAuthentication(imageName, username, password)
	assert.NoError(t, err, "PullImageWithAuthentication() returned an unexpected error")

	images, err := d.ListAllImages()
	assert.NoError(t, err, "ListAllImages() returned an unexpected error")
	assert.NotEmpty(t, images, "ListAllImages() should return a non-empty list of images")
	imagePulled := false
	for _, img := range images {
		if img.RepoTags != nil {
			for _, tag := range img.RepoTags {
				if tag == imageName {
					imagePulled = true
					break
				}
			}
		}
	}
	assert.True(t, imagePulled, "Image was not pulled successfully")
}
