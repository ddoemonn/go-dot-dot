package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ddoemonn/go-dot-dot/internal/app"
	"github.com/ddoemonn/go-dot-dot/internal/config"
)

func main() {
	// Check if .env file exists before loading configuration
	envExists := false
	if _, err := os.Stat(".env"); err == nil {
		envExists = true
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		// If setup was cancelled but .env file exists, try to continue
		if strings.Contains(err.Error(), "setup cancelled") && envExists {
			fmt.Println("Configuration exists. Continuing with saved configuration...")
			// Try to load again directly
			cfg, err = config.LoadFromEnv()
			if err != nil {
				log.Fatalf("Failed to load configuration: %v", err)
			}
		} else if strings.Contains(err.Error(), "setup cancelled") {
			fmt.Println("Configuration setup cancelled. Exiting...")
			os.Exit(0)
		} else {
			log.Printf("Warning: %v", err)
			fmt.Println("Continuing with default configuration.")
		}
	}

	// Give a moment for the user to see the success message
	time.Sleep(500 * time.Millisecond)

	// Clear the screen after setup
	fmt.Print("\033[H\033[2J")
	fmt.Println("Starting PostgreSQL Database Explorer...")

	// Initialize and run the application
	application, err := app.New(cfg)
	if err != nil {
		if strings.Contains(err.Error(), "connect") || strings.Contains(err.Error(), "database") {
			fmt.Println("\nError connecting to database. Please check your configuration in .env file.")
			fmt.Println("You can edit the .env file manually or delete it to run the setup wizard again.")
			fmt.Printf("\nError details: %v\n", err)
			os.Exit(1)
		}
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
