# QXBIS Backend (Go版本)

这是一个使用Go语言编写的BI事件收集系统后端服务。

## 功能特性

- 事件数据收集和存储
- PostgreSQL数据持久化
- Redis缓存和计数
- RESTful API接口
- 健康检查和监控
- 统计信息查询

## 技术栈

- **语言**: Go 1.21+
- **Web框架**: Gin
- **数据库**: PostgreSQL
- **缓存**: Redis
- **日志**: Logrus
- **配置**: 环境变量

## 项目结构

```
backend/
├── config/          # 配置管理
├── database/        # 数据库连接
├── handlers/        # HTTP处理器
├── models/          # 数据模型
├── routes/          # 路由配置
├── services/        # 业务逻辑
├── main.go          # 主程序入口
├── go.mod           # Go模块文件
├── Dockerfile       # Docker配置
└── env.example      # 环境变量示例
```

## API接口

### 事件管理

- `POST /api/v1/events/` - 收集事件数据
- `GET /api/v1/events/{event_type}` - 获取事件列表
- `GET /api/v1/events/{event_type}/count` - 获取事件计数

### 健康检查

- `GET /api/v1/health/` - 健康检查
- `GET /api/v1/health/ready` - 就绪检查

### 统计信息

- `GET /api/v1/stats/` - 获取总体统计
- `GET /api/v1/stats/{event_type}` - 获取特定事件类型统计

## 快速开始

### 本地开发

1. 安装Go 1.21+
2. 克隆项目
3. 进入backend目录
4. 安装依赖：
   ```bash
   go mod download
   ```
5. 配置环境变量（参考env.example）
6. 启动服务：
   ```bash
   go run main.go
   ```

### Docker部署

1. 构建镜像：
   ```bash
   docker build -t qxbis-backend .
   ```

2. 运行容器：
   ```bash
   docker run -p 8000:8000 qxbis-backend
   ```

### Docker Compose

使用项目根目录的docker-compose.yml文件：

```bash
docker-compose up -d
```

## 环境变量

| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| APP_NAME | BI事件收集系统 | 应用名称 |
| APP_VERSION | 1.0.0 | 应用版本 |
| DEBUG | false | 调试模式 |
| SERVER_PORT | 8000 | 服务端口 |
| DB_HOST | postgres | 数据库主机 |
| DB_PORT | 5432 | 数据库端口 |
| DB_NAME | bi_db | 数据库名称 |
| DB_USER | bi_user | 数据库用户 |
| DB_PASSWORD | bi_pass | 数据库密码 |
| REDIS_HOST | redis | Redis主机 |
| REDIS_PORT | 6379 | Redis端口 |
| REDIS_DB | 0 | Redis数据库 |
| MAX_RETRIES | 30 | 最大重试次数 |
| RETRY_DELAY | 2 | 重试延迟（秒） |

## 数据模型

### EventData
```go
type EventData struct {
    Type string                 `json:"type"`
    Data map[string]interface{} `json:"data"`
}
```

### EventResponse
```go
type EventResponse struct {
    Status  string `json:"status"`
    Message string `json:"message,omitempty"`
}
```

## 开发说明

### 添加新的API接口

1. 在`handlers/`目录下创建新的处理器
2. 在`routes/routes.go`中注册路由
3. 在`services/`目录下实现业务逻辑

### 数据库迁移

数据库表结构定义在`sql/db.sql`文件中，启动时会自动创建。

### 日志

使用Logrus进行日志记录，支持JSON格式输出。

## 性能优化

- 使用连接池管理数据库连接
- Redis缓存热点数据
- 异步处理非关键操作
- 优雅关闭机制

## 监控

- 健康检查接口：`/api/v1/health/`
- 就绪检查接口：`/api/v1/health/ready`
- 统计信息接口：`/api/v1/stats/` 