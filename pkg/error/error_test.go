package error

import (
	"testing"

	"github.com/bobmaertz/railcar/internal/assert"
)

func TestError(t *testing.T) {
	e := OAuthError{Type: "mock type", Description: "mock description"}

	assert.AssertEqual(t, e.Error(), "type: mock type, description: mock description")
}
