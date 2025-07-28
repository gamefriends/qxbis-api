package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEventData_MarshalJSON(t *testing.T) {
	event := EventData{
		Type: "test_event",
		Data: map[string]interface{}{
			"user_id": "12345",
			"action":  "login",
		},
	}

	data, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var result EventData
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if result.Type != event.Type {
		t.Errorf("期望类型 %s, 实际 %s", event.Type, result.Type)
	}

	if result.Data["user_id"] != event.Data["user_id"] {
		t.Errorf("期望user_id %v, 实际 %v", event.Data["user_id"], result.Data["user_id"])
	}
}

func TestEventResponse(t *testing.T) {
	response := EventResponse{
		Status:  "ok",
		Message: "事件收集成功",
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var result EventResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if result.Status != response.Status {
		t.Errorf("期望状态 %s, 实际 %s", response.Status, result.Status)
	}

	if result.Message != response.Message {
		t.Errorf("期望消息 %s, 实际 %s", response.Message, result.Message)
	}
}

func TestEvent(t *testing.T) {
	now := time.Now()
	event := Event{
		ID:        1,
		EventType: "test_event",
		Payload: map[string]interface{}{
			"user_id": "12345",
		},
		CreatedAt: now,
	}

	// 测试JSON标签
	data, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var result Event
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if result.ID != event.ID {
		t.Errorf("期望ID %d, 实际 %d", event.ID, result.ID)
	}

	if result.EventType != event.EventType {
		t.Errorf("期望事件类型 %s, 实际 %s", event.EventType, result.EventType)
	}
}
