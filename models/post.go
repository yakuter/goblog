package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

// Define Post table
type Post struct {
	Id      	int
	Title   	string `orm:"size(191)"`
	Slug   		string `orm:"size(191)"`
	Content 	string `orm:"type(text)`
	Category  	*Category  `orm:"rel(fk)"`
	Time    	time.Time `orm:"auto_now_add;type(datetime)"`
}

// Define Post table's name
func (u *Post) TableName() string {
	return "posts"
}

// Get all records form table
func (post Post) GetAll() (posts []Post) {
	o := orm.NewOrm()
	o.QueryTable("posts").RelatedSel().All(&posts)
	return
}

// Add new record
func (post Post) Create() (err error) {
	o := orm.NewOrm()
	if _, err = o.Insert(&post); err != nil {
		beego.Info(err)
	}
	return
}

func (post *Post) GetOne() (err error) {
	o := orm.NewOrm()
	if err = o.QueryTable("posts").Filter("id", post.Id).RelatedSel().One(post); err != nil {
		beego.Info(err)
	}
	return
}

func (post Post) Delete() (err error) {
	o := orm.NewOrm()
	if _, err =o.Delete(&Post{Id: post.Id}); err != nil {
		beego.Info(err)
	}
	return
}

func (post Post) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&post); err != nil {
		beego.Info(err)
	}
	return
}

func (post Post) ExistSlug() bool {
	o := orm.NewOrm()
	return o.QueryTable("posts").Filter("Slug", post.Slug).Exist()
}

func (post Post) ExistUpdateSlug() bool {
	o := orm.NewOrm()
	return o.QueryTable("posts").Exclude("id", post.Id).Filter("Slug", post.Slug).Exist()
}

func (post *Post) CountAll() (cnt int64) {
	var err error
	o := orm.NewOrm()
	if cnt, err = o.QueryTable("posts").Count(); err != nil {
		beego.Info(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(Post))
}