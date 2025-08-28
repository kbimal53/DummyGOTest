# Dummy Go API with PostgreSQL & Frontend

A complete full-stack application with a **Go REST API backend** and a **modern responsive frontend**. This API uses PostgreSQL database for data persistence and includes a beautiful web interface for managing users. **Now supports both traditional server deployment and Vercel serverless functions!**

## Features

- ✅ **Full-stack application** with Go backend + HTML/JS frontend
- ✅ **Beautiful responsive UI** built with Tailwind CSS and Alpine.js
- ✅ **CRUD operations** for users with real-time updates
- ✅ **RESTful API endpoints** with JSON responses
- ✅ **PostgreSQL database integration** with connection pooling
- ✅ **Vercel serverless function compatible**
- ✅ **Health monitoring dashboard** with API status
- ✅ **Environment variable configuration**
- ✅ **CORS support** for cross-origin requests
- ✅ **Error handling** with user-friendly notifications
- ✅ **Mobile-responsive design**

## Getting Started

### Quick Start - Frontend Dashboard
🎯 **Visit the live application**: [https://test-go-programme.vercel.app/](https://test-go-programme.vercel.app/)

The web dashboard provides:
- Real-time API health monitoring
- Complete user management interface (CRUD operations)
- Beautiful responsive design
- Interactive data tables
- Instant notifications

### Local Development

#### Prerequisites

- Go 1.21 or higher
- PostgreSQL database (Vercel Postgres, Neon, or local PostgreSQL)
- Git (optional)

#### Installation

1. Clone or download this project
2. Navigate to the project directory:
   ```bash
   cd TestGoProgramme
   ```

3. Copy the environment file and configure your database:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` file with your database credentials.

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Run the full-stack application:
   ```bash
   make run
   # OR directly:
   go run main.go
   ```

The application will start on `http://localhost:8080`:
- **Frontend Dashboard**: `http://localhost:8080/`
- **API Base URL**: `http://localhost:8080/api/`

## API Endpoints

### Base URLs
- **Production (Vercel)**: `https://test-go-programme.vercel.app/api`
- **Local Development**: `http://localhost:8080/api`

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Frontend Dashboard |
| GET | `/api/health` | Health check |
| GET | `/api/users` | Get all users |
| GET | `/api/users/{id}` | Get user by ID |
| POST | `/api/users` | Create new user |
| PUT | `/api/users/{id}` | Update user |
| DELETE | `/api/users/{id}` | Delete user |

### Example API Usage

```bash
# Check API health
curl https://test-go-programme.vercel.app/api/health

# Get all users
curl https://test-go-programme.vercel.app/api/users

# Create a user
curl -X POST https://test-go-programme.vercel.app/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# Update a user
curl -X PUT https://test-go-programme.vercel.app/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Smith","email":"johnsmith@example.com"}'

# Delete a user
curl -X DELETE https://test-go-programme.vercel.app/api/users/1
```

### Sample Data

The API automatically creates a `users` table on startup and populates it with 3 sample users if the table is empty:
1. John Doe (john@example.com)
2. Jane Smith (jane@example.com)
3. Bob Johnson (bob@example.com)

## Database Schema

The `users` table has the following structure:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Project Structure

```
TestGoProgramme/
├── api/                    # Vercel serverless functions
│   └── api.go             # Main serverless function handler
├── public/                 # Frontend assets
│   ├── index.html         # Main dashboard page
│   └── style.css          # Custom CSS styles
├── main.go                # Traditional server (local development)
├── database.go            # Database connection & schema
├── vercel.json            # Vercel deployment configuration
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
├── Makefile               # Development automation
├── .env.example           # Environment variables template
└── README.md              # Project documentation
```
├── database.go            # Database connection for traditional server
├── migrate.go             # Database migration script
├── main_test.go           # Unit tests
├── go.mod                 # Main project dependencies
├── vercel.json            # Vercel configuration
├── deploy.sh              # Deployment preparation script
├── .env                   # Environment variables (not in git)
├── .env.example           # Environment variables template
├── README.md              # This file
└── Makefile               # Development commands
```

## Frontend Features

### 🎨 Modern Web Dashboard
The frontend provides a beautiful, responsive user interface built with:

- **Tailwind CSS**: Modern, utility-first styling
- **Alpine.js**: Lightweight reactive framework
- **Font Awesome**: Professional iconography
- **Responsive Design**: Works perfectly on desktop, tablet, and mobile

### 🔧 Interactive Features

1. **Real-time API Health Monitoring**
   - Live status indicator with color coding
   - Health endpoint response time tracking
   - Automatic health checks every 30 seconds

2. **Complete User Management**
   - View all users in a beautiful data table
   - Create new users with instant validation
   - Edit users inline with click-to-edit functionality
   - Delete users with confirmation dialogs
   - Real-time updates without page refresh

3. **User Experience**
   - Loading states for all operations
   - Success/error notifications
   - Form validation with helpful error messages
   - Mobile-optimized responsive layout
   - Intuitive navigation and interaction patterns

### 📱 UI Components
- **Dashboard Header**: API status and branding
- **Stats Cards**: User count and system metrics
- **Data Table**: Sortable, interactive user listing
- **Forms**: Create and edit user modals
- **Notifications**: Toast messages for user feedback

## Environment Variables

Create a `.env` file in the project root with the following variables:

```env
# Database connection (required)
DATABASE_URL=postgresql://username:password@host:port/database?sslmode=require

# Optional: Server port (defaults to 8080)
PORT=8080

# Additional PostgreSQL variables (for Vercel/Neon)
POSTGRES_URL=postgresql://username:password@host:port/database?sslmode=require
POSTGRES_USER=username
POSTGRES_HOST=host
POSTGRES_PASSWORD=password
POSTGRES_DATABASE=database
```

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
- **Database**: PostgreSQL with connection pooling
- **Environment Config**: Using godotenv for environment variables
- **CORS**: Cross-Origin Resource Sharing enabled
- **JSON**: All responses in JSON format
- **Error Handling**: Proper HTTP status codes and error messages
- **Validation**: Input validation for required fields
- **Middleware**: CORS middleware implementation
- **Auto-Migration**: Automatic table creation on startup
- **Sample Data**: Automatic insertion of sample data

## Deployment

### Vercel (Serverless) - Recommended

This project is now optimized for Vercel's serverless architecture!

1. **Prepare for deployment:**
   ```bash
   ./deploy.sh
   ```

2. **Push to GitHub:**
   ```bash
   git add .
   git commit -m "Add Vercel serverless support"
   git push origin main
   ```

3. **Deploy to Vercel:**
   - Go to [vercel.com](https://vercel.com) and sign in
   - Click "New Project" and import your GitHub repository
   - Set environment variables in Vercel dashboard:
     - `DATABASE_URL`
     - `POSTGRES_URL`
     - `POSTGRES_USER`
     - `POSTGRES_HOST`
     - `POSTGRES_PASSWORD`
     - `POSTGRES_DATABASE`
   - Deploy!

### Traditional Server Deployment

For traditional server hosting (Railway, Render, VPS):

1. **Build and run:**
   ```bash
   go build -o api main.go database.go
   ./api
   ```

2. **Using Docker:**
   ```bash
   # Build image
   docker build -t go-api .
   
   # Run container
   docker run -p 8080:8080 --env-file .env go-api
   ```

### Railway/Render

1. Connect your GitHub repository
2. Set environment variables
3. Deploy with automatic builds

## Notes

- Uses PostgreSQL for persistent data storage
- Automatic database table creation and migration
- Environment variables for secure credential management
- Production-ready with proper error handling
- CORS enabled for frontend integration
