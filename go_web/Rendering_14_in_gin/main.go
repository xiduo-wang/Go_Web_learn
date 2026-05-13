package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由及路由组

func main() {
	r := gin.Default()
	// 访问/index的GET请求会走这一条处理逻辑
	// 路由 以GET等请求访问/index 会执行内部定义的匿名函数
	// 查
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	// 增
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	// 改
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	// 删
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	// 可以处理所有类型的请求
	r.Any("/home", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case "POST":
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
		// ...
	})
	// 没有路由的页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "liwenzhou.com"})
	})

	// 视频的首页和详情页
	// 商城的首页和详情页
	// 路由组 多用于区分不同的业务线或API版本
	// 使用Group方法，把公用的前缀提取出来
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/index",
			})
		})
		shopGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/index",
			})
		})
		shopGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/index",
			})
		})

	}
	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
