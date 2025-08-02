package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// DateResponse 日期响应结构体
type DateResponse struct {
	Date      string    `json:"date"`
	Timestamp time.Time `json:"timestamp"`
}

// Handler 是 Vercel serverless 函数的主入口
func Handler(w http.ResponseWriter, r *http.Request) {
	// 设置CORS头
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "86400")

	// 处理OPTIONS请求
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 根据路径处理不同的请求
	switch r.URL.Path {
	case "/":
		// 健康检查
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   "API服务正常运行",
			"timestamp": time.Now(),
		})

	case "/date":
		// 只处理POST请求
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 获取当前时间并返回
		now := time.Now()
		response := DateResponse{
			Date:      now.Format("2006-01-02"),
			Timestamp: now,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	case "/delay":
		// 只处理GET请求
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 获取延迟时间参数
		timeStr := r.URL.Query().Get("time")
		if timeStr == "" {
			http.Error(w, "Missing time parameter", http.StatusBadRequest)
			return
		}

		// 转换延迟时间
		delayTime, err := strconv.Atoi(timeStr)
		if err != nil {
			http.Error(w, "Invalid time parameter", http.StatusBadRequest)
			return
		}

		// 验证延迟时间
		if delayTime < 0 || delayTime > 10 {
			http.Error(w, "Delay time must be between 0 and 10 seconds", http.StatusBadRequest)
			return
		}

		// 执行延迟
		time.Sleep(time.Duration(delayTime) * time.Second)

		// 返回响应
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   "延迟完成！",
			"delay_time": delayTime,
			"timestamp":  time.Now(),
		})

	default:
		http.NotFound(w, r)
	}
}