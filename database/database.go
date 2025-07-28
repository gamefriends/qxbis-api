package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"qxbis-backend/config"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// DatabaseManager 数据库连接管理器
type DatabaseManager struct {
	PGConn      *sql.DB
	RedisClient *redis.Client
	settings    *config.Settings
}

// NewDatabaseManager 创建数据库管理器
func NewDatabaseManager(settings *config.Settings) *DatabaseManager {
	return &DatabaseManager{
		settings: settings,
	}
}

// ConnectPostgres 连接PostgreSQL
func (dm *DatabaseManager) ConnectPostgres() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dm.settings.DBHost,
		dm.settings.DBPort,
		dm.settings.DBUser,
		dm.settings.DBPassword,
		dm.settings.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logrus.Errorf("PostgreSQL连接失败: %v", err)
		return err
	}

	// 设置连接池参数
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 测试连接
	if err := db.Ping(); err != nil {
		logrus.Errorf("PostgreSQL连接测试失败: %v", err)
		return err
	}

	dm.PGConn = db
	logrus.Info("PostgreSQL连接成功")
	return nil
}

// ConnectRedis 连接Redis
func (dm *DatabaseManager) ConnectRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", dm.settings.RedisHost, dm.settings.RedisPort),
		DB:       dm.settings.RedisDB,
		Password: "", // 如果有密码，从配置中读取
	})

	// 测试连接
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		logrus.Errorf("Redis连接失败: %v", err)
		return err
	}

	dm.RedisClient = rdb
	logrus.Info("Redis连接成功")
	return nil
}

// Close 关闭所有连接
func (dm *DatabaseManager) Close() {
	if dm.PGConn != nil {
		dm.PGConn.Close()
		logrus.Info("PostgreSQL连接已关闭")
	}

	if dm.RedisClient != nil {
		dm.RedisClient.Close()
		logrus.Info("Redis连接已关闭")
	}
}

// GetPGConn 获取PostgreSQL连接
func (dm *DatabaseManager) GetPGConn() *sql.DB {
	return dm.PGConn
}

// GetRedisClient 获取Redis客户端
func (dm *DatabaseManager) GetRedisClient() *redis.Client {
	return dm.RedisClient
}
