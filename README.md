# Go API with Docker, Hot Reload, and PostgreSQL

A complete Go REST API with Docker containerization, hot reload development environment, database migrations, and PostgreSQL.

## Features

- 🚀 **Hot Reload Development** - Uses Air for automatic rebuilding on code changes
- 🐳 **Docker & Docker Compose** - Fully containerized development environment
- 🗃️ **PostgreSQL Database** - With automated migrations using golang-migrate
- 📊 **Database Management** - Includes PgAdmin for database administration
- 🔧 **GORM Integration** - Modern ORM for Go with auto-migrations
- 🛣️ **REST API** - Complete CRUD operations for Users and Posts
- 🏗️ **Clean Architecture** - Well-structured project layout

## Quick Start

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone and navigate to the project**:

   ```bash
   cd go_api
   ```

2. **Start all services with hot reload**:

   ```bash
   docker-compose up --build
   ```

3. **Access the services**:

   - **API**: http://localhost:8080
   - **PgAdmin**: http://localhost:5050 (admin@admin.com / admin)
   - **PostgreSQL**: localhost:5432

4. **Health Check**:
   ```bash
   curl http://localhost:8080/health
   ```

## API Endpoints

### Users

- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get a specific user
- `PUT /api/v1/users/:id` - Update a user
- `DELETE /api/v1/users/:id` - Delete a user

### Posts

- `POST /api/v1/posts` - Create a new post
- `GET /api/v1/posts` - Get all posts
- `GET /api/v1/posts/:id` - Get a specific post
- `PUT /api/v1/posts/:id` - Update a post
- `DELETE /api/v1/posts/:id` - Delete a post

## Example API Usage

### Create a User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Create a Post

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{"title": "My First Post", "content": "Hello World!", "user_id": 1}'
```

### Get All Users

```bash
curl http://localhost:8080/api/v1/users
```

## Development

### Hot Reload

The development environment uses Air for hot reload. Any changes to `.go` files will automatically rebuild and restart the application.

### Database Migrations

Migrations are automatically run when the containers start. To create new migrations:

```bash
# Create a new migration
docker run --rm -v "${PWD}/migrations:/migrations" \
  migrate/migrate:v4.16.2 create -ext sql -dir /migrations -seq add_new_table
```

### Manual Migration Commands

```bash
# Run migrations
docker-compose run --rm migrate

# Rollback migrations
docker-compose run --rm migrate \
  -path /migrations \
  -database "postgres://postgres:password@postgres:5432/go_api_db?sslmode=disable" \
  down 1
```

## Project Structure

```
go_api/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection
│   ├── handlers/               # HTTP handlers
│   ├── models/                 # Data models
│   └── routes/                 # Route definitions
├── migrations/                 # Database migrations
├── docker-compose.yml          # Docker Compose configuration
├── Dockerfile                  # Multi-stage Docker build
├── .air.toml                   # Air configuration for hot reload
├── go.mod                      # Go module definition
└── README.md                   # This file
```

## Environment Variables

| Variable      | Default     | Description              |
| ------------- | ----------- | ------------------------ |
| `DB_HOST`     | `localhost` | PostgreSQL host          |
| `DB_PORT`     | `5432`      | PostgreSQL port          |
| `DB_USER`     | `postgres`  | PostgreSQL username      |
| `DB_PASSWORD` | `password`  | PostgreSQL password      |
| `DB_NAME`     | `go_api_db` | PostgreSQL database name |
| `DB_SSLMODE`  | `disable`   | PostgreSQL SSL mode      |
| `GIN_MODE`    | `debug`     | Gin framework mode       |
| `PORT`        | `8080`      | Application port         |

## Stopping the Application

```bash
docker-compose down
```

To remove all data:

```bash
docker-compose down -v
```

## Production Build

To build for production:

```bash
docker build --target production -t go-api:prod .
```
