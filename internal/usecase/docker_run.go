package usecase

import (
	"github/angolarti/pluralcloud/internal/entity"
)

type ContainerInput struct {
	ImageName string   `json:"image"`
	Name      string   `json:"name"`
	Port      int      `json:"port"`
	Command   []string `json:"command"`
}

type ContainerOutput struct {
	ContainerID string
	Image       string
	Port        int
	Name        string
	Command     []string
}

type DockerRun struct {
	Container entity.ContainerInterface // implement
}

func NewDockerRun(container entity.ContainerInterface) *DockerRun {
	return &DockerRun{
		Container: container,
	}
}

func (dc *DockerRun) Execute(input ContainerInput) (*ContainerOutput, error) {

	container := entity.NewContainer(input.ImageName, input.Name, input.Port, input.Command)
	containrID, err := dc.Container.Run(container)
	if err != nil {
		return nil, err
	}

	return &ContainerOutput{
		ContainerID: containrID,
		Image:       container.Image,
		Port:        container.Port,
		Name:        container.Name,
		Command:     container.Command,
	}, nil
}
