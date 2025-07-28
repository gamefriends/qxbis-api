package handlers

import (
	"net/http"
	"strconv"

	"qxbis-backend/models"
	"qxbis-backend/services"

	"github.com/gin-gonic/gin"
)

// EventHandler 事件处理器
type EventHandler struct {
	eventService *services.EventService
}

// NewEventHandler 创建事件处理器
func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

// CollectEvent 收集事件数据
// @Summary 收集事件数据
// @Description 收集业务事件数据并存储到数据库
// @Tags 事件管理
// @Accept json
// @Produce json
// @Param event body models.EventData true "事件数据"
// @Success 200 {object} models.EventResponse "成功"
// @Failure 400 {object} map[string]interface{} "请求数据格式错误"
// @Failure 500 {object} models.EventResponse "服务器内部错误"
// @Router /events/ [post]
func (h *EventHandler) CollectEvent(c *gin.Context) {
	var event models.EventData

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求数据格式错误: " + err.Error(),
		})
		return
	}

	result := h.eventService.CollectEvent(&event)

	if result.Status == "error" {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetEventsByType 根据类型获取事件列表
// @Summary 获取事件列表
// @Description 根据事件类型获取事件列表
// @Tags 事件管理
// @Accept json
// @Produce json
// @Param event_type path string true "事件类型"
// @Param limit query int false "限制数量" default(100)
// @Success 200 {object} models.EventListResponse "成功"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /events/{event_type} [get]
func (h *EventHandler) GetEventsByType(c *gin.Context) {
	eventType := c.Param("event_type")
	if eventType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "事件类型不能为空",
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}

	events, err := h.eventService.GetEventsByType(eventType, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取事件列表失败: " + err.Error(),
		})
		return
	}

	response := models.EventListResponse{
		EventType: eventType,
		Events:    events,
		Count:     len(events),
	}

	c.JSON(http.StatusOK, response)
}

// GetEventCount 获取事件计数
// @Summary 获取事件计数
// @Description 获取特定事件类型的计数
// @Tags 事件管理
// @Accept json
// @Produce json
// @Param event_type path string true "事件类型"
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /events/{event_type}/count [get]
func (h *EventHandler) GetEventCount(c *gin.Context) {
	eventType := c.Param("event_type")
	if eventType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "事件类型不能为空",
		})
		return
	}

	count, err := h.eventService.GetEventCount(eventType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取事件计数失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"event_type": eventType,
		"count":      count,
	})
}
