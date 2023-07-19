package main

import (
	"context"
	"fmt"

	"github/angolarti/pluralcloud/pkg/container"

	"github.com/docker/docker/client"
)

func main() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	const image = "docker.io/library/nginx"
	command := []string{"echo", "hello world"}

	dc := container.NewDocker(ctx, cli, image, command)
	dc.ImageName = image

	// resp, err := dc.Run()
	// containres, err := dc.List()

	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println("Container com ID", resp.ID)

	// for _, container := range containres {
	// 	fmt.Println("Container ID: ", container.ID, "Names: ", container.Names[0], "Status: ", container.Status)
	// 	dc.Logs(container.ID)
	// }

	images, err := dc.ListAllImages()

	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}

}
