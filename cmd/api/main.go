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
	"github.com/krishkumar84/golang-project/pkg/http/handler/users"
	"github.com/krishkumar84/golang-project/pkg/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// load config

	cfg := config.MustLoad()


	//database
    storage,err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Database connected",cfg.Env)

	//setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello krish this side  and server is up and running"))
	})

	router.HandleFunc("POST /api/users",users.New(storage))
    // router.HandleFunc("GET /api/users/{id}",users.GetUser(storage))
	// router.HandleFunc("GET /api/users",users.GetAll(storage))
    
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