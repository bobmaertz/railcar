package error

import "fmt"

type OAuthError struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (e OAuthError) Error() string {
	return fmt.Sprintf("type: %v, description: %v", e.Type, e.Description)
}
