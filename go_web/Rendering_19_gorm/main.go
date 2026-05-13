package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 小清单项目
// To do Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	dsn := "root:Root#123@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 测试联通性
	return DB.DB().Ping()
}

func main() {
	// 创建数据库
	// sql :CREATE DATABASE bubble
	// 连接数据库 调用函数
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() // 程序退出关闭数据库
	// 绑定模型
	DB.AutoMigrate(&Todo{})
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "./dist/static")
	// 告诉gin去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项，点击提交，发请求到这里
			// 从请求中，把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 存入数据库 返回响应
			if err = DB.Create(&todo).Error; err != nil {
				// 创建失败返回error
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项（根据id）
		v1Group.GET("/todo/:id", func(c *gin.Context) {})
		// 修改状态
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			//  1. 从 URL 路径参数取 id
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id is not exist",
				})
			}
			var todo Todo
			// 2. 根据 id 查数据库，找到这条记录
			if err = DB.Where("id = ?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
			// 3. 把用户提交的 JSON 绑定到 todo 结构体（覆盖查出来的数据）
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id is not exist",
				})
			}

			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	err = r.Run(":8080")
	if err != nil {
		fmt.Printf("Run err: %v\n", err)
		return
	}
}
