package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Container struct {
	ID          uuid.UUID
	ContainerID string
	Image       string
	Created     string
	Status      string
	Port        int
	Name        string
	Command     []string
}

func NewContainer(name string, image string, port int, command []string) *Container {

	return &Container{
		Name:    name,
		Image:   image,
		Port:    port,
		Command: command,
	}
}

func (c *Container) Validate() error {

	if c.Image == "" {
		return errors.New("image is required")
	}

	return nil
}
