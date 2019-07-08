package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sea_log/common"
	"sea_log/master/etcd"
)

func Mapping(prefix string, app *gin.Engine) {
	admin := app.Group(prefix)
	admin.GET("all", ListAllLogJob)
	admin.GET("nodes", ListNodes)
	admin.GET("nodes/pree", ListNodePree)

}

func ListAllLogJob(ctx *gin.Context) {
	data := etcd.GetAllJob()
	fmt.Println(data)
	ctx.JSON(common.SuccessWithDate(data))
	return
}

func ListNodes(ctx *gin.Context) {
	data, err := etcd.GetAllNode()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err,
			"data":    "",
		})
	} else {
		ctx.JSON(common.SuccessWithDate(data))
	}
	return
}

func ListNodePree(ctx *gin.Context) {
	data := etcd.GetNodeInfo()
	ctx.JSON(common.SuccessWithDate(data))
	return
}
