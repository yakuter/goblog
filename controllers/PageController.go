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

type PageController struct {
	BaseController
}

func (this *PageController) Prepare() {
	this.CheckLogin()
}

// LIST
func (this *PageController) List() {
	beego.ReadFromRequest(&this.Controller)

	page 				:= &models.Page{}
	pages 				:= page.GetAll()
	this.Data["pages"]   = pages

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName = "admin/pages/pages/list.tpl"
}

// CREATE
func (this *PageController) Create() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName = "admin/pages/pages/new.tpl"
}

// STORE
func (this *PageController) Store() {
	flash := beego.ReadFromRequest(&this.Controller)

	title 		:= this.GetString("title")
	slug		:= this.GetString("slug")
	content		:= this.GetString("content")

	valid 		:= validation.Validation{}

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
			page := &models.Page{
				Title:title,
				Slug:slug,
				Content:content,
			}
			switch {
			case page.ExistSlug():
				flash.Error("This slug is in use")
				flash.Store(&this.Controller)
			default:
				err := page.Create()
				if err != nil {
					return
				}
				flash.Notice("Page added successfully!")
				flash.Store(&this.Controller)
				this.Redirect("/admin/pages", 302)
				return
			}

	}

	this.Redirect("/admin/pages/create", 302)
	this.Abort("302")
	return
}

// SHOW
func (this *PageController) Show() {
	id, _ 			 := strconv.Atoi(this.Ctx.Input.Param(":id"))
	page 			 := &models.Page{Id:id}
	page.GetOne()
	this.Data["page"] = page

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName 	  = "admin/pages/pages/show.tpl"
}

// EDIT
func (this *PageController) Edit() {
	beego.ReadFromRequest(&this.Controller)

	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())

	id, _ 			 := strconv.Atoi(this.Ctx.Input.Param(":id"))
	page 			 := &models.Page{Id:id}
	page.GetOne()
	this.Data["page"] = page

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user
	this.TplName 	  = "admin/pages/pages/edit.tpl"
}

// UPDATE
func (this *PageController) Update() {
	flash 	 := beego.ReadFromRequest(&this.Controller)

	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	page := &models.Page{Id:id}
	page.GetOne()

	title 		:= this.GetString("title")
	slug		:= this.GetString("slug")
	content		:= this.GetString("content")

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
		page.Title 		 = title
		page.Slug 		 = slug
		page.Content 	 = content
		page.Time    	 = time.Now()

		switch {
		case page.ExistUpdateSlug():
			flash.Error("This slug is in use")
			flash.Store(&this.Controller)
		default:
			err := page.Update()
			if err != nil {
				return
			}
			flash.Notice("Page updated successfully!")
			flash.Store(&this.Controller)
			this.Redirect("/admin/pages", 302)
			return
		}

	}

	redirectUrl := "/admin/pages/edit/" + strconv.Itoa(id)
	this.Redirect(redirectUrl, 302)
	this.Abort("302")
	return
}

// DELETE
func (this *PageController) Delete() {
	flash 	 := beego.ReadFromRequest(&this.Controller)
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	page := &models.Page{Id:id}
	page.Delete()
	flash.Notice("Page deleted successfully!")
	flash.Store(&this.Controller)
	this.Redirect("/admin/pages", 302)
	return
}
