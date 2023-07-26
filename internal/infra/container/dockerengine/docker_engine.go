package dockerengine

import (
	"github/angolarti/pluralcloud/internal/entity"
	"github/angolarti/pluralcloud/pkg/container"
)

type DockerEngine struct {
	Docker *container.Docker
}

func NewDockerEngine(docker *container.Docker) *DockerEngine {
	return &DockerEngine{
		Docker: docker,
	}
}

func (e *DockerEngine) Run(c *entity.Container) (string, error) {
	resp, err := e.Docker.Run(c.Image, c.Command)

	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (e *DockerEngine) List(c *entity.Container) ([]string, error) {
	return []string{""}, nil
}
