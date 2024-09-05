# go-gin-boilerplate

A robust boilerplate for building web applications with Go and Gin framework.

## Features

- Built with Go and Gin framework
- Database migrations support
- Environment configuration using .env files
- Command-line interface using Cobra
- Structured project layout

## Prerequisites

- Go 1.23 or higher
- PostgreSQL database

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/madevara24/go-gin-boilerplate.git
   cd go-gin-boilerplate
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Set up your environment variables by copying the `.env.example` file to `.env` and modifying it as needed.

4. Run the application:
   ```
   go run main.go server
   ```

## Project Structure

- `cmd/`: Contains the command-line interface commands
- `config/`: Configuration loading and management
- `migrations/`: Database migration files and templates
- `main.go`: Entry point of the application

## Available Commands

- `server`: Start the web server
- `migrate create`: Create a new migration file
- `migrate up`: Run all pending migrations
- `migrate down`: Revert the last applied migration
- `migrate status`: Show the status of all migrations

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.