package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 文件上传
// 上传文件本质就是客户端发送请求，服务端响应的过程
// 处理multipart forms提交文件时默认的内存限制是32MiB
// 可以通过如下方式修改
// router.MaxMultipartMemory = 8 << 20 // 8MiB
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			// 将读取到的文件保存在服务端本地 第一个参数为路径，第二个为读取到的文件名称
			file_path := fmt.Sprintf("./%s", file.Filename)
			err = c.SaveUploadedFile(file, file_path)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}

// 多个文件上传
//func main() {
//	router := gin.Default()
//	// 处理multipart forms提交文件时默认的内存限制是32 MiB
//	// 可以通过下面的方式修改
//	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
//	router.POST("/upload", func(c *gin.Context) {
//		// Multipart form
//		form, _ := c.MultipartForm()
//		files := form.File["file"]
//
//		for index, file := range files {
//			log.Println(file.Filename)
//			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
//			// 上传文件到指定的目录
//			c.SaveUploadedFile(file, dst)
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"message": fmt.Sprintf("%d files uploaded!", len(files)),
//		})
//	})
//	router.Run()
//}
