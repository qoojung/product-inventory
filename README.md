# Product Inventory
![example workflow](https://github.com/qoojung/product-inventory/actions/workflows/go.yml/badge.svg)
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
## Tables

### Products Table

| Column Name | Data Type | Description |
|------------|-----------|-------------|
| id | int8 | Primary key, auto-incrementing |
| sku | varchar(64) | Stock Keeping Unit, unique identifier for product |
| name | varchar(64) | Name of the product |
| description | varchar(128) | Detailed description of the product |
| quantity | int8 | Available quantity of the product |
| created_at | timestamp with timezone| Timestamp when the record was created |
| updated_at | timestamp with timezone | Timestamp when the record was last updated |

#### Indexes
- Primary Key on `id`

## Additional Information
- The application will create necessary database tables automatically on first run
- For development, you can use the provided Makefile commands (if available)
