package handlers

import (
	"net/http"

	"qxbis-backend/database"

	"github.com/gin-gonic/gin"
)

// HealthHandler 健康检查处理器
type HealthHandler struct {
	dbManager *database.DatabaseManager
}

// NewHealthHandler 创建健康检查处理器
func NewHealthHandler(dbManager *database.DatabaseManager) *HealthHandler {
	return &HealthHandler{
		dbManager: dbManager,
	}
}

// HealthCheck 健康检查
// @Summary 健康检查
// @Description 检查服务健康状态
// @Tags 健康检查
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "服务正常"
// @Failure 503 {object} map[string]interface{} "服务异常"
// @Router /health/ [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	status := gin.H{
		"status":  "ok",
		"message": "服务运行正常",
	}

	// 检查PostgreSQL连接
	if err := h.dbManager.GetPGConn().Ping(); err != nil {
		status["postgres"] = "error"
		status["postgres_error"] = err.Error()
	} else {
		status["postgres"] = "ok"
	}

	// 检查Redis连接
	ctx := c.Request.Context()
	if err := h.dbManager.GetRedisClient().Ping(ctx).Err(); err != nil {
		status["redis"] = "error"
		status["redis_error"] = err.Error()
	} else {
		status["redis"] = "ok"
	}

	// 如果任何数据库连接失败，返回503状态码
	if status["postgres"] == "error" || status["redis"] == "error" {
		c.JSON(http.StatusServiceUnavailable, status)
		return
	}

	c.JSON(http.StatusOK, status)
}

// ReadinessCheck 就绪检查
// @Summary 就绪检查
// @Description 检查服务是否就绪
// @Tags 健康检查
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "服务就绪"
// @Failure 503 {object} map[string]interface{} "服务未就绪"
// @Router /health/ready [get]
func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	status := gin.H{
		"status":  "ready",
		"message": "服务已就绪",
	}

	// 检查PostgreSQL连接
	if err := h.dbManager.GetPGConn().Ping(); err != nil {
		status["status"] = "not_ready"
		status["postgres"] = "error"
		status["postgres_error"] = err.Error()
		c.JSON(http.StatusServiceUnavailable, status)
		return
	}

	// 检查Redis连接
	ctx := c.Request.Context()
	if err := h.dbManager.GetRedisClient().Ping(ctx).Err(); err != nil {
		status["status"] = "not_ready"
		status["redis"] = "error"
		status["redis_error"] = err.Error()
		c.JSON(http.StatusServiceUnavailable, status)
		return
	}

	status["postgres"] = "ok"
	status["redis"] = "ok"
	c.JSON(http.StatusOK, status)
}
