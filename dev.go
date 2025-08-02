package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	handler "api-go-a/api"
)

func main() {
	// 设置服务器地址和端口
	addr := "localhost:8080"

	// 创建HTTP服务器
	server := &http.Server{
		Addr:         addr,
		Handler:      http.HandlerFunc(handler.Handler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 打印启动信息
	fmt.Println("API Go A 服务启动中...")
	fmt.Printf("可用接口：\n")
	fmt.Printf("GET  http://%s/       - 健康检查\n", addr)
	fmt.Printf("GET  http://%s/delay  - 延迟接口\n", addr)
	fmt.Printf("POST http://%s/date   - 日期接口\n", addr)

	// 启动服务器
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}