package handler

import (
	"encoding/json"
	"io"
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
	// 设置Content-Type
	w.Header().Set("Content-Type", "application/json")

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
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Method not allowed",
			})
			return
		}

		// ⚠️ 必须读取并关闭请求体以避免 Vercel 阻塞/超时
		_, _ = io.ReadAll(r.Body)
		_ = r.Body.Close()

		// 获取当前时间
		now := time.Now()
		response := DateResponse{
			Date:      now.Format("2006-01-02"),
			Timestamp: now,
		}
		json.NewEncoder(w).Encode(response)

	case "/delay":
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 获取延迟时间参数
		timeStr := r.URL.Query().Get("time")
		if timeStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Missing time parameter",
			})
			return
		}

		delayTime, err := strconv.Atoi(timeStr)
		if err != nil || delayTime < 0 || delayTime > 10 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Delay time must be between 0 and 10 seconds",
			})
			return
		}

		// 延迟执行
		time.Sleep(time.Duration(delayTime) * time.Second)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    "延迟完成！",
			"delay_time": delayTime,
			"timestamp":  time.Now(),
		})

	default:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Not Found",
		})
	}
}
