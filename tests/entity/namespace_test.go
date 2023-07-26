package entity

import (
	"github/angolarti/pluralcloud/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNamespace(t *testing.T) {
	cont1 := entity.NewContainer("container1", "image1", 8080, []string{"command1"})
	cont2 := entity.NewContainer("container2", "image2", 9090, []string{"command2"})
	network1 := entity.NewContainerNetwork("network1", "bridge", "local")
	network2 := entity.NewContainerNetwork("network2", "overlay", "swarm")

	n := entity.NewNamespace([]entity.Container{*cont1, *cont2}, []entity.ContainerNetwork{*network1, *network2})

	assert.NotNil(t, n)
	assert.Equal(t, 2, len(n.Containers))
	assert.Equal(t, 2, len(n.Networks))
}

func TestNewNamespace_Empty(t *testing.T) {
	n := entity.NewNamespace([]entity.Container{}, []entity.ContainerNetwork{})

	assert.NotNil(t, n)
	assert.Empty(t, n.Containers)
	assert.Empty(t, n.Networks)
}

func TestNamespace_Validate(t *testing.T) {
	cont1 := entity.NewContainer("container1", "image1", 8080, []string{"command1"})
	cont2 := entity.NewContainer("container2", "image2", 9090, []string{"command2"})
	network1 := entity.NewContainerNetwork("network1", "bridge", "local")

	n := entity.NewNamespace([]entity.Container{*cont1, *cont2}, []entity.ContainerNetwork{*network1})

	err := n.Validate()
	assert.Nil(t, err)
}

func TestNamespace_Validate_Invalid(t *testing.T) {
	n := entity.NewNamespace([]entity.Container{}, []entity.ContainerNetwork{})

	err := n.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "namespace must have at least one container or network", err.Error())
}
