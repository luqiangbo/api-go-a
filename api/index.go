package handler

import (
	"net/http"
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

// 全局路由实例
var router *gin.Engine

// Handler 是 Vercel serverless 函数的主入口
func Handler(w http.ResponseWriter, r *http.Request) {
	if router == nil {
		initRouter()
	}
	router.ServeHTTP(w, r)
}

// initRouter 初始化路由
func initRouter() {
	// 设置Gin为发布模式
	gin.SetMode(gin.ReleaseMode)
	
	router = gin.Default()
	
	// 添加CORS中间件
	router.Use(func(c *gin.Context) {
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
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API服务正常运行",
			"timestamp": time.Now(),
		})
	})
	
	// 延迟接口
	router.GET("/delay", func(c *gin.Context) {
		// 从查询参数获取延迟时间
		timeStr := c.Query("time")
		if timeStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "缺少必需的time参数",
			})
			return
		}

		// 将时间字符串转换为整数
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "time参数必须是有效的整数",
			})
			return
		}

		// 验证延迟时间（Vercel 限制为 10 秒）
		if time < 0 || time > 10 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "延迟时间必须在0-10秒之间（Vercel 限制）",
			})
			return
		}

		// 记录开始时间
		startTime := time.Now()

		// 延迟指定时间
		time.Sleep(time.Duration(time) * time.Second)

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
}