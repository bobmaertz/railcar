package main

import (
	"flag"
	"fmt"
	"log"
    "log/slog"
	"net/http"

    "os"
	"github.com/bobmaertz/railcar/cmd/server/actions"
	"github.com/bobmaertz/railcar/pkg/config"
)

var (
	configFile = flag.String("config_file", "", "The path to the config file")
)

func main() {

	flag.Parse()

	conf, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    slog.SetDefault(logger)

	handler, err := actions.App(conf)
	if err != nil {
		log.Fatal(err)
	}

	if conf.Pprof != nil {
		go func() {
			log.Printf("Starting `pprof` at port %v", conf.Pprof.Port)
			log.Fatal(http.ListenAndServe(fmt.Sprint(conf.Pprof.Port), nil))
		}()
	}

    slog.Info("Starting server at port 3333")
	log.Fatal(http.ListenAndServe(":3333", handler))

}
