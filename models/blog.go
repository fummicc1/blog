package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	// import mysql
	_ "github.com/go-sql-driver/mysql"
)

// Blog is struct
type Blog struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// BlogManager stores blog data and handles CRUD.
type BlogManager struct {
}

func init() {
	orm.RegisterModel(new(Blog))
	orm.RegisterDataBase("default", "mysql", "root:PASSWORD@tcp(localhost:3306)/blog?charset=utf8", 30)
}

// Read is a function that get blogs from mysql.
func (manager BlogManager) Read() []Blog {
	blogs := []Blog{}
	o := orm.NewOrm()
	id, err := o.QueryTable(new(Blog)).All(&blogs)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("id: %d", id)
	return blogs
}

// Create is a function that persist blog into mysql.
func (manager BlogManager) Create(blog Blog) {
	o := orm.NewOrm()
	id, err := o.Insert(&blog)

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%d", id)

}
