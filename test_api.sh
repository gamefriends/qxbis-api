#!/bin/bash

# API测试脚本
BASE_URL="http://localhost:8000"

echo "测试QXBIS Go API..."

# 测试根路径
echo "1. 测试根路径..."
curl -s "$BASE_URL/" | jq .

# 测试健康检查
echo -e "\n2. 测试健康检查..."
curl -s "$BASE_URL/api/v1/health/" | jq .

# 测试事件收集
echo -e "\n3. 测试事件收集..."
curl -s -X POST "$BASE_URL/api/v1/events/" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "user_login",
    "data": {
      "user_id": "12345",
      "ip": "192.168.1.1",
      "timestamp": "2024-01-01T12:00:00Z"
    }
  }' | jq .

# 测试获取事件列表
echo -e "\n4. 测试获取事件列表..."
curl -s "$BASE_URL/api/v1/events/user_login" | jq .

# 测试获取事件计数
echo -e "\n5. 测试获取事件计数..."
curl -s "$BASE_URL/api/v1/events/user_login/count" | jq .

# 测试统计信息
echo -e "\n6. 测试统计信息..."
curl -s "$BASE_URL/api/v1/stats/" | jq .

echo -e "\n测试完成！" 