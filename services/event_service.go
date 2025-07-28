package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"qxbis-backend/database"
	"qxbis-backend/models"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// EventService 事件服务
type EventService struct {
	dbManager *database.DatabaseManager
}

// NewEventService 创建事件服务
func NewEventService(dbManager *database.DatabaseManager) *EventService {
	return &EventService{
		dbManager: dbManager,
	}
}

// CollectEvent 收集事件数据
func (es *EventService) CollectEvent(event *models.EventData) *models.EventResponse {
	// 存储到 PostgreSQL
	payloadJSON, err := json.Marshal(event.Data)
	if err != nil {
		logrus.Errorf("事件数据序列化失败: %v", err)
		return &models.EventResponse{
			Status:  "error",
			Message: fmt.Sprintf("事件数据序列化失败: %v", err),
		}
	}

	query := `INSERT INTO events (event_type, payload, created_at) VALUES ($1, $2, $3)`
	_, err = es.dbManager.GetPGConn().Exec(query, event.Type, payloadJSON, time.Now())
	if err != nil {
		logrus.Errorf("事件存储到PostgreSQL失败: %v", err)
		return &models.EventResponse{
			Status:  "error",
			Message: fmt.Sprintf("事件存储失败: %v", err),
		}
	}

	// 更新 Redis 计数
	ctx := context.Background()
	key := fmt.Sprintf("event_count:%s", event.Type)
	err = es.dbManager.GetRedisClient().Incr(ctx, key).Err()
	if err != nil {
		logrus.Errorf("Redis计数更新失败: %v", err)
		// 不返回错误，因为主要数据已经存储成功
	}

	logrus.Infof("事件收集成功: %s", event.Type)
	return &models.EventResponse{
		Status: "ok",
	}
}

// GetEventsByType 根据类型获取事件列表
func (es *EventService) GetEventsByType(eventType string, limit int) ([]models.Event, error) {
	query := `SELECT id, event_type, payload, created_at FROM events WHERE event_type = $1 ORDER BY created_at DESC LIMIT $2`

	rows, err := es.dbManager.GetPGConn().Query(query, eventType, limit)
	if err != nil {
		logrus.Errorf("查询事件列表失败: %v", err)
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		var payloadJSON []byte

		err := rows.Scan(&event.ID, &event.EventType, &payloadJSON, &event.CreatedAt)
		if err != nil {
			logrus.Errorf("扫描事件数据失败: %v", err)
			continue
		}

		// 解析JSON数据
		if err := json.Unmarshal(payloadJSON, &event.Payload); err != nil {
			logrus.Errorf("解析事件数据失败: %v", err)
			event.Payload = make(map[string]interface{})
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		logrus.Errorf("遍历事件数据失败: %v", err)
		return nil, err
	}

	return events, nil
}

// GetEventCount 获取事件计数
func (es *EventService) GetEventCount(eventType string) (int64, error) {
	ctx := context.Background()
	key := fmt.Sprintf("event_count:%s", eventType)

	count, err := es.dbManager.GetRedisClient().Get(ctx, key).Int64()
	if err != nil && err != redis.Nil {
		logrus.Errorf("获取事件计数失败: %v", err)
		return 0, err
	}

	return count, nil
}
