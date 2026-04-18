package main

import (
	"log"

	"github.com/CookingApps/drone-telemetry-simulator-go/internal/server"
)

func main() {
	log.Println("🚁 Naija Drone Telemetry Simulator (Golang) Starting...")
	srv := server.NewServer()
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
