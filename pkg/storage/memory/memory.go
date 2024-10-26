package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/bobmaertz/railcar/pkg/storage"
)

var (
	NotImplementedError = errors.New("not implemented")
	NotFoundError       = errors.New("not found")
)

type Memory struct {
	clients map[string]storage.Client
	codes   []storage.AuthorizationCode
	session []storage.Session

	mu sync.Mutex
}

// NewMemory creates a new memory structure
func NewMemory() (*Memory, error) {
	return &Memory{
		clients: map[string]storage.Client{
			"29352735982374239857": {
				Name:         "Mock Test Client",
				RedirectUris: []string{"http://localhost"},
			},
		},
	}, nil
}

func (m *Memory) GetClient(id string) (storage.Client, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.clients {
		if k == id {
			v.Id = k
			return v, nil
		}
	}
	return storage.Client{}, NotFoundError
}

func (m *Memory) SaveClient(client storage.Client) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[client.Id] = client

	return nil
}

func (m *Memory) CreateAuthorizationCode(code string, client storage.Client, expiry time.Time) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.codes = append(m.codes, storage.AuthorizationCode{Code: code, ClientId: client.Id, Expiry: expiry})

	return nil
}

func (m *Memory) CreateSession(userId string, scopes []string, client storage.Client, expiry time.Time) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.session = append(m.session, storage.Session{UserId: userId, ClientId: client.Id, Expiry: expiry, Scopes: scopes})

	return nil
}
