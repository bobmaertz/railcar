package actions

import (
	"net/http"

	"github.com/bobmaertz/railcar/pkg/config"
	"github.com/bobmaertz/railcar/pkg/storage/memory"
)

func App(conf config.Config) (http.Handler, error) {

	mux := http.DefaultServeMux

	s, _ := memory.NewMemory()
	defineRoutes(mux, s)

	return mux, nil
}
