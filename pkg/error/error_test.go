package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	e := OAuthError{Type: "mock type", Description: "mock description"}

	assert.Equal(t, e.Error(), "type: mock type, description: mock description")
}
