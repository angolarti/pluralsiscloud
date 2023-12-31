package main

import (
	"context"
	"fmt"

	"github/angolarti/pluralcloud/internal/infra/container/dockerengine"
	"github/angolarti/pluralcloud/internal/usecase"
	de "github/angolarti/pluralcloud/pkg/container/dockerengine"

	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	config := de.NewConfig("httpd", []string{"tail", "-f", "/dev/null"})

	if err != nil {
		panic(err)
	}

	docker := de.NewDocker(ctx, cli, *config)
	dockerEngine := dockerengine.NewDockerEngine(docker)
	us := usecase.NewDockerRun(dockerEngine)

	var input usecase.ContainerInput

	input.ImageName = config.Image
	input.Command = config.Cmd
	input.Name = "apache"
	input.Port = 80

	output, err := us.Execute(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)

}
