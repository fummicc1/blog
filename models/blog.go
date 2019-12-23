package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	// import mysql
	_ "github.com/go-sql-driver/mysql"
)

// Blog is struct
type Blog struct {
	ID   int       `orm:"auto;pk" form:"-"`
	Body string    `form:"body"`
	Date time.Time `form:"-"`
}

// BlogManager handles CRUD.
type BlogManager struct {
	O orm.Ormer
}

func init() {
	logs.Debug("init blog in \"models\" package")
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

	logs.Debug("success reading blogs")

	if len(blogs) <= 10 {
		return []Blog{
			Blog{
				Body: "これはテストデータ１",
				Date: time.Now(),
			},
			Blog{
				Body: "これはテストデータ2",
				Date: time.Now(),
			},
		}
	}

	return blogs
}

// Create is a function that persist blog into mysql.
func (manager *BlogManager) Create(blog *Blog) error {
	o := manager.O
	id, err := o.Insert(blog)

	if err != nil {
		logs.Error("%s", err)
		return err
	}
	logs.Debug("id: %d", id)
	return nil
}
