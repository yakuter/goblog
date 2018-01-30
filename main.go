package main

import (
	_ "goblog/routers"
	"github.com/astaxie/beego"
)

func main() {

	/*
	// Database alias.
	name := "goblog"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := false

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	*/

	//authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	//beego.InsertFilter("*", beego.BeforeRouter,authPlugin)
	beego.Run()
}

/*
func SecretAuth(username, password string) bool {
	if username == "hello" && password == "world" {
		return true
	}
	return false
}
*/