package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 重定向
// 指的是当请求传到服务器后，把请求传到其它服务器或当前服务器上的其它网址
// 具体应用场景：比如一个用户是未登录状态，当他点击个人中心时，将页面跳转回登录页
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		// 重定向到百度
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r.GET("/a", func(c *gin.Context) {
		// 跳转到b对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的URI修改
		r.HandleContext(c)        // 继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
