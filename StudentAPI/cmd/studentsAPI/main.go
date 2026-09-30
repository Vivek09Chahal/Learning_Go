package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/Vivek09Chahal/studentsAPI/internal/config"
	"github.com/Vivek09Chahal/studentsAPI/internal/http/handlers/student"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()
	
	router.HandleFunc("POST /api/students", student.New())
	
	// setup server
	server := http.Server { 
	    Addr: cfg.Addr,
		Handler:  router,
	}
	slog.Info("server started", "address", cfg.HTTPServer.Addr)
	fmt.Printf("server started %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
    	err := server.ListenAndServe()
    	if err  != nil {
    	    log.Fatal("failed to start serve r")
    	}
	} ()

	<- done

	slog.Info("shutting down server")
	 
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := server.Shutdown(ctx)

	if err !=  nil {
	    slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shut down successfully")
	
}
