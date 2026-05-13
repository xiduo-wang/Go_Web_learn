package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 静态文件 html页面上用到的的样式文件.css js文件 图片

func main() {
	r := gin.Default()
	// gin定义一个自定义函数safe展示不转义操作
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	// 当我们渲染的HTML文件中引用了静态文件时，我们只需要按照以下方法在渲染页面前调用gin.statics方法即可
	// 加载静态文件
	r.Static("/xxx", "./statics")
	// 解析模板，在gin框架中使用LoadHTMLFlies 或 LoadHTMLGLOB 来加载 这里由于使用前者虽然也可以，但如果对解析的文件加量的话显然会变得冗余
	// 所以这里使用GLOB方法，通过正则匹配的加载方式解析模板
	// r.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl")
	// 代表解析templates下的所有目录中的所有文件
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		// type H map[string]any
		// http请求及模板渲染 H方法即为渲染步骤
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts/index.tmpl",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		// type H map[string]any
		// http请求及模板渲染 H方法即为渲染步骤
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})

	// 启动Server
	err := r.Run(":1010")
	if err != nil {
		fmt.Printf("run error: %v\n", err)
		return
	}
}
