package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerImagePull(t *testing.T) {
	docker := Docker{}

	// if docker.Validate() == nil {
	// 	t.Error("image name is required")
	// }
	assert.Error(t, docker.Validate(), "image name is required")
}
