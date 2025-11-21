package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	//要求 ：
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	db, err := CreateDbThree()
	if err != nil {
		fmt.Println(err)
	}
	//ems := findemployess(db)
	ems := findemployesSalaryTop(db)
	for _, em := range ems {
		fmt.Println(em)
	}
}

func CreateDbThree() (*sqlx.DB, error) {
	dsn := "root:root123@tcp(211.159.169.85:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	return db, err
}

func findemployess(db *sqlx.DB) []Employee {
	var techEmployees []Employee
	query1 := `SELECT id, name, department, salary FROM employees WHERE department = ?`
	err := db.Select(&techEmployees, query1, "技术部")
	if err != nil {
		log.Fatal(err)
	}
	return techEmployees
}

func findemployesSalaryTop(db *sqlx.DB) []Employee {
	var techEmployees []Employee
	query1 := `SELECT id, name, department, salary FROM employees order by salary desc limit 1`
	err := db.Select(&techEmployees, query1)
	if err != nil {
		fmt.Println(err)
	}
	return techEmployees
}
