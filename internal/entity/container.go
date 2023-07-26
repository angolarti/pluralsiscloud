package entity

import (
	"errors"
)

type Container struct {
	ContainerID string
	Image       string
	Created     string
	Status      string
	Port        int
	Name        string
	Command     []string
}

func NewContainer(image string, name string, port int, command []string) *Container {

	return &Container{
		Image:   image,
		Name:    name,
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
