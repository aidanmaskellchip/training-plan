# Training Plan Creator

## Overview
This project is an evolving application designed to generate personalized running training plans. It aims to provide tailored plans based on a user's current running abilities, historical activity data, and specific race goals (e.g., 5K, 10K, Half Marathon, Full Marathon). The system will intelligently adapt training schedules to help users achieve their fitness objectives.

## Features
- **Personalized Plan Generation:** Create training plans customized to individual running profiles and goals.
- **Activity Tracking Integration:** (Future) Integrate with external services to import and analyze user activity data.
- **Progress Monitoring:** (Future) Visualize user progress and adapt plans dynamically.
- **API-Driven:** A robust backend API to manage users, running profiles, and training plans.

## Technologies
- **Backend:** Go
- **Database:** PostgreSQL
- **Containerization:** Docker, Docker Compose
- **Dependency Management:** Go Modules
- **Testing:** Go's built-in testing framework, Testify

## Local Development

### Prerequisites
- Go (version 1.24.4 or higher recommended)
- Docker and Docker Compose

### Setup
To get the project up and running locally, execute the setup command:
```bash
make setup
```
This command will:
1.  Build the Docker images for the application services.
2.  Start the PostgreSQL database and other necessary services using Docker Compose.
3.  Run database migrations to set up the schema.
4.  Install Go dependencies.

The API will typically be accessible at `http://localhost:4000`.

### Local Database
- The local PostgreSQL database is accessible via port `8432`.
- Database credentials can be found within the `docker-compose.yml` file.

### Database Utilities
- **Run Migrations:** To apply any new database migrations, use:
  ```bash
  make run-migrations
  ```

### Linting
To run the Go linter on the project:
```bash
make go-lint
```

### Testing
To run all Go tests and generate a coverage report:
```bash
make go-tests
```
