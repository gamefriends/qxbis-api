# Swagger API文档支持

## 概述

QXBIS Go版本已完全支持Swagger API文档，提供交互式的API文档界面。

## 功能特性

### ✅ 已实现的功能

1. **自动生成API文档**
   - 基于代码注释自动生成Swagger文档
   - 支持所有API端点的文档化
   - 包含请求参数、响应格式、错误码

2. **交互式API文档**
   - Swagger UI界面
   - 在线API测试功能
   - 实时API调用演示

3. **完整的API覆盖**
   - 事件管理API
   - 健康检查API
   - 统计信息API

## 访问方式

### Swagger UI
```
http://localhost:8000/swagger/index.html
```

### Swagger JSON文档
```
http://localhost:8000/swagger/doc.json
```

### Swagger YAML文档
```
http://localhost:8000/swagger/swagger.yaml
```

## API文档结构

### 事件管理 (Events)
- `POST /api/v1/events/` - 收集事件数据
- `GET /api/v1/events/{event_type}` - 获取事件列表
- `GET /api/v1/events/{event_type}/count` - 获取事件计数

### 健康检查 (Health)
- `GET /api/v1/health/` - 健康检查
- `GET /api/v1/health/ready` - 就绪检查

### 统计信息 (Stats)
- `GET /api/v1/stats/` - 获取总体统计
- `GET /api/v1/stats/{event_type}` - 获取特定事件类型统计

## 数据模型

### EventData
```json
{
  "type": "string",
  "data": "object"
}
```

### EventResponse
```json
{
  "status": "string",
  "message": "string"
}
```

### EventListResponse
```json
{
  "event_type": "string",
  "events": "array",
  "count": "integer"
}
```

## 开发指南

### 添加新的API文档

1. **添加Swagger注释**
```go
// @Summary API摘要
// @Description API详细描述
// @Tags 标签名
// @Accept json
// @Produce json
// @Param param_name path string true "参数描述"
// @Success 200 {object} ResponseModel "成功响应"
// @Failure 400 {object} ErrorResponse "错误响应"
// @Router /api/path [method]
func (h *Handler) APIHandler(c *gin.Context) {
    // 实现代码
}
```

2. **重新生成文档**
```bash
swag init
```

3. **重启应用**
```bash
go run main.go
```

### 测试Swagger

使用提供的测试脚本：
```bash
./test_swagger.sh
```

## 技术实现

### 依赖包
- `github.com/swaggo/swag` - Swagger文档生成
- `github.com/swaggo/gin-swagger` - Gin集成
- `github.com/swaggo/files` - 静态文件服务

### 配置
- 自动扫描代码注释
- 生成JSON和YAML格式文档
- 提供Swagger UI界面

## 与Python版本的对比

| 特性 | Python FastAPI | Go Gin + Swagger |
|------|----------------|-------------------|
| 自动文档生成 | ✅ 内置 | ✅ 需要注释 |
| 交互式测试 | ✅ 内置 | ✅ Swagger UI |
| 文档格式 | OpenAPI 3.0 | Swagger 2.0 |
| 代码注释 | 自动解析 | 需要手动注释 |
| 实时文档 | ✅ | ✅ |

## 优势

1. **完整的API文档**
   - 所有API端点都有详细文档
   - 包含请求/响应示例
   - 支持在线测试

2. **开发者友好**
   - 交互式API测试
   - 实时参数验证
   - 清晰的错误信息

3. **维护简单**
   - 基于代码注释自动生成
   - 代码和文档同步更新
   - 版本控制友好

## 使用示例

### 1. 访问Swagger UI
打开浏览器访问：`http://localhost:8000/swagger/index.html`

### 2. 测试API
在Swagger UI中可以直接：
- 查看API文档
- 测试API调用
- 查看响应格式
- 下载API规范

### 3. 集成到其他工具
- 导入到Postman
- 生成客户端代码
- API网关集成

## 注意事项

1. **注释格式**
   - 必须使用标准的Swagger注释格式
   - 注释必须紧贴函数定义

2. **模型定义**
   - 响应模型需要正确定义
   - 使用JSON标签标注字段

3. **路径参数**
   - 路径参数必须正确标注
   - 参数类型必须准确

## 故障排除

### 常见问题

1. **文档不更新**
   - 运行 `swag init` 重新生成
   - 重启应用

2. **Swagger UI无法访问**
   - 检查应用是否正常运行
   - 确认端口8000未被占用

3. **注释不生效**
   - 检查注释格式是否正确
   - 确认注释紧贴函数定义

### 调试命令
```bash
# 重新生成文档
swag init

# 测试Swagger
./test_swagger.sh

# 检查应用状态
ps aux | grep "go run main.go"
```

## 总结

QXBIS Go版本已完全支持Swagger API文档，提供了：

- ✅ 完整的API文档覆盖
- ✅ 交互式API测试界面
- ✅ 自动文档生成
- ✅ 开发者友好的体验

这大大提升了API的可维护性和开发效率！ 