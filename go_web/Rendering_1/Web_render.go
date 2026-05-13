package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Go语言内置了文本模板引擎text/template和用于HTML文档的html/template。它们的作用机制可以简单归纳如下：
//
//模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
//模板文件中使用{{和}}包裹和标识需要传入的数据。
//传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
//除{{和}}包裹的内容外，其他内容均不做修改原样输出。

// 定义好了模板文件之后，可以使用下面的常用方法去解析模板文件，得到模板对象：
// func (t *Template) Parse(src string) (*Template, error)
// func ParseFiles(filenames ...string) (*Template, error)
// func ParseGlob(pattern string) (*Template, error)

// 渲染模板简单来说就是使用数据去填充模板，当然实际上可能会复杂很多。
// func (t *Template) Execute(wr io.Writer, data interface{}) error
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
// 在项目中我已经定义了hello.tmpl模板，下面我将对其进行实操
func sayHello2(w http.ResponseWriter, r *http.Request) {
	// 解析模板 找到哪里需要填入数据
	t, err := template.ParseFiles("./hello2.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed err:%v", err)
		return
	}
	// 渲染模板 设置填入的具体内容，使用Execute方法将数据写入w
	err = t.Execute(w, "duoduo")
	if err != nil {
		fmt.Printf("render template failed err:%v", err)
	}
}

func main() {
	http.HandleFunc("/", sayHello2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}
