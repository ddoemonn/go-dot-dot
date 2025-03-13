package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
    DB DBConfig
}

// DBConfig holds database connection parameters
type DBConfig struct {
    User     string
    Password string
    Name     string
    Host     string
    Port     string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
    // Check if .env file exists
    if _, err := os.Stat(".env"); os.IsNotExist(err) {
        fmt.Println("\nNo .env file found. Starting configuration wizard...")
        fmt.Println("Please enter your PostgreSQL database connection details.")
        fmt.Println("These will be saved to a .env file in the current directory.\n")
        return RunSetup()
    }

    // Load .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("\nError loading .env file. Starting configuration wizard...")
        fmt.Println("Please enter your PostgreSQL database connection details.")
        fmt.Println("These will be saved to a .env file in the current directory.\n")
        return RunSetup()
    }

    // Check if required environment variables are set
    if os.Getenv("DB_USER") == "" && os.Getenv("DB_HOST") == "" {
        fmt.Println("\nInvalid or empty .env file. Starting configuration wizard...")
        fmt.Println("Please enter your PostgreSQL database connection details.")
        fmt.Println("These will be saved to a .env file in the current directory.\n")
        return RunSetup()
    }

    // Read DB credentials from .env or environment
    cfg := &Config{
        DB: DBConfig{
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            Name:     getEnv("DB_NAME", "postgres"),
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnv("DB_PORT", "5432"),
        },
    }

    return cfg, nil
}

// LoadFromEnv loads configuration directly from the .env file
func LoadFromEnv() (*Config, error) {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("error loading .env file: %w", err)
    }

    // Read DB credentials from .env or environment
    cfg := &Config{
        DB: DBConfig{
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            Name:     getEnv("DB_NAME", "postgres"),
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnv("DB_PORT", "5432"),
        },
    }

    return cfg, nil
}

// ConnectionString returns a formatted connection string
func (c *DBConfig) ConnectionString() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
        c.User, c.Password, c.Host, c.Port, c.Name)
}

// ConnectionDetails returns a user-friendly connection string for display
func (c *DBConfig) ConnectionDetails() string {
    return fmt.Sprintf("Connected to: %s@%s:%s/%s", 
        c.User, c.Host, c.Port, c.Name)
}

