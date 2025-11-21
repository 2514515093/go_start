package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	age       int
	Email     string
	PostCount int    //文章数量
	Posts     []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint
	CommentCount  int       //评论数量
	CommentStatus string    //评论状态
	Comments      []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}

func main() {
	//题目1：模型定义,
	//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	//要求 ：
	//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。,
	//编写Go代码，使用Gorm创建这些模型对应的数据库表。,
	//,
	//,
	//题目2：关联查询,
	//基于上述博客系统的模型定义。
	//要求 ：
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。,
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。,
	//,
	//,
	//题目3：钩子函数,
	//继续使用博客系统的模型。
	//要求 ：
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	db, _ := CreateDbFive()
	//insertUsers(db)
	//insertPosts(db)
	//insertComments(db)
	//users := findUserPostsComments(db, 1)
	//fmt.Println(users)
	//post := findUserPostsByCmmax(db)
	//fmt.Println(post)
	deleteComment(db, 4)

}

func CreateDbFive() (*gorm.DB, error) {
	dsn := "root:root123@tcp(211.159.169.85:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//创建表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	return db, err
}

func findUserPostsComments(db *gorm.DB, userId uint) []User {
	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。,
	//var userPosts []User
	//db.Debug().Where("id = ?", userId).Find(&userPosts)
	var posts []User
	db.Debug().Preload("Posts.Comments").Where("id = ?", userId).Find(&posts)
	return posts
}

func findUserPostsByCmmax(db *gorm.DB) Post {
	//使用Gorm查询评论数量最多的文章信息。,
	//var userPosts []User
	//db.Debug().Where("id = ?", userId).Find(&userPosts)
	var posts []Post
	db.Debug().Preload("Comments").Find(&posts)
	var cmmax Post
	temp := 0
	for _, post := range posts {
		if len(post.Comments) > temp {
			cmmax = post
			temp = len(post.Comments)
		}
	}
	return cmmax
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	//模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
	tx.Debug().Model(&User{}).Where("id = ?", p.UserID).Select("PostCount").Updates((map[string]interface{}{
		"PostCount": gorm.Expr("post_count + ?", 1),
	}))
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

	// 先查出完整数据 不然会拿不到 PostID
	tx.Where("id = ?", c.ID).First(c)
	var comments []Comment
	tx.Debug().Where("post_id = ?", c.PostID).Find(&comments)
	//AfterDelete  查到本身  BeforeDelete 查不到
	if len(comments) == 0 {
		fmt.Println("无评论啊无评论")
		tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).Updates((map[string]interface{}{
			"comment_status": "无评论",
			"comment_count":  0,
		}))
	}
	return nil
}

func deleteComment(db *gorm.DB, commentID uint) {
	db.Debug().Delete(&Comment{Model: gorm.Model{ID: commentID}})
}

func insertUsers(db *gorm.DB) {
	users := []User{
		{Name: "张三", Email: "zhangsan@qq.com", age: 18},
		{Name: "李四", Email: "lisi@qq.com", age: 22},
		{Name: "王五", Email: "wangwu@qq.com", age: 21},
	}
	db.Debug().Create(&users)
}

func insertPosts(db *gorm.DB) {
	posts := []Post{
		{Title: "go 入门", Content: "go很简单", UserID: 1},
		{Title: "GORM 使用", Content: "GORM很强大", UserID: 1},
		{Title: "go 深入", Content: "go 很难", UserID: 2},
	}
	db.Debug().Create(&posts)
}

func insertComments(db *gorm.DB) {
	comments := []Comment{
		{Content: "很好", PostID: 1},
		{Content: "不错", PostID: 1},
		{Content: "一般", PostID: 2},
	}
	db.Debug().Create(&comments)
	db.Model(&Post{}).Where("id = ?", 1).
		Updates(map[string]interface{}{
			"comment_count":  2,
			"comment_status": "有评论",
		})
	db.Model(&Post{}).Where("id = ?", 2).
		Updates(map[string]interface{}{
			"comment_count":  1,
			"comment_status": "有评论",
		})
}
