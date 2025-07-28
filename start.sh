#!/bin/bash

# QXBIS Go API 启动脚本

echo "🚀 启动QXBIS Go API..."

# 设置环境变量
export DB_HOST=localhost
export REDIS_HOST=localhost

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go未安装，请先安装Go"
    exit 1
fi

# 检查依赖
echo "📦 检查依赖..."
go mod tidy

# 构建应用
echo "🔨 构建应用..."
go build -o main .

# 启动应用
echo "🌟 启动应用..."
./main 