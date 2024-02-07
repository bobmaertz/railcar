package memory

import (
	"testing"

	"github.com/bobmaertz/railcar/pkg/storage"
	"github.com/bobmaertz/railcar/pkg/internal/assert"
)

func TestGetClient(t *testing.T) {

	m, err := NewMemory()
	if err != nil {
		t.Error(err)
	}

	actualClient, e := m.GetClient("abcd")

	expectedClient := storage.Client{Id: "abcd", Name: "Mock Test Client", RedirectUris: []string{"http://localhost"}}
	assert.AssertNil(t, e)
	assert.AssertEqual(t, expectedClient.Id, actualClient.Id)
	assert.AssertEqual(t, expectedClient.Name, actualClient.Name)
}

func TestSaveClient(t *testing.T) {

	m, err := NewMemory()
	if err != nil {
		t.Error(err)
	}
	c := storage.Client{Id: "defg"}

	e := m.SaveClient(c)


    assert.AssertEqual(t, e, NotImplementedError)
}
