package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giridharan-7/home-energy-monitor/internal/handlers"
	"github.com/giridharan-7/home-energy-monitor/pkg/config"
	"github.com/giridharan-7/home-energy-monitor/pkg/database"
	"github.com/giridharan-7/home-energy-monitor/web/routes"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// energyRepo := database.NewEnergyRepository(db)
	energyHandler := handlers.NewEnergyHandler()
	router := mux.NewRouter()
	routes.SetupEnergyRoutes(router, energyHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	log.Printf("Server listening on port %d", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
