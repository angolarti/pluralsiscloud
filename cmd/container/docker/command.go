package dockerk

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"github/angolarti/pluralcloud/pkg/container"
// 	"os"

// 	"github.com/docker/docker/client"
// )

// func DockerCommand() {

// 	ctx := context.Background()
// 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

// 	if err != nil {
// 		panic(err)
// 	}

// 	provider := flag.NewFlagSet("provider", flag.ExitOnError)
// 	image := provider.String("image", "docker.io/library/alpine:latest", "Image name")

// 	c // create Slice
// 	flag.Parse()

// 	fmt.Println(os.Args[1])

// 	switch os.Args[1] {
// 	case "provider":
// 		provider.Parse(os.Args[2:])
// 		dc := container.NewDocker(*image, command)
// 		dc.ImageName = *image
// 		fmt.Println(dc.ImageName)
// 		resp, err := dc.Run(ctx, cli)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Container com ID", resp.ID)
// 	default:
// 		fmt.Println("expected 'provider' subcommands")
// 		os.Exit(1)
// 	}

// 	// for _, container := range containres {
// 	// 	fmt.Println("Container ID: ", container.ID, "Names: ", container.Names[0], "Status: ", container.Status)
// 	// 	// dc.Logs(container.ID)
// 	// }

// 	// images, err := dc.ListAllImages()

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// for _, image := range images {
// 	// 	fmt.Println(image.ID)
// 	// }

// }
