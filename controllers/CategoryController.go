package controllers

import (
	"goblog/models"
	"strconv"
	"github.com/astaxie/beego/validation"
	"log"
	"github.com/astaxie/beego"
	"html/template"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Prepare() {
	this.CheckLogin()
}

// LIST
func (this *CategoryController) List() {
	//beego.ReadFromRequest(&this.Controller)

	category 				:= &models.Category{}
	categories 				:= category.GetAll()
	this.Data["categories"]  = categories

	user 			 	:= this.GetSession("user").(models.User)
	this.Data["user"] 	 = user
	this.TplName 		 = "admin/pages/categories/list.tpl"
}

// CREATE
func (this *CategoryController) Create() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	user 			 	:= this.GetSession("user").(models.User)
	this.Data["user"] 	 = user
	this.TplName = "admin/pages/categories/new.tpl"
}

// STORE
func (this *CategoryController) Store() {
	flash := beego.ReadFromRequest(&this.Controller)

	title 		:= this.GetString("title")
	slug		:= this.GetString("slug")
	description	:= this.GetString("description")

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
			category := &models.Category{
				Title:title,
				Slug:slug,
				Description:description,
			}

			switch {
			case category.ExistSlug():
				flash.Error("This slug is in use")
				flash.Store(&this.Controller)
			default:
				err := category.Create()
				if err != nil {
					return
				}
				flash.Notice("Category added successfully!")
				flash.Store(&this.Controller)
				this.Redirect("/admin/categories", 302)
				return
			}

	}

	this.Redirect("/admin/categories/create", 302)
	this.Abort("302")
	return

}

// SHOW
func (this *CategoryController) Show() {
	id, _ 					:= strconv.Atoi(this.Ctx.Input.Param(":id"))
	category 				:= &models.Category{Id:id}
	category.GetOne()
	this.Data["category"]  	 = category

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName 	  = "admin/pages/categories/show.tpl"
}

// EDIT
func (this *CategoryController) Edit() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	id, _ 					:= strconv.Atoi(this.Ctx.Input.Param(":id"))
	category 				:= &models.Category{Id:id}
	category.GetOne()
	this.Data["category"]  	 = category

	user 			   		:= this.GetSession("user").(models.User)
	this.Data["user"] 		 = user
	this.TplName 			 = "admin/pages/categories/edit.tpl"
}

// UPDATE
func (this *CategoryController) Update() {
	flash 	 := beego.ReadFromRequest(&this.Controller)

	id, _ 	 := strconv.Atoi(this.Ctx.Input.Param(":id"))
	category := &models.Category{Id:id}
	category.GetOne()

	title 		:= this.GetString("title")
	slug		:= this.GetString("slug")
	description	:= this.GetString("description")

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
			category.Title 		 = title
			category.Slug 		 = slug
			category.Description = description

			switch {
			case category.ExistUpdateSlug():
				flash.Error("This slug is in use")
				flash.Store(&this.Controller)
			default:
				err := category.Update()
				if err != nil {
					return
				}
				flash.Notice("Category updated successfully!")
				flash.Store(&this.Controller)
				this.Redirect("/admin/categories", 302)
				return
			}

	}

	redirectUrl := "/admin/categories/edit/" + strconv.Itoa(id)
	this.Redirect(redirectUrl, 302)
	this.Abort("302")
	return
}

// DELETE
func (this *CategoryController) Delete() {
	flash 	 := beego.ReadFromRequest(&this.Controller)
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	category := &models.Category{Id:id}
	category.Delete()
	flash.Notice("Category deleted successfully!")
	flash.Store(&this.Controller)
	this.Redirect("/admin/categories", 302)
	return
}