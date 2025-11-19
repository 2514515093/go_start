package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Name  string
	Age   int
	Grade string
	gorm.Model
}

func main() {
	//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）
	//、 grade （学生年级，字符串类型）。
	//要求 ：
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录

	db, _ := createDb()
	//insertStudent(db, &[]Student{
	//	{Name: "张三", Age: 20, Grade: "三年级"},
	//	{Name: "李四", Age: 14, Grade: "二年级"},
	//	{Name: "王五", Age: 19, Grade: "三年级"},
	//})

	//result := findStudent(db, gorm.Expr("age > ?", 18))
	//fmt.Println(result)
	//updateStudent(db, map[string]interface{}{"name": "张三"}, map[string]interface{}{"name": "小三子"})
	deleteStudent(db, gorm.Expr("age < ?", 15))
}

func createDb() (*gorm.DB, error) {
	dsn := "root:root123@tcp(211.159.169.85:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//创建表
	db.AutoMigrate(&Student{})
	return db, err
}

func insertStudent(db *gorm.DB, m *[]Student) int64 {
	result := db.Debug().Create(m)
	return result.RowsAffected
}

func updateStudent(db *gorm.DB, where map[string]interface{}, m map[string]interface{}) int64 {
	result := db.Model(&Student{}).Debug().Where(where).Updates(m)
	return result.RowsAffected
}

func findStudent(db *gorm.DB, m interface{}) []Student {
	var stu []Student
	db.Debug().Where(m).Find(&stu)
	return stu
}

func deleteStudent(db *gorm.DB, m interface{}) int64 {
	res := db.Debug().Unscoped().Where(m).Delete(&Student{})
	return res.RowsAffected
}
