package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
	JWT      JWTConfig
	Fraud    FraudConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Host string
	Port string
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	VHost    string
}

type JWTConfig struct {
	Secret             string
	ExpiryHours        int
	RefreshExpiryHours int
}

type FraudConfig struct {
	MaxDailyAmount        float64
	MaxTransactionAmount  float64
	VelocityThreshold     int
	VelocityWindowMinutes int
}

type LoggingConfig struct {
	Level  string
	Format string
}

func Load() (*Config, error) {
	_ = godotenv.Load() // Load .env if exists

	config := &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "financial_user"),
			Password: getEnv("DB_PASSWORD", "financial_pass"),
			Name:     getEnv("DB_NAME", "financial_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		RabbitMQ: RabbitMQConfig{
			Host:     getEnv("RABBITMQ_HOST", "localhost"),
			Port:     getEnv("RABBITMQ_PORT", "5672"),
			User:     getEnv("RABBITMQ_USER", "financial_user"),
			Password: getEnv("RABBITMQ_PASS", "financial_pass"),
			VHost:    getEnv("RABBITMQ_VHOST", "financial_vhost"),
		},
		JWT: JWTConfig{
			Secret:             getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
			ExpiryHours:        getEnvAsInt("JWT_EXPIRY_HOURS", 24),
			RefreshExpiryHours: getEnvAsInt("JWT_REFRESH_EXPIRY_HOURS", 168),
		},
		Fraud: FraudConfig{
			MaxDailyAmount:        getEnvAsFloat("FRAUD_MAX_DAILY_AMOUNT", 50000.00),
			MaxTransactionAmount:  getEnvAsFloat("FRAUD_MAX_TRANSACTION_AMOUNT", 10000.00),
			VelocityThreshold:     getEnvAsInt("FRAUD_VELOCITY_THRESHOLD", 5),
			VelocityWindowMinutes: getEnvAsInt("FRAUD_VELOCITY_WINDOW_MINUTES", 10),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
	}

	return config, nil
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
}

func (c *Config) GetRabbitMQURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		c.RabbitMQ.User,
		c.RabbitMQ.Password,
		c.RabbitMQ.Host,
		c.RabbitMQ.Port,
		c.RabbitMQ.VHost,
	)
}

func (c *Config) GetJWTExpiry() time.Duration {
	return time.Duration(c.JWT.ExpiryHours) * time.Hour
}

func (c *Config) GetJWTRefreshExpiry() time.Duration {
	return time.Duration(c.JWT.RefreshExpiryHours) * time.Hour
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsFloat(name string, defaultVal float64) float64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultVal
}
