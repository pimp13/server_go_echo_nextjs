package main

import (
	"backend_next_echo/internal/server"
	"backend_next_echo/pkg/config"
	"log"
)

func init() {
	if err := config.Initialize(); err != nil {
		log.Fatal("error in not exists .env file:", err.Error())
	}
}

func main() {
	if err := server.ConnectDb(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer server.CloseDb()

	apiServer := server.NewAPI(":8070", nil)
	apiServer.Run()
}
