package main

import (
	"log"

	"github.com/secopsninjaaws/e-commerce/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("Erro de conexão com o banco %v", err)
	}
}
