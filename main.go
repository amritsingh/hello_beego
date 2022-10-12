package main

import "github.com/beego/beego"

type HomeController struct {
	beego.Controller
}

func (controller *HomeController) Get() {
	controller.TplName = "index.html"
	controller.Render()
}

func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
