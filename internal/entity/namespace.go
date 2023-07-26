package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Namespace struct {
	ID         uuid.UUID
	Containers []Container
	Networks   []ContainerNetwork
}

func NewNamespace(containers []Container, networks []ContainerNetwork) *Namespace {
	return &Namespace{
		Containers: containers,
		Networks:   networks,
	}
}

func (n *Namespace) Validate() error {
	if len(n.Containers) == 0 && len(n.Networks) == 0 {
		return errors.New("namespace must have at least one container or network")
	}
	return nil
}
