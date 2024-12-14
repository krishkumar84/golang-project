package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/krishkumar84/golang-project/pkg/config"
)

func main() {
	// load config

	cfg := config.MustLoad()

	//database


	//setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! hello krish"))
	})
    
	//start server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

    fmt.Println("Server is running on port", cfg.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Hello, World!")
}