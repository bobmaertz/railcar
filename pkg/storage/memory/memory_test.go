package memory

import (
	"testing"

	"github.com/bobmaertz/railcar/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestGetClient(t *testing.T) {

	m, err := NewMemory()
	if err != nil {
		t.Error(err)
	}

	actualClient, e := m.GetClient("abcd")

	expectedClient := storage.Client{Id: "abcd", Name: "Mock Test Client", RedirectUris: []string{"http://localhost"}}
	assert.Nil(t, e)
	assert.Equal(t, expectedClient, actualClient)

}

func TestSaveClient(t *testing.T) {

	m, err := NewMemory()
	if err != nil {
		t.Error(err)
	}
	c := storage.Client{Id: "defg"}

	e := m.SaveClient(c)

	assert.EqualError(t, e, "not implemented")
}
