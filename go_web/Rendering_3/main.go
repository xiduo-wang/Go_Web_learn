package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 本目录用于展示模板中的自定义函数及模板嵌套操作，由于模板中内置的函数无法满足很多需求，大部分情况下需要我们自定义函数

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值为error类型
	kua := func(name string) (string, error) {
		return name + "nb", nil
	}
	// 告诉模板我现在多了一个自定义的函数
	t, err := template.New("f.tmpl").Funcs(template.FuncMap{
		"kua": kua,
	}).ParseFiles("./f.tmpl") // 解析模板
	if err != nil {
		fmt.Printf("template parse failed err:%v\n", err)
		return
	}
	// 渲染模板
	name := "duoduo"
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板 被包含的模板要写在后面
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("template parse failed err:%v\n", err)
		return
	}
	name := "dio"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/demo1", demo1)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		fmt.Printf("http serve failed err:%v\n", err)
		return
	}
}
