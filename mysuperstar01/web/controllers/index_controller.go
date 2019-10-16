package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"superstartPro/mysuperstar01/services"
)

type IndexController struct {
	Ctx iris.Context
	service services.SuperstarService
}

func (index *IndexController)Get() mvc.Result  {
	//todo......
	return nil
}

func (index *IndexController) GetBy(id int) mvc.Result  {
	//todo......
	return nil
}

func (index *IndexController) GetSearch() mvc.Result {
	//todo......
	return nil
}
