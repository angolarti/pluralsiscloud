package entity

type ContainerInterface interface {
	Run(c *Container) (string, error)
	List(c *Container) ([]string, error)
}
