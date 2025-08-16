# Financial Transaction Processing System

A robust backend system built with Go for handling financial transactions with ACID properties, fraud detection, and comprehensive audit logging.

## 🚀 Features

- **Account Management**: User registration, authentication, and account management
- **Transaction Processing**: Secure money transfers with ACID properties
- **Fraud Detection**: Real-time fraud detection rules and risk assessment
- **Audit Logging**: Comprehensive audit trail for all operations
- **Balance Reconciliation**: Automated balance calculations and reconciliation
- **Transaction History**: Detailed transaction statements and history
- **Real-time Notifications**: Transaction alerts via message queues

## 🛠 Tech Stack

- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Authentication**: JWT tokens
- **Documentation**: Swagger/OpenAPI
- **Testing**: Go testing framework + Testify
- **CI/CD**: GitHub Actions
- **Containerization**: Docker & Docker Compose

## 📋 Prerequisites

- Go 1.21+
- Docker & Docker Compose
- PostgreSQL 15+
- Redis 7+
- RabbitMQ 3.12+

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd financial-transaction-system
   ```

2. **Start services with Docker Compose**
   ```bash
   docker-compose up -d
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run database migrations**
   ```bash
   go run cmd/migrate/main.go up
   ```

5. **Start the application**
   ```bash
   go run cmd/server/main.go
   ```

6. **Access the API documentation**
   - Swagger UI: http://localhost:8080/swagger/index.html
   - API Base URL: http://localhost:8080/api/v1

## 📚 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh JWT token

### Account Management
- `GET /api/v1/accounts` - Get user accounts
- `POST /api/v1/accounts` - Create new account
- `GET /api/v1/accounts/{id}` - Get account details
- `GET /api/v1/accounts/{id}/balance` - Get account balance

### Transactions
- `POST /api/v1/transactions/transfer` - Transfer money between accounts
- `GET /api/v1/transactions` - Get transaction history
- `GET /api/v1/transactions/{id}` - Get transaction details
- `GET /api/v1/transactions/statement` - Generate account statement

### Admin
- `GET /api/v1/admin/audit-logs` - Get audit logs
- `GET /api/v1/admin/fraud-alerts` - Get fraud alerts
- `POST /api/v1/admin/reconcile` - Trigger balance reconciliation

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run integration tests:
```bash
go test -tags=integration ./...
```

## 🏗 Project Structure

```
financial-transaction-system/
├── cmd/
│   ├── server/          # Application entry point
│   └── migrate/         # Database migration tool
├── internal/
│   ├── api/             # HTTP handlers and routes
│   ├── auth/            # Authentication logic
│   ├── config/          # Configuration management
│   ├── db/              # Database connection and queries
│   ├── fraud/           # Fraud detection engine
│   ├── models/          # Data models
│   ├── queue/           # Message queue handlers
│   ├── services/        # Business logic
│   └── utils/           # Utility functions
├── migrations/          # SQL migration files
├── docker-compose.yml   # Docker services configuration
├── Dockerfile          # Container image definition
├── .github/workflows/  # CI/CD pipeline
└── docs/               # Additional documentation
```

## 🔒 Security Features

- JWT-based authentication
- Password hashing with bcrypt
- Input validation and sanitization
- SQL injection prevention
- Rate limiting
- Fraud detection algorithms
- Comprehensive audit logging

## 📊 Monitoring & Logging

- Structured logging with Logrus
- Transaction audit trails
- Performance metrics
- Error tracking
- Real-time fraud alerts

## 🚀 Deployment

The application includes a complete CI/CD pipeline with:
- Automated testing
- Code quality checks
- Security scanning
- Docker image building
- Deployment automation

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📞 Support

For support and questions, please open an issue in the GitHub repository. 