package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 本目录讲解如何修改模版引擎的标识符，即{{}} 为什么修改移步README 同时也包含html/template的自动转义危险字符功能的相关内容

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板 在此处使用Delims方法和New搭配实现
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("template parse error: %v\n", err)
		return
	}
	// 渲染模板
	name := "duoduo"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("template execute error: %v\n", err)
		return
	}
}

// 定义一个危险字符函数
func xss(w http.ResponseWriter, r *http.Request) {
	// 解析模板之前创建一个函数safe来控制转义
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			// 这个函数做的是把传进来的字符串转成HTML类型
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Printf("template parse error: %v\n", err)
		return
	}
	// 危险字符
	str1 := "<script>alert(123);</script>"
	// 正常内容 由于html/template会对所有内容进行转义 所以我们可以自己定义一个函数来控制是否转义 移步解析文件前处
	str2 := "<script><a href='duoduo.com'>duoduo的博客</a></script>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Printf("template execute error: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":2020", nil)
	if err != nil {
		fmt.Printf("web serve failed, err:%v\n", err)
		return
	}
}
