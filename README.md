# go-dot-dot

A PostgreSQL database explorer TUI (Terminal User Interface) application written in Go.

## Table of Contents

- [go-dot-dot](#go-dot-dot)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Screenshots](#screenshots)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [From Source](#from-source)
  - [Configuration](#configuration)
  - [Usage](#usage)
    - [Key Bindings](#key-bindings)
  - [Project Structure](#project-structure)
  - [License](#license)
  - [Contributing](#contributing)

## Features

- Browse database tables in an interactive terminal interface
- View table data
- Search table contents
- Detailed row view for examining specific records
- Keyboard-driven navigation with intuitive shortcuts

## Screenshots
![Screenshot 2025-03-13 at 15 18 14](https://github.com/user-attachments/assets/68d1f830-3432-4598-bb24-f264c14485da)
![Screenshot 2025-03-13 at 15 18 00](https://github.com/user-attachments/assets/72f4d8e8-c49d-48bf-9046-647c6e901185)
![Screenshot 2025-03-13 at 15 18 08](https://github.com/user-attachments/assets/3695364e-21b9-4ac6-a88c-b931fcbaa52b)
![Screenshot 2025-03-13 at 15 18 23](https://github.com/user-attachments/assets/87fe7ea5-63ea-46a5-9643-624175f1b406)

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

# Move the executable to a directory in your PATH (e.g., /usr/local/bin)
sudo mv go-dot-dot /usr/local/bin

# Run the application
go-dot-dot
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

If no `.env` file exists, the application will automatically prompt you for your database credentials when you run `go-dot-dot`. After entering the credentials, it will generate a `.env` file with your provided information for future use.

## Usage

After starting the application, you'll see a list of tables in your database. 

### Key Bindings

- `↑/↓`: Navigate through tables or rows
- `Enter`: Select a table or view row details
- `/`: Enter search mode
- `Esc`: Exit search mode or return to previous view
- `q`: Quit the application
- `?`: Toggle help view

## Project Structure

```
go-dot-dot/
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

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
