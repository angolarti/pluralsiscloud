package main

import (
	"context"
	"encoding/json"
	"github/angolarti/pluralcloud/internal/infra/container/dockerengine"
	"github/angolarti/pluralcloud/internal/usecase"
	de "github/angolarti/pluralcloud/pkg/container/dockerengine"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/container", CreateContainer)
	router.Get("/container", ListContainer)
	http.ListenAndServe(":8888", router)
}

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	defer r.Body.Close()
	var input usecase.ContainerInput

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.Unmarshal(body, &input)

	if err != nil {
		panic(err)
	}
	json.Marshal(input)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	config := de.NewConfig(input.ImageName, []string{"tail", "-f", "/dev/null"})

	if err != nil {
		panic(err)
	}

	docker := de.NewDocker(ctx, cli, *config)
	dockerEngine := dockerengine.NewDockerEngine(docker)
	us := usecase.NewDockerRun(dockerEngine)

	output, err := us.Execute(input)

	if err != nil {
		panic(err)
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
