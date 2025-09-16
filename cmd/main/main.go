package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tudemaha/tujuhin-be/pkg/controller"
	"github.com/tudemaha/tujuhin-be/pkg/database"
	"github.com/tudemaha/tujuhin-be/pkg/server"
)

func main() {
	log.Println("INFO load env: loading .env")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ERROR load env fatal error: %v", err)
	}

	httpAddress := os.Getenv("ADDRESS")
	httpServer := server.NewServer()
	db := database.GetDatabase()

	log.Println("INFO controllers: initializing controllers")
	controller.InitializeControllers(httpServer.Router, db)

	if err := httpServer.Start(httpAddress); err != nil {
		log.Fatalf("ERROR start server fatal error: %v", err)
	}

}
