# Dummy Go API

A simple REST API built with Go for managing users. This is a demonstration API with in-memory storage.

## Features

- ✅ CRUD operations for users
- ✅ RESTful endpoints
- ✅ JSON responses
- ✅ CORS support
- ✅ Error handling
- ✅ Health check endpoint

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git (optional)

### Installation

1. Clone or download this project
2. Navigate to the project directory:
   ```bash
   cd TestGoProgramme
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Root endpoint - Welcome message |
| GET | `/api/v1/health` | Health check |
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/{id}` | Get user by ID |
| POST | `/api/v1/users` | Create new user |
| PUT | `/api/v1/users/{id}` | Update user |
| DELETE | `/api/v1/users/{id}` | Delete user |

### Sample Data

The API comes with 3 sample users:
1. John Doe (john@example.com)
2. Jane Smith (jane@example.com)
3. Bob Johnson (bob@example.com)

## Example Usage

### Get all users
```bash
curl http://localhost:8080/api/v1/users
```

### Get a specific user
```bash
curl http://localhost:8080/api/v1/users/1
```

### Create a new user
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Wilson",
    "email": "alice@example.com"
  }'
```

### Update a user
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com"
  }'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### Health check
```bash
curl http://localhost:8080/api/v1/health
```

## Response Format

All API responses follow this format:

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { /* response data */ }
}
```

## User Object Structure

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created": "2023-08-27T10:30:00Z"
}
```

## Building for Production

To build a binary:

```bash
go build -o api main.go
```

To run the binary:

```bash
./api
```

## Features Included

- **Router**: Using Gorilla Mux for routing
- **CORS**: Cross-Origin Resource Sharing enabled
- **JSON**: All responses in JSON format
- **Error Handling**: Proper HTTP status codes and error messages
- **Validation**: Basic input validation
- **Middleware**: CORS middleware implementation
- **In-memory Storage**: Simple data persistence for demo

## Notes

- This is a demo API with in-memory storage
- Data will be lost when the server restarts
- No authentication/authorization implemented
- Suitable for development and testing purposes
