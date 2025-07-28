.PHONY: build run test clean docker-build docker-run

# 构建应用
build:
	go build -o main .

# 运行应用
run:
	go run main.go

# 测试应用
test:
	go test ./...

# 清理构建文件
clean:
	rm -f main

# 下载依赖
deps:
	go mod download
	go mod tidy

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# Docker构建
docker-build:
	docker build -t qxbis-backend .

# Docker运行
docker-run:
	docker run -p 8000:8000 qxbis-backend

# 开发模式（热重载）
dev:
	air

# 安装开发工具
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 帮助信息
help:
	@echo "可用的命令:"
	@echo "  build        - 构建应用"
	@echo "  run          - 运行应用"
	@echo "  test         - 运行测试"
	@echo "  clean        - 清理构建文件"
	@echo "  deps         - 下载依赖"
	@echo "  fmt          - 格式化代码"
	@echo "  lint         - 代码检查"
	@echo "  docker-build - Docker构建"
	@echo "  docker-run   - Docker运行"
	@echo "  dev          - 开发模式（需要安装air）"
	@echo "  install-tools- 安装开发工具"
	@echo "  help         - 显示帮助信息" 