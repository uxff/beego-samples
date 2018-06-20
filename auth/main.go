package main

import (
	"github.com/astaxie/beego"
	_ "github.com/uxff/beego-samples/auth/conf/inits"
	_ "github.com/uxff/beego-samples/auth/routers"
)

func main() {
	beego.Run()
}
