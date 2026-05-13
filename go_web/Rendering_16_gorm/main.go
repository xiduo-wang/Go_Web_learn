package main

// GORM
// 指的是go中流行的一种ORM框架
// ORM指的是Object Relational Mapping(对象关系映射)
// 对象指的是程序中的对象/实例 例如Go中的结构体实例
// 关系指的是关系数据库 例如MySQL
// 所以ORM框架做的就是把一个程序中的实例与关系数据库的一个数据表形成一个映射关系
// 虽然ORM提高了开发效率，但牺牲了执行性能、牺牲了灵活性、弱化了SQL能力
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Age    int
	Hobby  string
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:Root#123@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()

	// 创建表 自动迁移 把结构体和数据表进行对应
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{1, "duoduo", "男", 19, "game"}
	db.Create(&u1)
	// 查询
	var u UserInfo
	db.First(&u) // 查询表中第一条数据保存到u中
	fmt.Printf("U:%v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "篮球")

	// 删除
	db.Delete(&u)
}
