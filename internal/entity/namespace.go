package entity

import "github.com/google/uuid"

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
