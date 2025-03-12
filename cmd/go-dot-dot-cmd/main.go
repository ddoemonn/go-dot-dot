package go_dot_dot_cmd

import (
	"fmt"
	"log"

	"github.com/ddoemonn/go-dot-dot/internal/app"
	"github.com/ddoemonn/go-dot-dot/internal/config"
)

// Execute runs the application
func Execute() {
    // Load configuration
    cfg, err := config.Load()
    if err != nil {
        log.Printf("Warning: %v", err)
        fmt.Println("Continuing with default configuration.")
    }

    // Initialize and run the application
    application, err := app.New(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize application: %v", err)
    }

    if err := application.Run(); err != nil {
        log.Fatalf("Error running program: %v", err)
    }
}