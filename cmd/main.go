package main

import (
	"log"

	todo "example.com/m/v2"
	"example.com/m/v2/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured %s", err.Error())
	}
}
