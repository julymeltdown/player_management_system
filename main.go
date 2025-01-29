package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"player_management_system/config"
	playerHttpHandler "player_management_system/internal/handlers/http"
	platformPostgres "player_management_system/internal/platform/postgres"
	"player_management_system/internal/repositories/player/postgres"
	"player_management_system/internal/services/player"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Database configuration
	dbConfig := platformPostgres.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	}

	// Connect to the database
	db, err := platformPostgres.New(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repository, service, and handler
	playerRepo := postgres.NewPlayerRepository(db)
	playerService := player.NewPlayerService(playerRepo)
	playerHandler := playerHttpHandler.NewPlayerHandler(playerService)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	playerHandler.RegisterRoutes(e)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}
