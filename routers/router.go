package routers

import (
	"goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/admin", &controllers.MainController{})

	// Auth
	beego.Router("/login", &controllers.BaseController{}, `get:LoginPage;post:LoginControl`)
	beego.Router("/logout", &controllers.BaseController{}, `get:DoLogout`)
	beego.Router("/register", &controllers.BaseController{}, `get:RegisterPage;post:RegisterControl`)
	/*beego.Router("/register", &controllers.UserController{}, `post:Register`)
	beego.Router("/setting", &controllers.UserController{}, `get:PageSetting;post:Setting`)*/

	// Categories
	beego.Router("/admin/categories", &controllers.CategoryController{}, "get:List")
	beego.Router("/admin/categories/show/:id([0-9]+)", &controllers.CategoryController{}, `get:Show`)
	beego.Router("/admin/categories/create", &controllers.CategoryController{}, "get:Create;post:Store")
	beego.Router("/admin/categories/edit/:id([0-9]+)", &controllers.CategoryController{}, `get:Edit;post:Update`)
	beego.Router("/admin/categories/delete/:id([0-9]+)", &controllers.CategoryController{}, `get:Delete`)

	// Posts
	beego.Router("/admin/posts", &controllers.PostController{}, "get:List")
	beego.Router("/admin/posts/show/:id([0-9]+)", &controllers.PostController{}, `get:Show`)
	beego.Router("/admin/posts/create", &controllers.PostController{}, "get:Create;post:Store")
	beego.Router("/admin/posts/edit/:id([0-9]+)", &controllers.PostController{}, `get:Edit;post:Update`)
	beego.Router("/admin/posts/delete/:id([0-9]+)", &controllers.PostController{}, `get:Delete`)

	// Pages
	beego.Router("/admin/pages", &controllers.PageController{}, "get:List")
	beego.Router("/admin/pages/show/:id([0-9]+)", &controllers.PageController{}, `get:Show`)
	beego.Router("/admin/pages/create", &controllers.PageController{}, "get:Create;post:Store")
	beego.Router("/admin/pages/edit/:id([0-9]+)", &controllers.PageController{}, `get:Edit;post:Update`)
	beego.Router("/admin/pages/delete/:id([0-9]+)", &controllers.PageController{}, `get:Delete`)

	// Pages
	beego.Router("/admin/users", &controllers.UserController{}, "get:List")
	beego.Router("/admin/users/create", &controllers.UserController{}, "get:Create;post:Store")
	beego.Router("/admin/users/edit/:id([0-9]+)", &controllers.UserController{}, `get:Edit;post:Update`)
	beego.Router("/admin/users/delete/:id([0-9]+)", &controllers.UserController{}, `get:Delete`)
	/*

	beego.Router("/admin/pages/show/:id([0-9]+)", &controllers.PageController{}, `get:Show`)
	beego.Router("/admin/pages/edit/:id([0-9]+)", &controllers.PageController{}, `get:Edit;post:Update`)
	 */

	//beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")

	/*
	- SignUp.
	- Account activation (send activation link email).
	- SignIn.
	- Captcha
	- Logout.
	- ResetPassword (Send reset password link email).
	- Delete account.
	- Profile/Account page.

	URI entry points

	/
	(just an example how other controllers can use sessions from redis)
	/accounts/signup
	/accounts/signin
	/accounts/signout
	/accounts/verify/(uuid)
	/accounts/delete
	/accounts/profile
	/accounts/forgotpassword
	/accounts/resetpassword/(uuid)
	 */
}

