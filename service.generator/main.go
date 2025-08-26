package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vhall1/password-generator/service.generator/handler"
)

func main() {
	router := handler.NewRouter()
	router.SetupRoutes()

	fmt.Println("Starting server on :8080")

	go func() {
		if err := router.Start(); err != nil && err.Error() != "http: Server closed" {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}
