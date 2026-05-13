package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// gin框架中间件
// 帮助开发者在接收请求的时候可以加入自己的函数(Hook)
// 中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
// 这里公共的业务逻辑指的是每个路由，如/index、/home等它们的函数中都需要做的公有部分

// 定义一个中间件m1:统计请求处理函数耗时
// 中间件必须是HandleFunc类型
// 即func(c *gin.Context){}
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()
	// 在中间件中使用goroutine去开线程的时候，必须使用上下文的副本，（使用copy方法），
	// 因为如果使用原context的话，如果在goroutine中修改c，那么中间件中其他的操作可能会受到影响，这是并发不安全的
	// 如go func xx(c.Copy())
	c.Next() // 调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost %v\n", cost)
	fmt.Println("m1 end")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "duoduo") // 跨中间件存取值，应用场景为中间件之间可能需要共享一些数据
	//c.Abort() // 阻止调用后续的处理函数
	fmt.Println("m2 end")
}

func indexHandle(c *gin.Context) {
	fmt.Println("indexHandle in ...")
	name, ok := c.Get("name")
	if !ok {
		name = "nil"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}

// 登录状态检测
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 使用闭包操作
	// 在return前可以做一些连接数据库等前置操作
	return func(c *gin.Context) {
		// if 是登录状态
		// c.Next()
		// else c.Abort()
		if doCheck {

		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(true)) // 全局注册中间件m1、m2、autoMiddleware
	// r.GET("/index", m1, indexHandle)
	r.GET("/index", indexHandle)

	// 定义路由组中间件
	//xxGroup := r.Group("/xx", authMiddleware(false))
	//{
	//	xxGroup.GET("/", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": "xxGroup",
	//		})
	//	})
	//}

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run Error: %v\n", err)
		return
	}
}
