package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 本目录用于展示在gin框架下如何返回json格式的数据
// 为什么要返回json格式的数据？
// 在使用如Vue、React等前端框架时，模板样式一般是写好的
// 这个时候服务器只需要把json数据传输过去即可
func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法1：使用map
		//data := map[string]interface{}{
		//	"name":    "duoduo",
		//	"age":     20,
		//	"message": "hello world",
		//}
		// 使用gin内置的快捷类型
		data := gin.H{"name": "duoduo", "msg": "hello world", "age": 18}
		c.JSON(http.StatusOK, data)
	})

	// 方法2：结构体
	type msg struct {
		Name    string `json:"name"`
		Age     int
		Message string
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "duoduo",
			Age:     18,
			Message: "hello world",
		}
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("run server error:%s\n", err)
		return
	}
}
