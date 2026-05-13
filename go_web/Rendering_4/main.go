package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 本目录用于展示模板的继承操作
// 观察目录下的index和home模板可以发现，这两个模板的大部分内容相同
// 此时我们可以定义一个block（根模板）来让子模版继承
// 实现请看templates目录

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板（继承方式）
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template filed,err:%v\n", err)
		return
	}
	// 渲染模板
	msg := "duoduo"
	err = t.ExecuteTemplate(w, "index2.tmpl", msg)
	if err != nil {
		fmt.Printf("execute template filed,err:%v\n", err)
		return
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板(模板继承)
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template filed,err:%v\n", err)
		return
	}
	// 渲染模板
	msg := "dio"
	err = t.ExecuteTemplate(w, "home2.tmpl", msg)
	if err != nil {
		fmt.Printf("execute template filed,err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
