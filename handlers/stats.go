package handlers

import (
	"database/sql"
	"net/http"

	"qxbis-backend/database"

	"github.com/gin-gonic/gin"
)

// StatsHandler 统计处理器
type StatsHandler struct {
	dbManager *database.DatabaseManager
}

// NewStatsHandler 创建统计处理器
func NewStatsHandler(dbManager *database.DatabaseManager) *StatsHandler {
	return &StatsHandler{
		dbManager: dbManager,
	}
}

// GetStats 获取统计信息
// @Summary 获取总体统计
// @Description 获取系统总体统计信息
// @Tags 统计信息
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "统计信息"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /stats/ [get]
func (h *StatsHandler) GetStats(c *gin.Context) {
	stats := gin.H{
		"total_events":  0,
		"event_types":   []string{},
		"recent_events": 0,
	}

	// 获取总事件数
	var totalEvents int
	query := `SELECT COUNT(*) FROM events`
	err := h.dbManager.GetPGConn().QueryRow(query).Scan(&totalEvents)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取总事件数失败: " + err.Error(),
		})
		return
	}
	stats["total_events"] = totalEvents

	// 获取事件类型列表
	rows, err := h.dbManager.GetPGConn().Query(`SELECT DISTINCT event_type FROM events ORDER BY event_type`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取事件类型列表失败: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var eventTypes []string
	for rows.Next() {
		var eventType string
		if err := rows.Scan(&eventType); err != nil {
			continue
		}
		eventTypes = append(eventTypes, eventType)
	}
	stats["event_types"] = eventTypes

	// 获取最近24小时的事件数
	var recentEvents int
	query = `SELECT COUNT(*) FROM events WHERE created_at >= NOW() - INTERVAL '24 hours'`
	err = h.dbManager.GetPGConn().QueryRow(query).Scan(&recentEvents)
	if err != nil && err != sql.ErrNoRows {
		// 不返回错误，只是统计信息
		recentEvents = 0
	}
	stats["recent_events"] = recentEvents

	c.JSON(http.StatusOK, stats)
}

// GetEventTypeStats 获取特定事件类型的统计
// @Summary 获取事件类型统计
// @Description 获取特定事件类型的详细统计信息
// @Tags 统计信息
// @Accept json
// @Produce json
// @Param event_type path string true "事件类型"
// @Success 200 {object} map[string]interface{} "事件类型统计"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /stats/{event_type} [get]
func (h *StatsHandler) GetEventTypeStats(c *gin.Context) {
	eventType := c.Param("event_type")
	if eventType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "事件类型不能为空",
		})
		return
	}

	stats := gin.H{
		"event_type":  eventType,
		"total_count": 0,
		"today_count": 0,
		"week_count":  0,
	}

	// 获取总计数
	var totalCount int
	query := `SELECT COUNT(*) FROM events WHERE event_type = $1`
	err := h.dbManager.GetPGConn().QueryRow(query, eventType).Scan(&totalCount)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取事件计数失败: " + err.Error(),
		})
		return
	}
	stats["total_count"] = totalCount

	// 获取今日计数
	var todayCount int
	query = `SELECT COUNT(*) FROM events WHERE event_type = $1 AND created_at >= CURRENT_DATE`
	err = h.dbManager.GetPGConn().QueryRow(query, eventType).Scan(&todayCount)
	if err != nil && err != sql.ErrNoRows {
		todayCount = 0
	}
	stats["today_count"] = todayCount

	// 获取本周计数
	var weekCount int
	query = `SELECT COUNT(*) FROM events WHERE event_type = $1 AND created_at >= DATE_TRUNC('week', CURRENT_DATE)`
	err = h.dbManager.GetPGConn().QueryRow(query, eventType).Scan(&weekCount)
	if err != nil && err != sql.ErrNoRows {
		weekCount = 0
	}
	stats["week_count"] = weekCount

	c.JSON(http.StatusOK, stats)
}
