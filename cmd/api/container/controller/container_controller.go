package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github/angolarti/pluralcloud/internal/infra/container/dockerengine"
	"github/angolarti/pluralcloud/internal/usecase"
	de "github/angolarti/pluralcloud/pkg/container/dockerengine"

	"github.com/docker/docker/client"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	defer r.Body.Close()

	var input usecase.ContainerInput
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		http.Error(w, "Failed to create Docker client", http.StatusInternalServerError)
		return
	}

	config := de.NewConfig(input.ImageName, []string{"tail", "-f", "/dev/null"})
	docker := de.NewDocker(ctx, cli, *config)
	dockerEngine := dockerengine.NewDockerEngine(docker)
	us := usecase.NewDockerRunUsecase(dockerEngine)

	output, err := us.Execute(input)
	if err != nil {
		http.Error(w, "Failed to execute Docker run use case", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(output)
}

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
