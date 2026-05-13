package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取form表单提交的参数

func main() {
	r := gin.Default()
	// 解析模板
	r.LoadHTMLFiles("./login.html", "./index.html")
	// 渲染模板
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 一次请求对应一个相应
	// 接收login post请求，即点击登录按钮之后
	// 这里写服务端接收前端返回的请求
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数据 三种方式
		// 1
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		// 2
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***")
		// 3
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "ss"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
