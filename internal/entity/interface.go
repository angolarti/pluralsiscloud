package entity

type ContainerInterface interface {
	Create(oc *Container) error
	List(c *Container) ([]string, error)
}
