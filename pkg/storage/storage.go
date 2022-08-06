package storage

import "time"

//TODO: Could make this into a smaller interfaces to build
type Backend interface {
	GetClient(id string) (Client, error)
	SaveClient(client Client) error

	CreateAuthorizationCode(code string, client Client, expiry time.Time) error

	// RemoveClients(id []string) error

	// GetAuthorizationRequest(id string) (AuthorizationRequest, error)
	// CreateAuthorizationRequest(id string) error
	// RemoveAuthorizationRequests(id []string) error
}

type Client struct {
	// Unique identifier for client
	Id string
	// Human readable name of client
	Name string
	//Allowed redirect URI's
	RedirectUris []string
}

type AuthorizationRequest struct {
	// Unique identifier authentication request
	Id string
	// Note: this could and should be a unique identifier but its supplied by the client so
	// adding it as a seperate field to prevent any future issues from clients.
	State        string
	ClientId     string
	ResponseType string
	Scope        []string
	//RedirectUri used as part of the request.
	RedirectUri string
}

type AuthorizationCode struct {
	// Unique identifier authentication request
	// Id       string
	Code     string
	ClientId string
	Expiry   time.Time
}
