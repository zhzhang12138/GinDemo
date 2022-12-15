package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.Default()

	Router.GET("/ping", func(c *gin.Context) {
		sleepTime, _ := strconv.Atoi(c.Query("time"))
		time.Sleep(time.Duration(sleepTime) * time.Second)
		//fmt.Println(123456789)
		c.JSON(200, gin.H{
			"message": "sleep---" + c.Query("time") + "s",
		})
	})

	return Router
}

// 初始化服务
func initServer(address string, router *gin.Engine) {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 100 * time.Second // 该参数表示读取请求头的最长时间。
	s.WriteTimeout = 100 * time.Second
	s.MaxHeaderBytes = 1 << 20
	// 使用 endless.ListenAndServe 方法启动 HTTP 服务器
	s.ListenAndServe()
}

func main() {
	// endless实现不停机重启 Go 程序
	// 参考文章：https://www.luozhiyun.com/archives/584
	Router := Routers()
	initServer(":8080", Router)
}
