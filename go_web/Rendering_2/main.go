package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 模板语法
// {{.}}
// 模板语法都包含在{{和}}之间。其中的.表示当前对象
// 当我们传入一个结构体对象的时候，我们就可以通过.来访问结构体中的字段
// 想传入map也是同理

type StudentInfo struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello2.tmpl")
	if err != nil {
		fmt.Printf("Parse failed err:%v\n", err)
		return
	}
	// 利用给定数据渲染模板并写入
	student := StudentInfo{ // 在tmpl中的{{}}中我们可以通过.Name等方式来单独传入字段
		Name:   "DUODUO",
		Age:    20,
		Gender: "male",
	}
	// 再来一个map
	m := map[string]interface{}{
		"name":   "duoduo",
		"age":    20,
		"gender": "male",
	}
	// 演示模板中的range操作
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	// 传入多个值进行渲染
	err = t.Execute(w, map[string]interface{}{
		"student": student,
		"m":       m,
		"hobby":   hobbyList,
	})
	if err != nil {
		fmt.Printf("Execute failed err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("http serve failed err:%v\n", err)
		return
	}
}
