package http

import (
	"go-micro-todoList/app/gateway/rpc"
	"go-micro-todoList/idl/pb"
	"go-micro-todoList/pkg/ctl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetUserInfo-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-TaskCreate"))
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetUserInfo-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskUpdate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-TaskUpdate"))
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetUserInfo-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskDelete(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-TaskDelete"))
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func ListTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetUserInfo-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-TaskList"))
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}

func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetUserInfo-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	taskRes, err := rpc.TaskGet(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-TaskGet"))
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, taskRes))
}
