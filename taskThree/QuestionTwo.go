package main

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type account struct {
	gorm.Model
	Blance decimal.Decimal
}

type transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        decimal.Decimal
}

func main() {
	//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
	//（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	//要求 ：
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
	//需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	//并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	db, _ := CreateDbTwo()
	//saveacc(db)
	transation(db, func() {
		err := atob(db, 1, 2, decimal.NewFromInt(100))
		if err != nil {
			fmt.Println(err)
		}
	})

}

func CreateDbTwo() (*gorm.DB, error) {
	dsn := "root:root123@tcp(211.159.169.85:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//创建表
	db.AutoMigrate(&account{})
	db.AutoMigrate(&transaction{})
	return db, err
}

func atob(db *gorm.DB, fromId int64, toId int64, amount decimal.Decimal) error {
	accountA := findBlance(db, fromId)
	accountB := findBlance(db, toId)
	fmt.Println(accountA.Blance)
	fmt.Println(amount)
	if accountA.Blance.LessThan(amount) {
		return errors.New("余额小于100元，无法转账")
	}
	accountB.Blance = accountB.Blance.Add(decimal.NewFromInt(100))
	accountA.Blance = accountA.Blance.Sub(decimal.NewFromInt(100))
	updateacc(db, []*account{&accountA, &accountB})
	saveatt(db, fromId, toId, amount)
	return nil
}

func saveatt(db *gorm.DB, fromId int64, toId int64, amount decimal.Decimal) {
	db.Debug().Create(&transaction{
		FromAccountId: uint(fromId),
		ToAccountId:   uint(toId),
		Amount:        amount,
	})
}

func saveacc(db *gorm.DB) {
	db.Debug().Create(&[]account{
		{Blance: decimal.NewFromInt(100)},
		{Blance: decimal.NewFromInt(500)},
	})
}

func updateacc(db *gorm.DB, accs []*account) {
	for _, acc := range accs {
		db.Debug().Model(&account{Model: gorm.Model{ID: acc.ID}}).Select("*").Updates(acc)
	}
}

func findBlance(db *gorm.DB, id int64) account {
	var account account
	db.Debug().Where("id = ?", id).First(&account)
	return account
}

func transation(db *gorm.DB, f func()) {
	db.Transaction(func(tx *gorm.DB) error {
		f()
		return nil
	})
}
