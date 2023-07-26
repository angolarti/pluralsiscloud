package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNamespace(t *testing.T) {
	cont1 := NewContainer("container1", "image1", 8080, []string{"command1"})
	cont2 := NewContainer("container2", "image2", 9090, []string{"command2"})
	network1 := NewContainerNetwork("network1", "bridge", "local")
	network2 := NewContainerNetwork("network2", "overlay", "swarm")

	n := NewNamespace([]Container{*cont1, *cont2}, []ContainerNetwork{*network1, *network2})

	assert.NotNil(t, n)
	assert.Equal(t, 2, len(n.Containers))
	assert.Equal(t, 2, len(n.Networks))
}

func TestNewNamespace_Empty(t *testing.T) {
	n := NewNamespace([]Container{}, []ContainerNetwork{})

	assert.NotNil(t, n)
	assert.Empty(t, n.Containers)
	assert.Empty(t, n.Networks)
}

func TestNamespace_Validate(t *testing.T) {
	cont1 := NewContainer("container1", "image1", 8080, []string{"command1"})
	cont2 := NewContainer("container2", "image2", 9090, []string{"command2"})
	network1 := NewContainerNetwork("network1", "bridge", "local")

	n := NewNamespace([]Container{*cont1, *cont2}, []ContainerNetwork{*network1})

	err := n.Validate()
	assert.Nil(t, err)
}

func TestNamespace_Validate_Invalid(t *testing.T) {
	n := NewNamespace([]Container{}, []ContainerNetwork{})

	err := n.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "namespace must have at least one container or network", err.Error())
}
