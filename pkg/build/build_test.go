package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	out := BuildInfo()

	assert.Equal(t, out, "sha1: , buildTime: ")
}
