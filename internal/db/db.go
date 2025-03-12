package db

import (
	"context"
	"fmt"

	"github.com/ddoemonn/go-dot-dot/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database represents a database connection
type Database struct {
    pool *pgxpool.Pool
}

// Connect establishes a connection to the database
func Connect(cfg *config.DBConfig) (*Database, error) {
    // Connect to the database
    pool, err := pgxpool.New(context.Background(), cfg.ConnectionString())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    return &Database{pool: pool}, nil
}

// Close closes the database connection
func (db *Database) Close() {
    if db.pool != nil {
        db.pool.Close()
    }
}

// GetPool returns the connection pool
func (db *Database) GetPool() *pgxpool.Pool {
    return db.pool
}

// FetchTables retrieves all tables from the database
func (db *Database) FetchTables() ([]string, error) {
    var tables []string
    rows, err := db.pool.Query(context.Background(), `
        SELECT tablename FROM pg_catalog.pg_tables 
        WHERE schemaname NOT IN ('pg_catalog', 'information_schema')
        ORDER BY tablename;
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var tableName string
        err := rows.Scan(&tableName)
        if err != nil {
            return nil, err
        }
        tables = append(tables, tableName)
    }
    return tables, nil
}

// FetchTableData retrieves data from a specific table
func (db *Database) FetchTableData(tableName string) ([][]string, []string, error) {
    query := fmt.Sprintf("SELECT * FROM %s LIMIT 1000", tableName) // Added limit for performance
    rows, err := db.pool.Query(context.Background(), query)
    if err != nil {
        return nil, nil, err
    }
    defer rows.Close()

    // Get column names
    fieldDescriptions := rows.FieldDescriptions()
    columns := make([]string, len(fieldDescriptions))
    for i, fd := range fieldDescriptions {
        columns[i] = string(fd.Name)
    }

    // Fetch rows
    var data [][]string
    for rows.Next() {
        values, err := rows.Values()
        if err != nil {
            return nil, nil, err
        }

        row := make([]string, len(values))
        for i, v := range values {
            if v == nil {
                row[i] = "NULL"
            } else {
                row[i] = fmt.Sprintf("%v", v)
            }
        }
        data = append(data, row)
    }

    return data, columns, nil
}