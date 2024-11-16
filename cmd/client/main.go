package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type AuthCodeRequest struct {
	Code  string
	Error string
}

func main() {
	mux := http.DefaultServeMux

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			params := r.URL.Query()

			e := params.Get("error")
			code := params.Get("code")

			acr := AuthCodeRequest{
				Code:  code,
				Error: e,
			}
			err = tmpl.Execute(w, acr)
			if err != nil {
				fmt.Println("template error execution: %v", err)
				return
			}

		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
