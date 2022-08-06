package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	e := OAuthError{Type: "mock type", Description: "mock description"}

	assert.EqualError(t, e, "type: mock type, description: mock description")
}
