package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 增删改查
// 定义模型
type User struct {
	ID   int64
	Name *string `gorm:"default:'dio'"` // 可以通过tag定义字段的默认值，定义默认值后在创建记录时生成的SQL语句会排除没有值或值为0的字段
	// 引入默认值会有一个问题，如果我想传入""、null、false这样的零值的时候，它们也不会传入数据库，而是传入默认值
	// 这种情况我们可以使用指针进行传入或使用Scanner或Valuer 指针的零值是nil，可以很好地规避这个问题
	Age *int `gorm:"default:'6'"`
}

func main() {

	db, err := gorm.Open("mysql", "root:Root#123@(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()

	// 将模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 创建记录
	// 不传入Name字段时填入默认值dio
	// 使用指针后可以传入零值，使用临时变量避免解引用
	u := User{Name: func(s string) *string { return &s }("duoduo"), Age: func(i int) *int { return &i }(18)} // 在代码层面创建一个结构体对象
	fmt.Println(db.NewRecord(&u))                                                                            // 判断这个结构体是不是还没进过数据库 true 这是一条新记录
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u)) // false 已经创建
}
