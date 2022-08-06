package storage

type Backend interface {
	GetClient(id string) (Client, error)
	SaveClient(client Client) error
}

type Client struct {
	Id          string
	Name        string
	RedirectUri []string
}
