package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 参数绑定
// 为了能够更方便的获取请求相关的参数，提高开发效率，我们可以基于请求的Content-type识别请求数据类型并利用反射机制自动提取请求中querystring等参数到结构体中
// 使用.ShouldBind()方法，它基于请求自动提取JSON、form表单和querystring类型的数据，并把值绑定到指定的结构体对象

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{username, password}
		var u UserInfo          // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u) // 把请求里面跟username和password相关的值取出来赋值给u 这里因为要对原始的u做修改，所以要传指针
		if err != nil {
			fmt.Printf("bind error: %v\n", err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
		}
		fmt.Printf("%v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			fmt.Printf("bind error: %v\n", err)
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
		fmt.Printf("%v\n", u)
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
