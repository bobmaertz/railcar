package memory

import (
	"errors"
	"time"

	"github.com/bobmaertz/railcar/pkg/storage"
)


var (
    NotImplementedError = errors.New("not implemented")
    NotFoundError = errors.New("not found")
) 
type Memory struct {
	clients []storage.Client
	codes   []storage.AuthorizationCode
	session []storage.Session
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
	return storage.Client{}, NotFoundError 
}

func (m *Memory) SaveClient(client storage.Client) error {
	return NotImplementedError
}

func (m *Memory) CreateAuthorizationCode(code string, client storage.Client, expiry time.Time) error {

	//TODO: ensure this is isolated..
	m.codes = append(m.codes, storage.AuthorizationCode{Code: code, ClientId: client.Id, Expiry: expiry})

	return nil
}

func (m *Memory) CreateSession(userId string, scopes []string, client storage.Client, expiry time.Time) error {

	//TODO: ensure this is isolated..
	m.session = append(m.session, storage.Session{UserId: userId, ClientId: client.Id, Expiry: expiry, Scopes: scopes})

	return nil
}
