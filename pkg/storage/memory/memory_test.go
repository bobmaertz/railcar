package memory

import (
	"testing"

	"github.com/bobmaertz/railcar/internal/assert"
	"github.com/bobmaertz/railcar/pkg/storage"
)

func TestGetClient(t *testing.T) {
	m, err := NewMemory()
	if err != nil {
		t.Error(err)
	}

	actualClient, e := m.GetClient("29352735982374239857")

	expectedClient := storage.Client{Id: "29352735982374239857", Name: "Mock Test Client", RedirectUris: []string{"http://localhost"}}
	assert.AssertNil(t, e)
	assert.AssertEqual(t, expectedClient.Id, actualClient.Id)
	assert.AssertEqual(t, expectedClient.Name, actualClient.Name)
}

func TestSaveClient(t *testing.T) {
	m, err := NewMemory()
	if err != nil {
		t.Error(err)
	}
	expected := storage.Client{Id: "defg"}

	e := m.SaveClient(expected)
	assert.AssertNil(t, e)

	got, e := m.GetClient(expected.Id)
	assert.AssertNil(t, e)

	assert.AssertEqual(t, expected.Id, got.Id)
	assert.AssertEqual(t, expected.Name, got.Name)
}
