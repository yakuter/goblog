package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

// Define Page table
type Page struct {
	Id      	int
	Title   	string `orm:"size(191)"`
	Slug   		string `orm:"size(191)"`
	Content 	string `orm:"type(text)`
	Time    	time.Time `orm:"auto_now_add;type(datetime)"`
}

// Define Page table's name
func (u *Page) TableName() string {
	return "pages"
}

// Get all records form table
func (page Page) GetAll() (pages []Page) {
	o := orm.NewOrm()
	o.QueryTable("pages").All(&pages)
	return
}

// Get one record form table
func (page *Page) GetOne() (err error) {
	o := orm.NewOrm()
	if err = o.QueryTable("pages").Filter("id", page.Id).One(page); err != nil {
		beego.Info(err)
	}
	return
}

// Add new record
func (page Page) Create() (err error) {
	o := orm.NewOrm()
	if _, err = o.Insert(&page); err != nil {
		beego.Info(err)
	}
	return
}

// Update record
func (page Page) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&page); err != nil {
		beego.Info(err)
	}
	return
}

// Delete record
func (page Page) Delete() (err error) {
	o := orm.NewOrm()
	if _, err = o.Delete(&Page{Id: page.Id}); err != nil {
		beego.Info(err)
	}
	return
}

func (page Page) ExistSlug() bool {
	o := orm.NewOrm()
	return o.QueryTable("pages").Filter("Slug", page.Slug).Exist()
}

func (page Page) ExistUpdateSlug() bool {
	o := orm.NewOrm()
	return o.QueryTable("pages").Exclude("id", page.Id).Filter("Slug", page.Slug).Exist()
}

func (page *Page) CountAll() (cnt int64) {
	var err error
	o := orm.NewOrm()
	if cnt, err = o.QueryTable("pages").Count(); err != nil {
		beego.Info(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(Page))
}