# Ecommerce-API-Assessment

## Overview
RESTful API for an e-commerce application. This API handle basic CRUD operations for products and orders, and provide user management and authentication.

## Technologies Used
- **Golang**: Programming language.
- **PostgreSQL**: An open-source object-relational database system.
- **Golang-Migrate**: An open-source migration tool.
- **Gorilla-Mux**: An open-source routing tool.
- **Godotenv**: An open-source library to load environment variables
- **golang-jwt**: An open-source library for jwt handling
- **Docker**: A containerization platform for seamless application building and sharing.

## Installation

### Docker (Preferred)
1. Clone the repository.
2. Create a `.env` file by copying from `.env.example`.
3. Run the following commands:
   ```bash
   make build      # Build the application
   make migrate_up # Apply migrations and create necessary tables
   make run          # Start the application
   ```
4. Access the application at `http://localhost:8080`.

### Local Installation
1. **Golang**: Install Golang. 
2. **PostgreSQL**: Install PostgreSQL. 
3. **Install Dependencies**:
   - Install all dependencies:
     ```bash
     go mod download
     ```
   - Install golang-migrate:
     ```bash
     go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
     ```
   - Create and configure a `.env` file using `.env.example`.
   - Export the below variables.
      ```bash
      export POSTGRESQL_URL="your postresql url in env"
      export ENVIRONMENT=dev
      ```
   - Run migrations:
     ```bash
     migrate -database ${POSTGRESQL_URL} -path db/migrations up
     ```
   - Start the server:
     ```bash
     go run main.go
     ```
4. Access the application at `http://localhost:8080`.

## API Endpoints
Access it here on postman: https://orange-desert-910094.postman.co/workspace/6df4b7b9-9934-49a0-87fa-1ec172988c0e/collection/25347207-52939234-7f7d-46bc-a6df-b76330b3c7ec