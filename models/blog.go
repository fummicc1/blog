package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	// import mysql
	_ "github.com/go-sql-driver/mysql"
)

var logger *log.Logger = logs.GetLogger()

// Blog is struct
type Blog struct {
	ID    int    `orm:"auto;pk" json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// BlogManager handles CRUD.
type BlogManager struct {
	O orm.Ormer
}

func init() {
	orm.RegisterModel(new(Blog))
	orm.RegisterDataBase("default", "mysql", "root:PASSWORD@tcp(localhost:3306)/blog?charset=utf8", 30)
}

// Read is a function that get blogs from mysql.
func (manager *BlogManager) Read() []Blog {
	blogs := []Blog{}
	o := manager.O
	_, err := o.QueryTable(new(Blog)).All(&blogs)
	if err != nil {
		logs.Error(err)
	}

	logs.Info("success reading blogs")

	if len(blogs) <= 10 {
		return []Blog{
			Blog{
				Title: "タイトル１",
				Body:  "これはテストデータ１",
			},
		}
	}

	return blogs
}

// Create is a function that persist blog into mysql.
func (manager *BlogManager) Create(blog Blog) error {
	o := manager.O
	id, err := o.Insert(&blog)

	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	fmt.Printf("%d", id)

	return nil

}
