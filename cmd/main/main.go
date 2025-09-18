package main

import (
	"fmt"
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

	httpPort := os.Getenv("PORT")
	httpServer := server.NewServer()
	db := database.GetDatabase()

	log.Println("INFO controllers: initializing controllers")
	controller.InitializeControllers(httpServer.Router, db)

	if err := httpServer.Start(fmt.Sprintf(":%s", httpPort)); err != nil {
		log.Fatalf("ERROR start server fatal error: %v", err)
	}

}
