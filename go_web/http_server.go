package main

import (
	"fmt"
	"net/http"
	"os"
)

// 输入是一个回应和一个请求的指针
func sayHello(w http.ResponseWriter, r *http.Request) {
	byteData, err := os.ReadFile("./hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 向w中写入hello.txt
	_, _ = fmt.Fprintf(w, string(byteData))
}

func main() {
	// 在浏览器中输入/hello 返回SayHello方法的内容
	http.HandleFunc("/hello", sayHello)
	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed,err:%v\n", err)
		return
	}
}
