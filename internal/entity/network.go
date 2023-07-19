package entity

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
