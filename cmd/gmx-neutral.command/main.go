package main

import (
	"log"

	grpc_server "github.com/gmx-delta-neutral/gmx-neutral.command/internal/server/grpc"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	grpc_server.StartServer()
}
