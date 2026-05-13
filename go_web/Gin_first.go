package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	// 返回一个Json格式的内容
	c.JSON(200, gin.H{
		// type H map[string]any 一个结构体map，key永远是string
		"message": "hello world",
	})
}

func main() {
	// 返回默认的路由引擎
	r := gin.Default()
	// 指定用户使用GET请求访问/hello2时，执行SayHello这个函数
	r.GET("/hello2", SayHello)

	// RESTful 风格API
	// 状态码可以通过http包的status方法拿到，演示的OK即为200
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
		// 开发RESTful API时我们通常使用Postman来作为客户端测试工具
		// 启动服务
		r.Run(":1100")
	})
}
