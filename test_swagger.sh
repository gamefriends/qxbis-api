#!/bin/bash

# Swagger测试脚本
BASE_URL="http://localhost:8000"

echo "🧪 测试QXBIS Swagger API文档..."

# 测试根路径
echo "1. 测试根路径..."
curl -s "$BASE_URL/" | jq .

# 测试Swagger文档JSON
echo -e "\n2. 测试Swagger文档JSON..."
curl -s "$BASE_URL/swagger/doc.json" | jq .info

# 测试Swagger UI页面
echo -e "\n3. 测试Swagger UI页面..."
curl -s -I "$BASE_URL/swagger/index.html" | head -5

# 测试API路径
echo -e "\n4. 测试API路径..."
curl -s "$BASE_URL/swagger/doc.json" | jq '.paths | keys'

# 测试健康检查
echo -e "\n5. 测试健康检查..."
curl -s "$BASE_URL/api/v1/health/" | jq .

echo -e "\n✅ Swagger测试完成！"
echo -e "📖 访问Swagger UI: http://localhost:8000/swagger/index.html" 