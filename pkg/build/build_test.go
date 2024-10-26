package build

import (
	"testing"

	"github.com/bobmaertz/railcar/pkg/internal/assert"
)

func TestError(t *testing.T) {
	out := BuildInfo()

	assert.AssertEqual(t, out, "sha1: , buildTime: ")
}
