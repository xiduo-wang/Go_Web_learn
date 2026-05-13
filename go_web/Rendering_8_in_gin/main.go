package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// querystring
// 即URL地址中？后面那一串用来传递信息的参数
// 基本结构就是一系列的Key-Value键值对
// 在实际开发中，常用于数据传递（如在搜索栏中输入‘手机’，网站会通过query string把这个关键词传给服务器，服务器再返回相应的商品列表）、状态记录（刷新页面时还能保留在之前打开的页面）
// 功能控制（比如商品的按价格排序可以写成?sort=price或?sort=time）等
// 多个key=value用&连接
func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器那边发请求携带的 query String
		name := c.Query("query") // 通过Query获取请求中携带的query string
		//name := c.DefaultQuery("query", "somebody") // 在其中按query关键词查找时取不到，就用默认值
		//name, ok := c.GetQuery("query") // 取不到query通过接到的false来输出默认值
		//if !ok {
		//	// 取不到的情况
		//	name = "somebody"
		//}
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
