package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"github.com/astaxie/beego/validation"
	"fmt"
	"log"
	"crypto/md5"
	"html/template"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) IsLogin() bool {
	return this.GetSession("user") != nil
}

func (this *BaseController) CheckLogin() {
	if !this.IsLogin() {
		this.Redirect("/login", 302)
		this.Abort("302")
	}
}

func (this *BaseController) DoLogin(user models.User) {
	this.SetSession("user", user)
}

func (this *BaseController) LoginPage() {
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	this.TplName = "auth/login.tpl"
}

func (this *BaseController) LoginControl() {

	username := this.GetString("user_name")
	password := this.GetString("password")

	valid := validation.Validation{}

	valid.Required(username, "Username")
	valid.Required(password, "Password")
	valid.MaxSize(password, 16, "Password")


	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	user 	:= &models.User{UserName:this.GetString("user_name")}
	err 	:= user.GetByUsername()
	switch  {
		case valid.HasErrors():
		case err != nil:
			valid.Error("Couldn't logged in")
		case PassCheck(password, user.Password) == false:
			valid.Error("Password doesn't correct")
		default:
			this.DoLogin(*user)
			this.Data["user"] = this.GetSession("user")
			this.Redirect("/admin", 302)
			return
	}

	this.Redirect("/login", 302)
	this.Abort("302")
	return

}

func (this *BaseController) RegisterPage() {
	this.TplName = "auth/register.tpl"
}

func (this *BaseController) RegisterControl() {

	namesurname		:= this.GetString("name_surname")
	username 		:= this.GetString("user_name")
	email 	 		:= this.GetString("email")
	password1 		:= this.GetString("password")
	password2 		:= this.GetString("password_confirm")

	if len(namesurname) < 1 {
		namesurname = username
	}

	valid := validation.Validation{}

	valid.Email(email, "Email")

	valid.Required(username, "Username")
	valid.Required(password1, "Password1")
	valid.Required(password2, "Password2")

	valid.MaxSize(username, 20, "Username")
	valid.MaxSize(password1, 16, "Password1")
	valid.MaxSize(password2, 16, "Password2")

	switch {
		case valid.HasErrors():
		case password1 != password2:
			valid.Error("Couldn't register")
		default:
			user := &models.User{
				NameSurname:namesurname,
				UserName:username,
				Email:email,
				Password:Md5(password1),
			}

			switch {
				case user.ExistUserName():
					valid.Error("This username is in use")
				case user.ExistEmail():
					valid.Error("This email is in use")
				default:
					err := user.Create()
					if err != nil {
						return
					}
					valid.Error(fmt.Sprintf("%v", err))
					this.Redirect("/admin", 302)
					return
			}

	}

	this.Redirect("/register", 302)
	this.Abort("302")
	return

}

func (this *BaseController) DoLogout() {
	this.DestroySession()
	this.Redirect("/login", 302)
}

func PassCheck(pwd, saved string) bool {
	return Md5(pwd) == saved
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

