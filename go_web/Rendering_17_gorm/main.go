package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 在使用ORM工具时，通常我们需要在代码中定义Models来与数据库中的表进行映射，在GORM中模型通常是正常定义的结构体、基本的go类型或它们的指针
// 同时也支持sql.Scanner及driver.Valuer接口
// 为了方便模型定义，GORM内置了一个gorm.Models结构体，它是一个包含了ID、CreatedAt、UpdatedAt、DeletedAt四个字段的结构体
// 即
//type Model struct {
//	ID        uint `gorm:"primary_key"`
//	CreatedAt time.Time 创建时间
//	UpdatedAt time.Time 更新时间
//	DeletedAt *time.Time 删除时间
//}
// 在gorm中，默认使用名为ID的字段作为表的主键
// 但可以通过结构体tag或内嵌的方式来修改主键
// 表名默认就是结构体名称的复数
// 如下面的User的默认表名就是users
// 我们可以通过TableName()方法来修改表名

type User struct {
	gorm.Model   // 内嵌
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;no null"` // 设置会员号唯一且不空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置num为自增类型
	Address      string  `gorm:"index:addr"`     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`              // 忽略本字段
}

type Animal struct {
	AnimalID uint `gorm:"primary_key"`
}

// 修改表名
func (Animal) TableName() string {
	return "Duoduo"
}

func main() {
	// 一个默认的有关表名处理的函数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "jojo_" + defaultTableName
	}

	// 连接数据库
	db, err := gorm.Open("mysql", "root:Root#123@(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
	// 禁用复数
	db.SingularTable(true)
	// 自动迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用User结构体创建名叫 dio 的表
	//db.Table("dio").CreateTable(&User{})

}
