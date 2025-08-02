# API Go Two - 延迟接口服务

这是一个部署在 Vercel 上的 Golang API 服务，使用 Gin 框架提供延迟接口。

## 功能特性

- 🚀 基于 Gin 框架的轻量级 API 服务
- ⏰ 提供延迟接口，支持自定义延迟时间
- 🌐 支持 CORS 跨域请求
- 📊 健康检查接口
- 🔒 输入验证和错误处理

## API 接口

### 1. 健康检查接口

**GET /**

返回服务状态信息

**响应示例：**

```json
{
  "message": "API服务正常运行",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

### 2. 延迟接口

**POST /delay**

根据指定时间延迟后返回响应

**请求体：**

```json
{
  "time": 5
}
```

**参数说明：**

- `time`: 延迟时间（秒），范围：0-60 秒

**响应示例：**

```json
{
  "message": "延迟完成！",
  "delay_time": 5,
  "timestamp": "2024-01-01T12:00:05Z"
}
```

## 本地开发

### 环境要求

- Go 1.21+

### 安装依赖

```bash
go mod tidy
```

### 运行服务

有多种方式启动项目：

#### 1. 直接使用 go run

```bash
go run main.go
```

#### 2. 使用 Makefile (Linux/Mac)

```bash
make run          # 运行项目
make build        # 构建项目
make test         # 运行测试
make deps         # 安装依赖
make fmt          # 格式化代码
make dev          # 开发模式(热重载)
make clean        # 清理文件
```

#### 3. 热重载开发模式

安装 air 工具：

```bash
go install github.com/cosmtrek/air@latest
```

然后运行：

```bash
air
```

服务将在 `http://localhost:8080` 启动

### 测试接口

使用 curl 测试延迟接口：

```bash
curl -X POST http://localhost:8080/delay \
  -H "Content-Type: application/json" \
  -d '{"time": 3}'
```

## 部署到 Vercel

1. 安装 Vercel CLI：

```bash
npm i -g vercel
```

2. 登录 Vercel：

```bash
vercel login
```

3. 部署项目：

```bash
vercel
```

4. 生产环境部署：

```bash
vercel --prod
```

## 项目结构

```
api-go-two/
├── main.go          # 主程序文件
├── go.mod           # Go模块文件
├── vercel.json      # Vercel配置文件
├── Makefile         # Makefile脚本
├── .air.toml        # 热重载配置
└── README.md        # 项目说明文档
```

## 注意事项

- 延迟时间限制在 0-60 秒之间
- Vercel 函数最大执行时间为 60 秒
- 服务支持 CORS 跨域请求
- 所有接口都返回 JSON 格式数据
