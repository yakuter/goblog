package controllers

import (
	"goblog/models"
	"strconv"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"fmt"
	"log"
	"html/template"
)

type UserController struct {
	BaseController
}

func (this *UserController) Prepare() {
	this.CheckLogin()
}

// LIST
func (this *UserController) List() {
	beego.ReadFromRequest(&this.Controller)

	user 				:= &models.User{}
	users 				:= user.GetAll()
	this.Data["users"]   = users


	this.TplName = "admin/pages/users/list.tpl"
}

// CREATE
func (this *UserController) Create() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	this.TplName = "admin/pages/users/new.tpl"
}

// STORE
func (this *UserController) Store() {
	flash := beego.ReadFromRequest(&this.Controller)

	namesurname 		:= this.GetString("name_surname")
	username 			:= this.GetString("user_name")
	email 				:= this.GetString("email")
	password	 		:= this.GetString("password")
	url			 		:= this.GetString("url")
	info				:= this.GetString("info")

	valid := validation.Validation{}

	valid.Email(email, "Email")

	valid.Required(username, "Username")
	valid.Required(password, "Password")

	valid.MaxSize(username, 20, "Username")
	valid.MaxSize(password, 16, "Password")

	switch {
	case valid.HasErrors():
		valid.Error("Problem creating user!")
		flash.Error("Problem creating user!")
		flash.Store(&this.Controller)
	default:
		user := &models.User{
			NameSurname	:namesurname,
			UserName	:username,
			Email		:email,
			Password	:Md5(password),
			Url			:url,
			Info		:info,
		}
		switch {
			case user.ExistUserName():
				valid.Error("This username is in use!")
				flash.Error("This username is in use!")
				flash.Store(&this.Controller)
			case user.ExistEmail():
				valid.Error("This email is in use!")
				flash.Error("This email is in use!")
				flash.Store(&this.Controller)
			default:
				err := user.Create()
				if err != nil {
					return
				}
				valid.Error(fmt.Sprintf("%v", err))
				flash.Notice("User added successfully!")
				flash.Store(&this.Controller)
				this.Redirect("/admin/users", 302)
				return
		}

	}

	this.Redirect("/admin/users/create", 302)
	this.Abort("302")
	return
}

// EDIT
func (this *UserController) Edit() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	user := &models.User{Id:id}
	user.GetOne()

	this.Data["user"]  	= user
	this.TplName 		= "admin/pages/users/edit.tpl"
}

// Save new record function
func (this *UserController) Update() {
	flash 	 := beego.ReadFromRequest(&this.Controller)

	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	user := &models.User{Id:id}
	user.GetOne()

	namesurname 		:= this.GetString("name_surname")
	username 			:= this.GetString("user_name")
	email 				:= this.GetString("email")
	password	 		:= this.GetString("password")
	url			 		:= this.GetString("url")
	info				:= this.GetString("info")

	valid := validation.Validation{}

	valid.Email(email, "Email")

	valid.Required(username, "Username")
	valid.Required(password, "Password")

	valid.MaxSize(username, 20, "Username")
	valid.MaxSize(password, 16, "Password")

	switch {
	case valid.HasErrors():
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		valid.Error("Problem creating user!")
		flash.Error("Problem creating user!")
		flash.Store(&this.Controller)
	default:
		user := &models.User{
			NameSurname		:namesurname,
			UserName		:username,
			Email			:email,
			Password		:Md5(password),
			Url				:url,
			Info			:info,
			RegisterTime   	:time.Now(),
		}
		switch {
			case user.ExistUserName():
				valid.Error("This username is in use!")
				flash.Error("This username is in use!")
				flash.Store(&this.Controller)
			case user.ExistEmail():
				valid.Error("This email is in use!")
				flash.Error("This email is in use!")
				flash.Store(&this.Controller)
			default:
				err := user.Update()
				if err != nil {
					return
				}
				valid.Error(fmt.Sprintf("%v", err))
				flash.Notice("User updated successfully!")
				flash.Store(&this.Controller)
				this.Redirect("/admin/users", 302)
				return
		}

	}

	redirectUrl := "/admin/users/edit/" + strconv.Itoa(id)
	this.Redirect(redirectUrl, 302)
	this.Abort("302")
	return
}

// Delete record function
func (this *UserController) Delete() {
	flash 	 := beego.ReadFromRequest(&this.Controller)
	id, _ 	 := strconv.Atoi(this.Ctx.Input.Param(":id"))
	user 	 := &models.User{Id:id}
	user.Delete()
	flash.Notice("User deleted successfully!")
	flash.Store(&this.Controller)
	this.Redirect("/admin/users", 302)
	return
}