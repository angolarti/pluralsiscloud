package controllers

import (
	"context"
	"encoding/json"
	"github/angolarti/pluralcloud/internal/usecase"
	de "github/angolarti/pluralcloud/pkg/container/dockerengine"

	"net/http"

	"github.com/docker/docker/client"
)

func ListContainer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response []usecase.ContainerOutput

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}
	config := de.NewConfig("", []string{"tail", "-f", "/dev/null"})
	docker := de.NewDocker(ctx, cli, *config)

	containers, _ := docker.List()

	for _, container := range containers {

		var privatePort int
		var publicPort int

		if len(container.Ports) > 0 {
			privatePort = int(container.Ports[0].PrivatePort)
			publicPort = int(container.Ports[0].PublicPort)
		}

		output := &usecase.ContainerOutput{
			ContainerID: container.ID,
			Image:       container.Image,
			PrivatePort: privatePort,
			PublicPort:  publicPort,
			Name:        container.Names[0],
			Command:     config.Cmd,
		}

		response = append(response, *output)
	}
	json.NewEncoder(w).Encode(response)
}
