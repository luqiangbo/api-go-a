package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// DelayRequest 延迟请求结构体
type DelayRequest struct {
	Time int `json:"time" binding:"required"` // 延迟时间，单位秒
}

// DelayResponse 延迟响应结构体
type DelayResponse struct {
	Message   string    `json:"message"`
	DelayTime int       `json:"delay_time"`
	Timestamp time.Time `json:"timestamp"`
}

// DateResponse 日期响应结构体
type DateResponse struct {
	Date      string    `json:"date"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	// 打印启动信息
	fmt.Println("🚀 API Go A 服务启动中...")
	fmt.Println("📅 启动时间:", time.Now().Format("2006-01-02 15:04:05"))
	
	// 设置Gin为发布模式
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()
	
	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})
	
	// 健康检查接口
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API服务正常运行",
			"timestamp": time.Now(),
		})
	})
	
	// 日期接口
	r.POST("/date", func(c *gin.Context) {
		// 获取当前时间
		now := time.Now()

		// 返回响应
		response := DateResponse{
			Date:      now.Format("2006-01-02"),
			Timestamp: now,
		}

		c.JSON(http.StatusOK, response)
	})

	// 延迟接口
	r.GET("/delay", func(c *gin.Context) {
		// 从查询参数获取延迟时间
		timeStr := c.Query("time")
		if timeStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "缺少必需的time参数",
			})
			return
		}

		// 将时间字符串转换为整数
		delayTime, err := strconv.Atoi(timeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "time参数必须是有效的整数",
			})
			return
		}

		// 验证延迟时间（本地开发可以更长）
		if delayTime < 0 || delayTime > 60 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "延迟时间必须在0-60秒之间（本地开发）",
			})
			return
		}

		// 记录开始时间
		startTime := time.Now()

		// 延迟指定时间
		time.Sleep(time.Duration(delayTime) * time.Second)

		// 计算实际延迟时间
		actualDelay := int(time.Since(startTime).Seconds())

		// 返回响应
		response := DelayResponse{
			Message:   "延迟完成！",
			DelayTime: actualDelay,
			Timestamp: time.Now(),
		}

		c.JSON(http.StatusOK, response)
	})
	
	// 启动服务器
	port := ":" + getPort()
	fmt.Println("🌐 服务器地址: http://localhost" + port)
	fmt.Println("📋 可用接口:")
	fmt.Println("   GET  /     - 健康检查")
	fmt.Println("   POST /delay - 延迟接口")
	fmt.Println("   POST /date  - 日期接口")
	fmt.Println("⏳ 正在启动服务器...")
	
	err := r.Run(port)
	if err != nil {
		fmt.Println("❌ 服务器启动失败:", err)
		os.Exit(1)
	}
}

// getPort 获取端口号
func getPort() string {
	if port := getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}

// getenv 获取环境变量
func getenv(key string) string {
	return os.Getenv(key)
}