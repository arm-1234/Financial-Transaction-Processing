# Testing Guide

This document provides comprehensive information about testing the Financial Transaction Processing System.

## 🧪 Testing Strategy

Our testing approach follows the testing pyramid:
- **Unit Tests**: Test individual functions and methods
- **Integration Tests**: Test component interactions
- **End-to-End Tests**: Test complete user workflows
- **Performance Tests**: Test system performance under load

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- jq (for JSON parsing in scripts)

### Setup Test Environment
```bash
# Start test services
make docker-up

# Run database migrations
make migrate-up

# Run all tests
make test
```

## 📋 Test Categories

### 1. Unit Tests
Test individual components in isolation.

```bash
# Run unit tests
go test ./internal/...

# Run with coverage
go test -cover ./internal/...

# Run with verbose output
go test -v ./internal/...
```

### 2. Integration Tests
Test component interactions with real databases.

```bash
# Run integration tests (requires database)
go test -tags=integration ./...
```

### 3. API Tests
Test HTTP endpoints with real server.

```bash
# Start the server first
make run

# In another terminal, run API tests
make test-api
```

## 🔧 Test Configuration

### Environment Variables
```bash
# Test database
DB_HOST=localhost
DB_PORT=5432
DB_USER=financial_user
DB_PASSWORD=financial_pass
DB_NAME=financial_db_test

# Test Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# Test JWT
JWT_SECRET=test-secret-key
```

## 📊 API Testing Examples

### Authentication Flow
```bash
# Register a new user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "testpassword123",
    "first_name": "Test",
    "last_name": "User"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "testpassword123"
  }'

# Use the returned access_token for authenticated requests
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### User Management
```bash
# Get user profile
curl -H "Authorization: Bearer TOKEN" \
  http://localhost:8080/api/v1/users/profile

# Update profile
curl -X PUT http://localhost:8080/api/v1/users/profile \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TOKEN" \
  -d '{
    "first_name": "Updated",
    "address": "123 New Address"
  }'

# Change password
curl -X POST http://localhost:8080/api/v1/users/change-password \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TOKEN" \
  -d '{
    "current_password": "oldpassword",
    "new_password": "newpassword123"
  }'
```

## 🧩 Test Data

### Sample User Data
```json
{
  "email": "john.doe@example.com",
  "password": "securepassword123",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890",
  "address": "123 Main St, New York, NY"
}
```

### Sample Account Data
```json
{
  "account_type": "checking",
  "account_name": "John's Checking Account",
  "currency": "USD",
  "daily_limit": 5000.00,
  "monthly_limit": 50000.00,
  "is_primary": true
}
```

## 🔍 Testing Patterns

### Testing Database Operations
```go
func TestUserRepository_Create(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    repo := NewUserRepository(db)
    
    user := &models.User{
        ID:        uuid.New(),
        Email:     "test@example.com",
        FirstName: "Test",
        LastName:  "User",
    }
    
    err := repo.Create(user)
    assert.NoError(t, err)
    
    // Verify user was created
    retrieved, err := repo.GetByID(user.ID)
    assert.NoError(t, err)
    assert.Equal(t, user.Email, retrieved.Email)
}
```

### Testing HTTP Handlers
```go
func TestHandleRegister(t *testing.T) {
    // Setup
    userService := &MockUserService{}
    handler := handleRegister(userService)
    
    // Create request
    reqBody := `{
        "email": "test@example.com",
        "password": "password123",
        "first_name": "Test",
        "last_name": "User"
    }`
    
    req := httptest.NewRequest("POST", "/auth/register", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    c.Request = req
    
    // Execute
    handler(c)
    
    // Assert
    assert.Equal(t, http.StatusCreated, w.Code)
}
```

## 🚦 Test Status Indicators

### Expected Response Codes
- `200 OK` - Successful GET/PUT requests
- `201 Created` - Successful POST requests
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Missing/invalid authentication
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server errors

### Health Check
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "service": "financial-transaction-system"
}
```

## 🔄 Continuous Integration

Our CI pipeline runs:
1. **Unit Tests** - Fast, isolated tests
2. **Linting** - Code quality checks
3. **Security Scan** - Vulnerability detection
4. **Integration Tests** - Component interaction tests
5. **Docker Build** - Container image creation
6. **End-to-End Tests** - Full workflow validation

### Local CI Simulation
```bash
# Run the same checks as CI
make deps
make test
make build
make docker-build
```

## 🐛 Debugging Tests

### Common Issues

1. **Database Connection Errors**
   ```bash
   # Check if PostgreSQL is running
   docker ps | grep postgres
   
   # Check database connectivity
   psql -h localhost -U financial_user -d financial_db
   ```

2. **Port Conflicts**
   ```bash
   # Check what's using port 8080
   lsof -i :8080
   
   # Kill process if needed
   kill -9 <PID>
   ```

3. **JWT Token Issues**
   ```bash
   # Verify token format
   echo "TOKEN" | base64 -d
   ```

### Debug Mode
```bash
# Run with debug logging
LOG_LEVEL=debug make run

# Run tests with verbose output
go test -v -race ./...
```

## 📈 Performance Testing

### Load Testing with Apache Bench
```bash
# Install Apache Bench
sudo apt-get install apache2-utils

# Test registration endpoint
ab -n 100 -c 10 -H "Content-Type: application/json" \
   -p user_data.json http://localhost:8080/api/v1/auth/register

# Test authenticated endpoint
ab -n 100 -c 10 -H "Authorization: Bearer TOKEN" \
   http://localhost:8080/api/v1/users/profile
```

### Memory and CPU Profiling
```bash
# Run with profiling
go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.

# Analyze profiles
go tool pprof cpu.prof
go tool pprof mem.prof
```

## 📝 Test Reports

### Coverage Report
```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# View in browser
open coverage.html
```

### Test Results
Tests generate detailed output including:
- Pass/fail status
- Execution time
- Coverage percentage
- Performance metrics

## 🔐 Security Testing

### Authentication Tests
- Invalid credentials
- Expired tokens
- Missing authorization headers
- Token tampering

### Input Validation Tests
- SQL injection attempts
- XSS payloads
- Invalid JSON formats
- Boundary value testing

### Rate Limiting Tests
```bash
# Test rate limiting
for i in {1..20}; do
  curl http://localhost:8080/api/v1/auth/login &
done
```

## 📚 Additional Resources

- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Testify Framework](https://github.com/stretchr/testify)
- [Gin Testing Guide](https://gin-gonic.com/docs/testing/)
- [PostgreSQL Testing](https://www.postgresql.org/docs/current/regress.html) 