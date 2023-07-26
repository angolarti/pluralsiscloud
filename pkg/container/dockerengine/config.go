package dockerengine

type Config struct {
	Image string
	Cmd   []string
}

func NewConfig(image string, command []string) *Config {
	return &Config{
		Image: image,
		Cmd:   command,
	}
}
