package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// 定义 -p 参数，用于指定端口号
	port := flag.String("p", "8080", "Port to run the HTTP server on")
	flag.Parse()

	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	fmt.Printf("Serving files from: %s\n", wd)

	// 创建一个 HTTP 文件服务器
	fs := http.FileServer(http.Dir(wd))

	// 自定义的根路由处理器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取当前时间
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		// 日志记录，包含时间、HTTP 方法、访问路径和客户端 IP
		log.Printf("[%s] [%s] %s %s\n", timestamp, r.Method, r.URL.Path, r.RemoteAddr)
		// 使用文件服务器提供的服务
		fs.ServeHTTP(w, r)
	})

	// 启动 HTTP 服务
	address := fmt.Sprintf(":%s", *port)
	fmt.Printf("Starting HTTP server on port %s...\n", *port)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
