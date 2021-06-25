package controller

import (
	"github.com/gin-gonic/gin"
	controllers "go-ddd/pkg/controller"
	"go-ddd/pkg/expection"
	"go-ddd/pkg/global_expection"
)

func Route() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(global_expection.Recovery(global_expection.RecoveryHandler))
	r.Use(CrossDomain())

	r.GET("customer", CustomerList())
	r.POST("customer", CustomerAdd())

	r.NoRoute(func(context *gin.Context) {
		controllers.ERROR(context, expection.Exception{Code: 404, Message: "NOT FOUND"})
	})
	return
}