package memory

import (
	"errors"

	"github.com/bobmaertz/railcar/pkg/storage"
)

type Memory struct {
	clients []storage.Client
}

func NewMemory() (Memory, error) {
	return Memory{clients: []storage.Client{{Id: "abcd"}}}, nil
}

func (m Memory) GetClient(id string) (storage.Client, error) {
	for _, v := range m.clients {
		if v.Id == id {
			return v, nil
		}
	}
	return storage.Client{}, errors.New("not found")
}

func (m Memory) SaveClient(client storage.Client) error {
	return errors.New("not implemented")
}
