package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	_const "go-ddd/config/const"
	"go-ddd/internal/app/command"
	"go-ddd/internal/app/query"
	"go-ddd/internal/client/expection"
	. "go-ddd/internal/client/request_object"
	"go-ddd/pkg/controller"
	expection2 "go-ddd/pkg/expection"
)

func CustomerList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var page int
		if p, ok := c.Get("page"); ok {
			page = cast.ToInt(p)
		}
		if page <= 0 {
			page = 1
		}
		list := query.ListCustomer(page)
		total := query.TotalCustomer()
		controller.PageResponse(c, list, total, page, _const.PERPAGE)
	}
}

func CustomerAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		ro := CusomterRO{
			c.GetInt64("idEmployee"), c.GetString("name"),
		}
		var err expection2.Exception
		if err = command.AddCustomer(ro); err == expection.Expections.EMPTY {
			controller.SingleResponse(c, nil)
		}else{
			controller.ERROR(c, err)
		}
	}
}