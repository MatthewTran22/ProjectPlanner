# Project Planner

A simple Go web application built with the Gin framework.

## Features

- RESTful API endpoints
- Built with Gin framework
- Containerized with Docker

## API Endpoints

- `GET /` - Returns a simple "OKAY" message
- `GET /data` - Returns sample data

## Running with Docker

### Option 1: Using Docker Compose (Recommended)

1. Build and run the application:
   ```bash
   docker-compose up --build
   ```

2. Access the application at `http://localhost:8000`

3. To stop the application:
   ```bash
   docker-compose down
   ```

### Option 2: Using Docker directly

1. Build the Docker image:
   ```bash
   docker build -t project-planner .
   ```

2. Run the container:
   ```bash
   docker run -p 8000:8000 project-planner
   ```

3. Access the application at `http://localhost:8000`

## Development

### Prerequisites

- Go 1.24.5 or later
- Docker (for containerized deployment)

### Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Access the application at `http://localhost:8000`

## Docker Commands

### Build the image
```bash
docker build -t project-planner .
```

### Run the container
```bash
docker run -p 8000:8000 project-planner
```

### Run in detached mode
```bash
docker run -d -p 8000:8000 --name project-planner-app project-planner
```

### Stop the container
```bash
docker stop project-planner-app
```

### Remove the container
```bash
docker rm project-planner-app
```

### View logs
```bash
docker logs project-planner-app
```

## Health Check

The application includes a health check endpoint at `GET /` that returns a 200 status code when the service is running properly.

## Environment Variables

- `GIN_MODE`: Set to `release` for production (default: `debug`)

## Project Structure

```
ProjectPlanner/
├── controllers/
│   └── get.go          # API controllers
├── main.go             # Application entry point
├── go.mod              # Go module file
├── go.sum              # Go module checksums
├── Dockerfile          # Docker configuration
├── docker-compose.yml  # Docker Compose configuration
├── .dockerignore       # Docker ignore file
└── README.md           # This file
``` 