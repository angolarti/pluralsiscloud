package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	name := "my_container"
	image := "my_image"
	port := 8080
	command := []string{"echo", "hello world"}

	container := NewContainer(name, image, port, command)

	assert.Equal(t, name, container.Name)
	assert.Equal(t, image, container.Image)
	assert.Equal(t, port, container.Port)
	assert.Equal(t, command, container.Command)
}

func TestContainer_Validate(t *testing.T) {
	name := "my_container"
	image := "my_image"
	port := 8080
	command := []string{"echo", "hello world"}

	container := NewContainer(name, image, port, command)

	err := container.Validate()
	assert.Nil(t, err)
}

func TestContainer_Validate_Invalid(t *testing.T) {
	name := "my_container"
	port := 8080
	command := []string{"echo", "hello world"}

	container := NewContainer(name, "", port, command)

	err := container.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "image is required", err.Error())
}
