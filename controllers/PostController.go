package controllers

import (
	"goblog/models"
	"strconv"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
	"html/template"
)

type PostController struct {
	BaseController
}

func (this *PostController) Prepare() {
	this.CheckLogin()
}

// LIST
func (this *PostController) List() {
	beego.ReadFromRequest(&this.Controller)

	post 				:= &models.Post{}
	posts 				:= post.GetAll()
	this.Data["posts"]   = posts

	user 			 	:= this.GetSession("user").(models.User)
	this.Data["user"] 	 = user
	this.TplName 		 = "admin/pages/posts/list.tpl"
}

// CREATE
func (this *PostController) Create() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	category				:= &models.Category{}
	categories 				:= category.GetAll()
	this.Data["categories"]  = categories

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName = "admin/pages/posts/new.tpl"
}

// STORE
func (this *PostController) Store() {
	flash := beego.ReadFromRequest(&this.Controller)

	title 		:= this.GetString("title")
	slug		:= this.GetString("slug")
	content		:= this.GetString("content")

	category_id, _ := strconv.Atoi(this.GetString("category_id"))
	category    := &models.Category{Id:category_id}

	valid := validation.Validation{}

	valid.Required(title, "Title")
	valid.Required(slug, "Slug")
	valid.AlphaDash(slug, "Slug")

	switch {
	case valid.HasErrors():
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		flash.Error("Please fill title and slug correctly")
		flash.Store(&this.Controller)
	default:
		post := &models.Post{
			Title:title,
			Slug:slug,
			Content:content,
			Category:category,
		}

		switch {
			case post.ExistSlug():
				flash.Error("This slug is in use")
				flash.Store(&this.Controller)
			default:
				err := post.Create()
				if err != nil {
					return
				}
				flash.Notice("Post added successfully!")
				flash.Store(&this.Controller)
				this.Redirect("/admin/posts", 302)
				return
			}

	}

	this.Redirect("/admin/posts/create", 302)
	this.Abort("302")
	return
}

// SHOW
func (this *PostController) Show() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	post := &models.Post{Id:id}
	post.GetOne()
	this.Data["post"]  = post

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName = "admin/pages/posts/show.tpl"
}

// EDIT
func (this *PostController) Edit() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	category				:= &models.Category{}
	categories 				:= category.GetAll()
	this.Data["categories"]  = categories

	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	post := &models.Post{Id:id}
	post.GetOne()
	this.Data["post"]  		= post

	user 			 	:= this.GetSession("user").(models.User)
	this.Data["user"] 	 = user
	this.TplName 		 = "admin/pages/posts/edit.tpl"
}

// UPDATE
func (this *PostController) Update() {
	flash 	 := beego.ReadFromRequest(&this.Controller)

	id, _ 	 := strconv.Atoi(this.Ctx.Input.Param(":id"))
	post 	 := &models.Post{Id:id}
	post.GetOne()

	title 			:= this.GetString("title")
	slug			:= this.GetString("slug")
	content			:= this.GetString("content")
	category_id, _  := strconv.Atoi(this.GetString("category_id"))
	category 		:= &models.Category{Id:category_id}

	valid := validation.Validation{}

	valid.Required(title, "Title")
	valid.Required(slug, "Slug")
	valid.AlphaDash(slug, "Slug")

	switch {
	case valid.HasErrors():
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		flash.Error("Please fill title and slug correctly")
		flash.Store(&this.Controller)

	default:
		post.Title 		= title
		post.Slug 		= slug
		post.Content 	= content
		post.Category	= category
		post.Time    	= time.Now()

		switch {
		case post.ExistUpdateSlug():
			flash.Error("This slug is in use")
			flash.Store(&this.Controller)
		default:
			err := post.Update()
			if err != nil {
				return
			}
			flash.Notice("Post updated successfully!")
			flash.Store(&this.Controller)
			this.Redirect("/admin/posts", 302)
			return
		}

	}

	redirectUrl := "/admin/posts/edit/" + strconv.Itoa(id)
	this.Redirect(redirectUrl, 302)
	this.Abort("302")
}

// DELETE
func (this *PostController) Delete() {
	flash 	 := beego.ReadFromRequest(&this.Controller)
	id, _ 	 := strconv.Atoi(this.Ctx.Input.Param(":id"))
	post 	 := &models.Post{Id:id}
	post.Delete()
	flash.Notice("Post deleted successfully!")
	flash.Store(&this.Controller)
	this.Redirect("/admin/posts", 302)
	return
}