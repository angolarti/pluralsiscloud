package entity

import "errors"

type ContainerNetwork struct {
	NetworkID string
	Name      string
	Driver    string
	Scope     string
}

func NewContainerNetwork(name string, driver string, scope string) *ContainerNetwork {
	return &ContainerNetwork{
		Name:   name,
		Driver: driver,
		Scope:  scope,
	}
}

func (cn *ContainerNetwork) Validate() error {
	if cn.Name == "" {
		return errors.New("name is required")
	}
	if cn.Driver == "" {
		return errors.New("driver is required")
	}
	if cn.Scope == "" {
		return errors.New("scope is required")
	}
	return nil
}
