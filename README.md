# Shorter - A URL Shortening Service

A minimalist URL shortening service built with Go-Zero framework that allows you to create, manage and use shortened URLs.

## Features

- Shorten long URLs to compact, easy-to-share links
- Redirect from shortened URLs to original destinations
- Health check endpoint
- Redis caching for improved performance
- MySQL database storage for persistence

## Technology Stack

- [Go-Zero](https://github.com/zeromicro/go-zero) - A web and RPC framework with many built-in engineering best practices
- MySQL - For storing URL mappings
- Redis - For caching frequent requests
- Docker & Docker Compose - For easy deployment and development

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/short` | POST | Create a shortened URL |
| `/s/:shortcode` | GET | Redirect to the original URL |
| `/healthz` | GET | Health check endpoint |

## Getting Started

### Prerequisites

- Go 1.16+
- Docker and Docker Compose (for local development)

### Installation & Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/shorter.git
cd shorter
```

2. Start the required services (MySQL and Redis) using Docker Compose:
```bash
docker-compose up -d
```

3. Run the application:
```bash
go run shorter.go
```

The service will start on port 8888 (by default).

### Configuration

The configuration is stored in `etc/shorter-api.yaml`. You can modify the following parameters:

- Host and Port for the API server
- Database connection string
- Redis cache settings
- Log level

## Usage Examples

### Shorten a URL
```bash
curl -X POST http://localhost:8888/api/short -H "Content-Type: application/json" -d '{"url": "https://example.com/very-long-url-that-needs-shortening"}'
```

Response:
```json
{
  "shortened": "http://localhost:8888/s/abc123"
}
```

### Use a shortened URL
Simply open the shortened URL in your browser or use:
```bash
curl -L http://localhost:8888/s/abc123
```

## Project Structure

```
├── api                 # API definitions
│   ├── healthz.api     # Health check API definition
│   └── shorter.api     # URL shortener API definition
├── etc                 # Configuration files
├── internal            
│   ├── config          # Configuration structures
│   ├── handler         # HTTP handlers
│   ├── logic           # Business logic
│   ├── model           # Database models
│   ├── svc             # Service context
│   └── types           # Type definitions
├── pkg
│   └── utils           # Utility functions
└── shorter.go          # Main entry point
```

## Development

### Generating Code From API Definition

The project uses go-zero's tool to generate code from API definitions:

```bash
goctl api go -api ./api/shorter.api -dir .
```

### Creating the Database

The database schema is defined in `internal/model/shorter.sql` and is automatically created when running with Docker Compose.

## License

This project is licensed under the [MIT License](LICENSE).
