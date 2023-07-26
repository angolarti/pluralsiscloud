package container_test

import (
	"context"
	"github/angolarti/pluralcloud/pkg/container"
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

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})

	resp, err := d.Run()

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

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})
	resp, err := d.Run()
	if err != nil {
		t.Fatalf("Error running container: %v", err)
	}

	err = d.Start(resp.ID)

	assert.NoError(t, err, "Start() returned an unexpected error")
}

func TestDocker_Validate(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})

	err = d.Validate()

	assert.NoError(t, err, "Validate() returned an unexpected error")
}
func TestDocker_Logs(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatalf("Error creating Docker client: %v", err)
	}
	defer cli.Close()

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})
	resp, err := d.Run()
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

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})
	resp, err := d.Run()
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

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})

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

	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})

	username := "any_username"
	password := "any_password"
	imageName := "alpine:latest"

	err = d.PullImageWithAuthentication(ctx, imageName, username, password)
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

// func TestDocker_CommitAcontainer(t *testing.T) {

// 	ctx := context.Background()
// 	cli, err := client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		t.Fatalf("Error creating Docker client: %v", err)
// 	}
// 	defer cli.Close()

// 	d := container.NewDocker(ctx, cli, "alpine", strslice.StrSlice{"echo", "Hello, World!"})

// 	resp, err := d.Run()
// 	if err != nil {
// 		t.Fatalf("Error running container: %v", err)
// 	}
// 	defer func() {
// 		_ = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true})
// 	}()

// 	err = d.CommitAcontainer(ctx, "helloworld", "cunga22", "X$PtV*2'?XAEfgw")
// 	assert.NoError(t, err, "CommitAcontainer() returned an unexpected error")

// 	images, err := d.ListAllImages()
// 	assert.NoError(t, err, "ListAllImages() returned an unexpected error")
// 	assert.NotEmpty(t, images, "ListAllImages() should return a non-empty list of images")
// 	imageCreated := false
// 	for _, img := range images {
// 		if img.RepoTags != nil {
// 			for _, tag := range img.RepoTags {
// 				if tag == "helloworld" {
// 					imageCreated = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	assert.True(t, imageCreated, "New image was not created successfully")
// }
