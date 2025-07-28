package models

import (
	"time"
)

// EventData 事件数据模型
type EventData struct {
	Type string                 `json:"type" binding:"required"` // 事件类型
	Data map[string]interface{} `json:"data" binding:"required"` // 事件数据
}

// EventResponse 事件响应模型
type EventResponse struct {
	Status  string `json:"status"`            // 状态
	Message string `json:"message,omitempty"` // 消息
}

// Event 数据库事件模型
type Event struct {
	ID        int64                  `json:"id" db:"id"`
	EventType string                 `json:"event_type" db:"event_type"`
	Payload   map[string]interface{} `json:"payload" db:"payload"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
}

// EventListResponse 事件列表响应
type EventListResponse struct {
	EventType string  `json:"event_type"`
	Events    []Event `json:"events"`
	Count     int     `json:"count"`
}
