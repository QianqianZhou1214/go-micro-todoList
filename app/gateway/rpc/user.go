package rpc

import (
	"context"
	"go-micro-todoList/idl/pb"
	"go-micro-todoList/pkg/e"
)

func UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil {
		if resp == nil {
			resp = &pb.UserResponse{}
		}
		resp.Code = e.Error
		return
	}
	if resp.Code != e.Success {
		return
	}
	return
}

func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil {
		if resp == nil {
			resp = &pb.UserResponse{}
		}
		resp.Code = e.Error
		return
	}
	if resp.Code != e.Success {
		return
	}
	return
}
