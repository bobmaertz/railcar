package memory

import (
	"errors"
	"time"

	"github.com/bobmaertz/railcar/pkg/storage"
)

type Memory struct {
	clients []storage.Client
	codes   []storage.AuthorizationCode
}

func NewMemory() (*Memory, error) {
	return &Memory{clients: []storage.Client{{Id: "abcd", Name: "Mock Test Client", RedirectUris: []string{"http://localhost"}}}}, nil
}

func (m *Memory) GetClient(id string) (storage.Client, error) {
	for _, v := range m.clients {
		if v.Id == id {
			return v, nil
		}
	}
	return storage.Client{}, errors.New("not found")
}

func (m *Memory) SaveClient(client storage.Client) error {
	return errors.New("not implemented")
}

func (m *Memory) CreateAuthorizationCode(code string, client storage.Client, expiry time.Time) error {

	//TODO: ensure this is isolated..
	m.codes = append(m.codes, storage.AuthorizationCode{Code: code, ClientId: client.Id, Expiry: expiry})

	return nil
}
