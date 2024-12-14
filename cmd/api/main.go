package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krishkumar84/golang-project/pkg/config"
)

func main() {
	// load config

	cfg := config.MustLoad()

	//database


	//setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! hello m, krishkhush"))
	})
    
	//start server

	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

    fmt.Println("Server is running on port", cfg.Addr)
  
     done := make(chan os.Signal,1)

	 signal.Notify(done, os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	<-done

	slog.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx) ; err != nil {

		slog.Error("Server Shutdown Failed",slog.String("error",err.Error()))
	}

	slog.Info("Server ShutDown Properly")
}