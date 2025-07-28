// Package main QXBIS Backend
//
// QXBIS事件收集系统后端服务
//
//	Schemes: http, https
//	Host: localhost:8000
//	BasePath: /api/v1
//	Version: 1.0.0
//	Title: QXBIS API
//	Description: 齐心BI系统事件收集API
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- api_key
//
// swagger:meta
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"qxbis-backend/config"
	"qxbis-backend/database"
	"qxbis-backend/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	// 配置日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Info("应用启动中...")

	// 加载配置
	settings := config.LoadSettings()

	// 创建数据库管理器
	dbManager := database.NewDatabaseManager(settings)

	// 连接数据库
	retryCount := 0
	for retryCount < settings.MaxRetries {
		err := dbManager.ConnectPostgres()
		if err != nil {
			logrus.Warnf("PostgreSQL连接失败，重试 %d/%d: %v", retryCount+1, settings.MaxRetries, err)
			retryCount++
			time.Sleep(time.Duration(settings.RetryDelay) * time.Second)
			continue
		}

		err = dbManager.ConnectRedis()
		if err != nil {
			logrus.Warnf("Redis连接失败，重试 %d/%d: %v", retryCount+1, settings.MaxRetries, err)
			retryCount++
			time.Sleep(time.Duration(settings.RetryDelay) * time.Second)
			continue
		}

		logrus.Info("所有数据库连接成功！")
		break
	}

	if retryCount >= settings.MaxRetries {
		logrus.Error("无法连接到数据库，应用启动失败")
		os.Exit(1)
	}

	// 设置路由
	router := routes.SetupRoutes(dbManager)

	// 创建HTTP服务器
	server := &http.Server{
		Addr:    ":" + settings.ServerPort,
		Handler: router,
	}

	// 启动服务器
	go func() {
		logrus.Infof("服务器启动在端口 %s", settings.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("应用关闭中...")

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("服务器强制关闭:", err)
	}

	// 关闭数据库连接
	dbManager.Close()

	logrus.Info("应用已关闭")
}
