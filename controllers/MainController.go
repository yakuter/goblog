package controllers

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goblog/models"
	"time"
	"math/rand"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/goblog?charset=utf8")
}

type MainController struct {
	BaseController
}

func (this *MainController) Prepare() {
	this.CheckLogin()
}

func (this *MainController) Get() {

	oQuery, _ 	:= orm.NewQueryBuilder("mysql")
	qb   		:= oQuery.Select("id, title, slug, time").From("posts").OrderBy("time").Desc().Limit(5)
	sql  		:= qb.String()

	type Post struct {
		Id			int
		Title   	string
		Slug   		string
		Time    	time.Time
	}

	var res []Post
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&res)

	this.Data["recentPosts"]  = res

	category 				:= &models.Category{}
	this.Data["catNumber"]   = category.CountAll()

	post 					:= &models.Post{}
	this.Data["postNumber"]  = post.CountAll()

	page 					:= &models.Page{}
	this.Data["pageNumber"]  = page.CountAll()

	userNum 				:= &models.User{}
	this.Data["userNumber"]  = userNum.CountAll()

	this.Data["word"]	= GetWord()

	user 			 := this.GetSession("user").(models.User)
	this.Data["user"] = user

	this.TplName 	  = "admin/pages/homepage.tpl"
}

func GetWord() string {

	words := []string{
		"Technology is anything that wasn’t around when you were born. <br> <em> Alan Kay (Computer Scientist)</em>",
		"Any sufficiently advanced technology is equivalent to magic. <br> <em> Arthur C. Clarke (Author)</em>",
		"Just because something doesn’t do what you planned it to do doesn’t mean it’s useless. <br> <em> Thomas Edison (Inventor)</em>",
		"It has become appallingly obvious that our technology has exceeded our humanity. <br> <em> Albert Einstein (Scientist)</em>",
		"Einstein worried that advancements in technology will exceed our humanity. <br> <em> Pablo Picasso (Artist)</em>",
		"One machine can do the work of fifty ordinary men.  No machine can do the work of one extraordinary man. <br> <em> Elbert Hubbard (Author)</em>",
		"Technology is a word that describes something that doesn’t work yet. <br> <em> Douglas Adams (Author)</em>",
		"Humanity is acquiring all the right technology for all the wrong reasons. <br> <em> R. Buckminster Fuller (Inventor and Author)</em>",
		"The human spirit must prevail over technology. <br> <em> Albert Einstein (Scientist)</em>",
		"The great myth of our times is that technology is communication. <br> <em> Libby Larsen (Composer)</em>",
		"You cannot endow even the best machine with initiative; the jolliest steamroller will not plant flowers. <br> <em> Walter Lippmann (Author)</em>",
		"We are stuck with technology when what we really want is just stuff that works. <br> <em> Douglas Adams (Author)</em>",
		"Technology made large populations possible; large populations now make technology indispensable. <br> <em> Joseph Krutch (Writer)</em>",
		"The real danger is not that computers will begin to think like men, but that men will begin to think like computers. <br> <em> Sydney Harris (Journalist)</em>",
		"If we continue to develop our technology without wisdom or prudence, our servant may prove to be our executioner. <br> <em> Omar Bradley (General, US Army)</em>",
		"The art challenges the technology, and the technology inspires the art. <br> <em> John Lasseter (Director)</em>",
		"Science and technology revolutionize our lives, but memory, tradition and myth frame our response. <br> <em> Arthur Schlesinger (Historian)</em>",
	}

	rand.Seed(time.Now().Unix())
	word := words[rand.Intn(len(words))]

	return word
}