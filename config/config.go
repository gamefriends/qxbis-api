package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Settings 应用配置结构
type Settings struct {
	// 应用基础配置
	AppName    string
	AppVersion string
	Debug      bool

	// 数据库配置
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string

	// Redis配置
	RedisHost string
	RedisPort int
	RedisDB   int

	// 应用配置
	MaxRetries int
	RetryDelay int
	APIPrefix  string
	ServerPort string
}

// LoadSettings 加载配置
func LoadSettings() *Settings {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		logrus.Warn("未找到.env文件，使用默认配置")
	}

	settings := &Settings{
		// 应用基础配置
		AppName:    getEnv("APP_NAME", "BI事件收集系统"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		Debug:      getEnvAsBool("DEBUG", false),

		// 数据库配置
		DBHost:     getEnv("DB_HOST", "postgres"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBName:     getEnv("DB_NAME", "bi_db"),
		DBUser:     getEnv("DB_USER", "bi_user"),
		DBPassword: getEnv("DB_PASSWORD", "bi_pass"),

		// Redis配置
		RedisHost: getEnv("REDIS_HOST", "redis"),
		RedisPort: getEnvAsInt("REDIS_PORT", 6379),
		RedisDB:   getEnvAsInt("REDIS_DB", 0),

		// 应用配置
		MaxRetries: getEnvAsInt("MAX_RETRIES", 30),
		RetryDelay: getEnvAsInt("RETRY_DELAY", 2),
		APIPrefix:  getEnv("API_PREFIX", "/api/v1"),
		ServerPort: getEnv("SERVER_PORT", "8000"),
	}

	return settings
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt 获取环境变量并转换为整数
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsBool 获取环境变量并转换为布尔值
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
