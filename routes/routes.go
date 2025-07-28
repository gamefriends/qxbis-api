package routes

import (
	"qxbis-backend/database"
	_ "qxbis-backend/docs" // 导入swagger文档
	"qxbis-backend/handlers"
	"qxbis-backend/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes 设置路由
func SetupRoutes(dbManager *database.DatabaseManager) *gin.Engine {
	router := gin.Default()

	// 创建服务和处理器
	eventService := services.NewEventService(dbManager)
	eventHandler := handlers.NewEventHandler(eventService)
	healthHandler := handlers.NewHealthHandler(dbManager)
	statsHandler := handlers.NewStatsHandler(dbManager)

	// 根路径
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "欢迎使用BI事件收集系统",
			"version": "1.0.0",
			"docs":    "/swagger/index.html",
		})
	})

	// Swagger文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/docs", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := router.Group("/api/v1")
	{
		// 事件相关路由
		events := api.Group("/events")
		{
			events.POST("/", eventHandler.CollectEvent)                  // 收集事件
			events.GET("/:event_type", eventHandler.GetEventsByType)     // 获取事件列表
			events.GET("/:event_type/count", eventHandler.GetEventCount) // 获取事件计数
		}

		// 健康检查路由
		health := api.Group("/health")
		{
			health.GET("/", healthHandler.HealthCheck)         // 健康检查
			health.GET("/ready", healthHandler.ReadinessCheck) // 就绪检查
		}

		// 统计相关路由
		stats := api.Group("/stats")
		{
			stats.GET("/", statsHandler.GetStats)                     // 获取总体统计
			stats.GET("/:event_type", statsHandler.GetEventTypeStats) // 获取特定事件类型统计
		}
	}

	return router
}
