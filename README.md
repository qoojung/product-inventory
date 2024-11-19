# Product Inventory

## Prerequisites

### Required Software (Tested)
- Ubuntu 24.04
- Docker 26.1.4
- Docker Compose v2.27.1
- Go 1.23.3 (for local development)
- GNU Make 4.3 (optional, for using Makefile commands)

### Environment Setup

1. Clone the repository
```bash
git clone <repository-url>
cd <project-directory>
```
## Running with Docker Compose

1. Build and start the containers:
```bash
docker-compose up -d --build
```

2. Stop the containers:
```bash
docker-compose down
```

## Development

### Local Setup

1. Install dependencies:
```bash
go mod download
```

2. (**Optional for local development**) Create a `.env` file in the root directory with the following variables:
```env
# Database Configuration
POSTGRES_HOST=pgdb
POSTGRES_PORT=5432
POSTGRES_USER=simple_pg
POSTGRES_PASSWORD=simple_pass
POSTGRES_DB=simple_pg_db
```


3. Run the application:
```bash
go run main.go
```
### Run Test
- run `make test` 

### Environment Variables
The application uses environment variables for configuration. You can:
- Create a `.env` file in the project root (recommended for development)
- Set environment variables directly in your system
- Use Docker Compose environment configuration (when running with Docker)

### Database
- The application uses PostgreSQL as the database
- Database migrations will run automatically on startup
- Default database port is 5432 (configurable via .env)

## API documentation
- API documentation can be found in the `/api_documentation.md` file

## Additional Information
- The application will create necessary database tables automatically on first run
- For development, you can use the provided Makefile commands (if available)

## Troubleshooting
- Ensure all required ports are available (5500 for API, 5432 for PostgreSQL)
- Check Docker logs if containers fail to start:
  ```bash
  docker-compose logs
  ```
- Verify that the .env file contains all required variables
- Ensure Docker daemon is running before using docker-compose commands