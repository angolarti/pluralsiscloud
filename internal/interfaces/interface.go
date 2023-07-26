package entity

import "github/angolarti/pluralcloud/internal/entity"

type ContainerInterface interface {
	Create(oc *entity.Container) error
	List(c *entity.Container) ([]string, error)
}
