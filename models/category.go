package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

// Define Category table
type Category struct {
	Id      		int		`orm:"pk"`
	Title   		string 	`orm:"size(191)"`
	Slug   			string 	`orm:"size(191)"`
	Description   	string
}

// Define Category table's name
func (u *Category) TableName() string {
	return "categories"
}

// Get all records form table
func (category Category) GetAll() (categories []Category) {
	o := orm.NewOrm()

	o.QueryTable("categories").All(&categories)
	return
}

func (category *Category) GetOne() (err error) {
	o := orm.NewOrm()
	if err = o.Read(category); err != nil {
		beego.Info(err)
	}
	return
}

// Add new record
func (category Category) Create() (err error) {
	o := orm.NewOrm()
	if _, err = o.Insert(&category); err != nil {
		beego.Info(err)
	}
	return
}

// Update record
func (category Category) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&category); err != nil {
		beego.Info(err)
	}
	return
}

// Delete record
func (category Category) Delete() (err error) {
	o := orm.NewOrm()
	if _,err = o.Delete(&Category{Id: category.Id}); err != nil {
		beego.Info(err)
	}
	return
}

func (category Category) ExistSlug() bool {
	o := orm.NewOrm()
	return o.QueryTable("categories").Filter("Slug", category.Slug).Exist()
}

func (category Category) ExistUpdateSlug() bool {
	o := orm.NewOrm()
	return o.QueryTable("categories").Exclude("id", category.Id).Filter("Slug", category.Slug).Exist()
}

func (category *Category) CountAll() (cnt int64) {
	var err error
	o := orm.NewOrm()
	if cnt, err = o.QueryTable("categories").Count(); err != nil {
		beego.Info(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(Category))
}