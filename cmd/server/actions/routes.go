package actions

import (
	"net/http"

	"github.com/bobmaertz/railcar/pkg/storage"
)

func defineRoutes(m *http.ServeMux, s storage.Backend) {
	m.HandleFunc("/version", versionHandler)

	m.HandleFunc("/authorize", authorizeHandler(s))
	m.HandleFunc("/token", tokenHandler(s))
}
