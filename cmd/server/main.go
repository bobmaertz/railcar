package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe(":3333", handler))

}
