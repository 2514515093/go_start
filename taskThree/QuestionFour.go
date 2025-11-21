package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Book struct {
	ID     int             `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}

func main() {
	//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	//要求 ：
	//定义一个 Book 结构体，包含与 books 表对应的字段。,
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	db, _ := CreateDbFour()
	var books []Book
	err := db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	if err != nil {
		fmt.Println(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}
}

func CreateDbFour() (*sqlx.DB, error) {
	dsn := "root:root123@tcp(211.159.169.85:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	return db, err
}
