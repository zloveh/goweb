package main

import (
	"context"
	"ginweb/src/conf"
	"ginweb/src/gin-server/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 加载配置文件
	conf.InitConfig("../../conf/conf.toml")

	// 初始化数据库
	conf.InitDB(conf.GlobalConfig)

	// 同时将日志写入文件和控制台
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 注册路由
	router.Router()

	// 启动服务
	StartServer()

}

func StartServer() {
	srv := &http.Server{
		Handler: router.RouterMux,
		Addr:    ":8080",
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\\n", err)
		}
	}()

	// 等待中断信号以超时 5 秒正常关闭服务器
	quit := make(chan os.Signal)
	// kill 命令发送信号 syscall.SIGTERM
	// kill -2 命令发送信号 syscall.SIGINT
	// kill -9 命令发送信号 syscall.SIGKILL

	//将对应的信号通知 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// 5 秒后捕获 ctx.Done() 信号
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
