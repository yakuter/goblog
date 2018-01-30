package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

type User struct {
	Id           	int  	`orm:"pk"`
	NameSurname  	string		`orm:"null"`
	UserName     	string
	Email        	string
	Password     	string
	Url          	string  	`orm:"null"`
	Info         	string  	`orm:"null"`
	RegisterTime    time.Time 	`orm:"auto_now_add;type(datetime)"`
}

// Define Page table's name
func (u *User) TableName() string {
	return "users"
}

// Get all records form table
func (user User) GetAll() (users []User) {
	o := orm.NewOrm()
	o.QueryTable("users").All(&users);
	return
}

func (user *User) GetOne() (err error) {
	o := orm.NewOrm()
	if err = o.QueryTable("users").Filter("id", user.Id).One(user); err != nil {
		beego.Info(err)
	}
	return
}
func (user *User) GetByUsername() (err error) {
	o := orm.NewOrm()
	if err = o.QueryTable("users").Filter("user_name", user.UserName).One(user); err != nil {
		beego.Info(err)
	}
	return
}

func (user User) Create() (err error) {
	o := orm.NewOrm()
	if _, err = o.Insert(&user); err != nil {
		beego.Info(err)
	}
	return
}

func (user User) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&user); err != nil {
		beego.Info(err)
	}
	return
}

// Delete record
func (user User) Delete() (err error) {
	o := orm.NewOrm()
	if _, err = o.Delete(&User{Id: user.Id}); err != nil {
		beego.Info(err)
	}
	return
}

func (user User) ExistUserName() bool {
	o := orm.NewOrm()
	return o.QueryTable("users").Filter("UserName", user.UserName).Exist()
}

func (user User) ExistEmail() bool {
	o := orm.NewOrm()
	return o.QueryTable("users").Filter("Email", user.Email).Exist()
}

func (user *User) CountAll() (cnt int64) {
	var err error
	o := orm.NewOrm()
	if cnt, err = o.QueryTable("users").Count(); err != nil {
		beego.Info(err)
	}
	return
}

func init() {
	orm.RegisterModel(new(User))
}