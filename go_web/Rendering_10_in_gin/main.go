package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取URI参数（path）
// 请求的参数通过URL路径传递，例如/user/search/duoduo/柳林

func main() {
	r := gin.Default()
	// GET方法中访问请求处写格式
	r.GET("/user/:username/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("username")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 第二个博客例子
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
