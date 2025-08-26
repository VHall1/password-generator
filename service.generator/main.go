package main

import (
	"fmt"

	"github.com/vhall1/passsword-generator/service.generator/handler"
)

func main() {
	router := handler.NewRouter()
	router.SetupRoutes()
	fmt.Println("Starting server on :8080")
	router.Start()
}
