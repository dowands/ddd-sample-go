package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"go-ddd/pkg/expection"
)

type response struct {
	Success bool      `json:"success"`
	ErrCode   string `json:"errCode"`
	ErrMessage   string `json:"errMessage"`
}

type singleResponse struct {
	response
	Data *interface{} `json:"data"`
}

func SingleResponse(ctx *gin.Context, data interface{}) {
	ctx.Abort()
	ctx.JSON(http.StatusOK, singleResponse{
		response: response{
			Success:    true,
		},
		Data:   &data,
	})
}

type pageResponse struct {
	response
	Data interface{} `json:"data"`
	TotalCount int64 `json:"totalCount"`
	PageSize int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}

func PageResponse(ctx *gin.Context, data interface{}, totalCount int64, pageSize int, pageIndex int) {
	ctx.Abort()
	ctx.JSON(http.StatusOK, pageResponse{
		response: response{
			Success:    true,
		},
		Data:   data,
		TotalCount: totalCount,
		PageSize: pageSize,
		PageIndex: pageIndex,
	})
}

func MultiResponse(ctx *gin.Context, data interface{}) {

}

func ERROR(ctx *gin.Context, err expection.Exception) {
	ctx.Abort()
	ctx.JSON(http.StatusOK, response{
		Success:    false,
		ErrCode:    cast.ToString(err.Code),
		ErrMessage: err.Message,
	})
}

