# go-dot-dot

A PostgreSQL database explorer TUI (Terminal User Interface) application written in Go.

## Features

- Browse database tables in an interactive terminal interface
- View table data with pagination
- Search and filter table contents
- Detailed row view for examining specific records
- Keyboard-driven navigation with intuitive shortcuts

## Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL database

### From Source

```bash
# Clone the repository
git clone https://github.com/ddoemonn/go-dot-dot.git
cd go-dot-dot

# Build the application
go build -o go-dot-dot

# Run the application
./go-dot-dot
```

### Using Go Install

```bash
go install github.com/ddoemonn/go-dot-dot@latest
```

## Configuration

The application uses environment variables for configuration. You can set these in your environment or create a `.env` file in the project root:

```
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_HOST=localhost
DB_PORT=5432
```

## Usage

After starting the application, you'll see a list of tables in your database. 

### Key Bindings

- `↑/↓`: Navigate through tables or rows
- `Enter`: Select a table or view row details
- `/`: Enter search mode
- `Esc`: Exit search mode or return to previous view
- `q`: Quit the application
- `?`: Toggle help view

## Development

### Project Structure

```
go-dot-dot/
├── cmd/
│   └── go-dot-dot-cmd/     # Command implementation
├── internal/
│   ├── app/                # Application logic
│   ├── config/             # Configuration handling
│   ├── db/                 # Database interactions
│   ├── model/              # Data structures
│   ├── ui/                 # User interface components
│   └── utils/              # Utility functions
├── main.go                 # Entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── .env                    # Environment variables
```

### Building from Source

```bash
go build -o go-dot-dot
```

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
