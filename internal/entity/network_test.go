package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainerNetwork(t *testing.T) {
	name := "my_network"
	driver := "bridge"
	scope := "local"

	network := NewContainerNetwork(name, driver, scope)

	assert.Equal(t, name, network.Name)
	assert.Equal(t, driver, network.Driver)
	assert.Equal(t, scope, network.Scope)
}

func TestContainerNetwork_Validate(t *testing.T) {
	name := "my_network"
	driver := "bridge"
	scope := "local"

	network := NewContainerNetwork(name, driver, scope)

	err := network.Validate()
	assert.Nil(t, err)
}

func TestContainerNetwork_Validate_Invalid(t *testing.T) {
	driver := "bridge"
	scope := "local"

	network := NewContainerNetwork("", driver, scope)

	err := network.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())

	network = NewContainerNetwork("my_network", "", scope)

	err = network.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "driver is required", err.Error())

	network = NewContainerNetwork("my_network", driver, "")

	err = network.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "scope is required", err.Error())
}
