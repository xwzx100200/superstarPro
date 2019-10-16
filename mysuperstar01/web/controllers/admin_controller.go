package controllers

import (
	"github.com/kataras/iris/mvc"
	"superstartPro/mysuperstar01/services"
	"github.com/kataras/iris"
)

type AdminController struct {
	Ctx iris.Context
	Service services.SuperstarService
}

func (admin *AdminController) Get() mvc.Result {
	//Todo.....
	return nil
}

func (admin *AdminController) GetEdit() mvc.Result {
	//Todo.....
	return nil
}

func (admin *AdminController) PostSave() mvc.Result {
	//Todo.....
	return nil
}

func (admin *AdminController) GetDelete() mvc.Result {
	//Todo.....
	return nil
}
